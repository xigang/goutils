package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

const (
	DATE         = "2006-01-02"
	TIME         = "2006-01-02 15:04:05"
	DEFAULT_TIME = "0001-01-01 00:00:00"
	DEFAULT_DATE = "0001-01-01"
)

//将string类型的浮点数转换为浮点数，bit为32 or 64.
func StringToFloat(s string, bit int) (float64, error) {
	return strconv.ParseFloat(s, bit)
}

//将string类型的字符串转换为整数
func StringToInt(s string, base int, bit int) (int64, error) {
	return strconv.ParseInt(s, base, bit)
}

//将string的字符串转换为布尔
func StringToBoolean(s string) (bool, error) {
	return strconv.ParseBool(s)
}

//将时间转换为字符串类型的日期
func TimeToStringDate(t time.Time) string {
	return t.Format(DATE)
}

//将时间转换为字符处类型的日期+时间
func TimeToStringDateTime(t time.Time) string {
	return t.Format(TIME)
}

//将string类型的字符串转换为日期
func ParseStringDate(t string) (time.Time, error) {
	return time.Parse(DATE, t)
}

//将string类型的字符串转化为日期时间
func ParseStringDateInLocation(t string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(TIME, t, loc)
}

//将整型转换为字符串类型
func FmtIntToString(value interface{}) (str string, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int8:
		str = fmt.Sprintf("%d", value.(int8))
	case reflect.Int:
		str = fmt.Sprintf("%d", value.(int))
	case reflect.Int16:
		str = fmt.Sprintf("%d", value.(int16))
	case reflect.Int32:
		str = fmt.Sprintf("%d", value.(int32))
	case reflect.Int64:
		str = fmt.Sprintf("%d", value.(int64))
	case reflect.String:
		str = value.(string)
	default:
		err = fmt.Errorf("type is valid: %s", reflect.TypeOf(value).String())
	}
	return
}

//将一个对象转换为json
func ToJson(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		res = []byte("")
	}
	return string(res), err
}
