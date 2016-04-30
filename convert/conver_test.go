package convert

import (
	"reflect"
	"testing"
	"time"
)

func TestStringToFloat(t *testing.T) {
	floatNumber, err := StringToFloat("1.01", 64)
	if err != nil {
		t.Error("string to float error: ", err)
	}
	t.Log("string to float success. floatNumber: ", floatNumber)
	t.Log("floatNumber type is: ", reflect.TypeOf(floatNumber))
}

func TestStringToInt(t *testing.T) {
	number, err := StringToInt("100", 10, 64)
	if err != nil {
		t.Error("string to integer error: ", err)
	}
	t.Log("string to integer success. number: ", number)
	t.Log("number type is: ", reflect.TypeOf(number))
}

func TestStringToBoolean(t *testing.T) {
	boolean, err := StringToBoolean("hello world")
	if err != nil {
		t.Error("string to bool error: ", err)
	}
	t.Log("string to boolean sueecss. boolean: ", boolean)
	t.Log("boolean type is:", reflect.TypeOf(boolean))
}

func TestTimeToStringDate(t *testing.T) {
	date := TimeToStringDate(time.Now())
	t.Log("time to strign date success. date: ", date)
	t.Log("date type is: ", reflect.TypeOf(date))
}

func TestTimeToStringDateTime(t *testing.T) {
	datetime := TimeToStringDateTime(time.Now())
	t.Log("time to strign datetime success. date: ", datetime)
	t.Log("datetime type is: ", reflect.TypeOf(datetime))
}

func TestParseStringDate(t *testing.T) {
	td, err := ParseStringDate("2016-04-16")
	if err != nil {
		t.Error("aprse string to date error: ", err)
	}
	t.Log("parse string to date success. td: ", td)
	t.Log("td type is: ", reflect.TypeOf(td))
}

func TestParseStringDateInLocation(t *testing.T) {
	tm, err := ParseStringDateInLocation("2016-04-16 10:00:00", time.Local)
	if err != nil {
		t.Error("aprse string to date error: ", err)
	}
	t.Log("parse string to date success. tm: ", tm)
	t.Log("tm type is: ", reflect.TypeOf(tm))
}

func TestFmtToString(t *testing.T) {
	str, err := FmtIntToString(123)
	if err != nil {
		t.Log("integer to string error: ", err)
	}
	t.Log("integer to string success. str: ", str)
	t.Log("str type is: ", reflect.TypeOf(str))
}
