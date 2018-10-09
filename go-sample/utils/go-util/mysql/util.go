package database

import (
	"database/sql"
	"fmt"
	"go-sample/utils/go-util/config"
	"time"
)

type DateTime struct {
	TimeStamp string
}

//Parse mysql datetime to Time object
func (t DateTime) Parse() (time.Time, error) {
	tm, err := time.ParseInLocation(`2006-01-02 15:04:05`, fmt.Sprint(t.TimeStamp), config.AppConf.Location)
	return tm, err
}

//Parse Time object datetime to mysql datetime
func ParseToString(t time.Time) (s string) {
	return t.Format(`2006-01-02 15:04:05`)
}

//Parse Time object datetime to mysql datetime
func ParseToNullableString(t time.Time) (s sql.NullString) {

	emptyTime := time.Time{}
	if t == emptyTime {
		return s
	}
	s.String = t.Format(`2006-01-02 15:04:05`)
	s.Valid = true
	return s
}

//Parse Time object datetime to mysql date
func ParseToDateString(t time.Time) string {
	return t.Format(`2006-01-02`)
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, config.AppConf.Location)
}
