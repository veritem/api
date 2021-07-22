package utils

import (
	"strconv"
	"time"
)

func FormatTime(parseTime time.Time) string {
	const numBase = 10
	return strconv.FormatInt(parseTime.UnixNano()/int64(time.Millisecond), numBase)
}
