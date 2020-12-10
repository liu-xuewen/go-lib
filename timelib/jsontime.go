package timelib

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(t).Format(timeFormat)
	if formatted == "0001-01-01 00:00:00" || formatted == "0000-00-00 00:00:00" {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", formatted)), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t JsonTime) String() string {
	str := time.Time(t).Format(timeFormat)
	if str == "0001-01-01 00:00:00" {
		return ""
	}
	return str
}

func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
