package utils

import (
	"strconv"
	"time"
)

func FormatTime(parseTime time.Time) string {
	return strconv.FormatInt(parseTime.UnixNano()/int64(time.Millisecond), 10)
}
