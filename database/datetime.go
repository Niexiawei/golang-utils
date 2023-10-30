package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	golangutils "github.com/Niexiawei/golang-utils"
	"time"
)

var (
	ErrWrongDateFormat = errors.New("错误的时间格式")
)

type LocalTime time.Time

func (l *LocalTime) UnmarshalJSON(bytes []byte) error {
	t, err := time.Parse("\"2006-01-02 15:04:05\"", golangutils.BytesToString(bytes))
	if err != nil {
		return err
	}
	*l = LocalTime(t)
	return nil
}

func (l LocalTime) Value() (driver.Value, error) {
	val := time.Time(l)
	return val.Format(time.DateTime), nil
}

func (l LocalTime) MarshalJSON() ([]byte, error) {
	val := time.Time(l)
	return []byte(fmt.Sprintf(`"%v"`, val.Format(time.DateTime))), nil
}

func (l *LocalTime) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*l = LocalTime(value)
		return nil
	}
	return fmt.Errorf("%w,time:%v", ErrWrongDateFormat, v)
}
