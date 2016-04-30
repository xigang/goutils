package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

const (
	DB_DEFEAULT_CONNS = 10
)

type DBConfig struct {
	DBName    string `json:"db_name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	Location  string `json:"location"`
	IsDefault bool   `json:"is_default"`
}

type DBConfigs struct {
	DBDriver  string     `json:"db_driver"`
	DBConfigs []DBConfig `json:"db_configs"`
}

func (p *DBConfig) DSN() string {
	loc := strings.Replace(p.Location, "/", "%2F", 1)
	if loc == "" {
		loc = "Asia%2FShanghai"
	}

	charset := "utf8"
	if p.Charset != "" {
		charset = p.Charset
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s", p.User, p.Password, p.Host, p.Port, p.DBName, charset, loc)
}

//InitalDBFromConfig 创建数据库引擎,使用默认的数据库连接数
func InitalDBFromConfig(filename string) (engines map[string]*xorm.Engine, err error) {
	return InitalDBFromConfigWithConns(filename, DB_DEFEAULT_CONNS)
}

//InitalDBFromConfigWithConns 创建数据库Orm引擎,可以选择最大的连接数
func InitalDBFromConfigWithConns(filename string, conns int) (engines map[string]*xorm.Engine, err error) {
	configs := DBConfigs{}
	configs, err = GetConfig(filename)
	if err != nil {
		return
	}

	haveDefaultDb := false
	tmpEngines := make(map[string]*xorm.Engine, 0)
	for _, config := range configs.DBConfigs {
		if config.IsDefault == true && haveDefaultDb == true {
			panic("only could have one default database")
		}

		engineName := config.DBName
		newEngine, e := xorm.NewEngine(configs.DBDriver, config.DSN()) //创建数据orm引擎
		if e != nil {
			err = e
			return
		} else {
			if config.IsDefault {
				haveDefaultDb = true
				tmpEngines["default"] = newEngine
			}
			tmpEngines[engineName] = newEngine
		}
	}

	engines = tmpEngines
	if conns <= 0 {
		conns = DB_DEFEAULT_CONNS
	}

	for _, engine := range engines {
		engine.SetMaxIdleConns(conns)           //设置连接池的空闲大小
		engine.SetMaxOpenConns(conns)           //设置最大打开的连接数
		engine.Logger().SetLevel(core.LOG_INFO) //设置日志级别
	}
	return
}

//GetConfig 获取配置文件中数据库相关的数据
func GetConfig(filename string) (configs DBConfigs, err error) {
	var data []byte
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("read file content from filename error: ", err)
		return
	}

	err = json.Unmarshal(data, &configs)
	if err != nil {
		log.Fatal("unmarshal json file to structure error: ", err)
		return
	}
	return
}
