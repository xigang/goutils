package goutils

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"time"
)

const (
	DURATION_DAY = time.Hour * 24
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

//对浮点数进行精度取舍
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

//对版本号的大小进行比较
func VersionOrdinal(version string) string {
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

//获取两个时间之间间隔的天数
func diffTime(t1, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

func DiffDays(t1, t2 time.Time) int64 {
	return (diffTime(t1, t2) / DURATION_DAY).Nanoseconds()
}

//MD5 生成
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
