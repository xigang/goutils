package database

import (
	"testing"
)

/*
CREATE TABLE `member_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `age` varchar(45) NOT NULL,
  `address` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8
*/

const (
	DATABASE_CONFIG_PATH = "./database.conf"
)

type MemberInfo struct {
	Id      int32  `xorm:"id"`
	Name    string `xorm:"name"`
	Age     int32  `xorm:"age"`
	Address string `xorm:"address"`
}

func TestInitalDBFromConfig(t *testing.T) {
	engines, err := InitalDBFromConfig(DATABASE_CONFIG_PATH)
	if err != nil {
		t.Error("inital database config message error: ", err)
		return
	}

	t.Log("inital database config message success!!!")

	var members []MemberInfo
	err = engines["test"].Find(&members)
	if err != nil {
		t.Error("batch get member info message error: ", err)
		return
	}
	t.Log(members)

	var member MemberInfo
	_, err = engines["test"].Where("`id` = ?", 1).Get(&member)
	if err != nil {
		t.Error("select member message error: ", err)
		return
	}
	t.Log(member)

	member.Name = "Dana"
	_, err = engines["test"].Where("`id` = ?", 1).Update(&member)
	if err != nil {
		t.Error("update member_info message error: ", err)
		return
	}
	t.Log("update member_info data success.")
}
