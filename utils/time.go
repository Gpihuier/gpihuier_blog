package utils

import (
	"time"
)

const (
	DEFAULT       = "2006-01-02 15:04:05"
	DEFAULT_YMD   = "2006-01-02"
	DEFAULT_YMDHI = "2006-01-02 15:04"
	RFC3339Milli  = "2006-01-02T15:04:05.000Z07:00"
)

// TimestampToDate 时间戳转字符串
func TimestampToDate(timestamp int64, layout string) string {
	if layout == "" {
		return time.Unix(timestamp, 0).Format(DEFAULT)
	}
	return time.Unix(timestamp, 0).Format(layout)
}

// DateToTime 解析字符串格式的时间
func DateToTime(date, layout string) (*time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}
	obj, err := time.ParseInLocation(layout, date, loc)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
