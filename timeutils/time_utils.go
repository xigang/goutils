package timeutils

import (
	"time"
)

const (
	DURATION_DAY = time.Hour * 24
)

//获取两个时间之间间隔的天数
func diffTime(t1, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

func DiffDays(t1, t2 time.Time) int64 {
	return (diffTime(t1, t2) / DURATION_DAY).Nanoseconds()
}
