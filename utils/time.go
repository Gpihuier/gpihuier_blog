package utils

import (
	"time"
)

func TimestampToDate(timestamp int64, layout string) string {
	if layout == "" {
		return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	}
	return time.Unix(timestamp, 0).Format(layout)
}

// DateToTime 解析字符串格式的时间
func DateToTime(layout string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}
	obj, err := time.ParseInLocation("2006-01-02 15:04:05", layout, loc)
	if err != nil {
		return time.Time{}, err
	}
	return obj, nil
}
