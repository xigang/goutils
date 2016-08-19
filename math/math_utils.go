package math

import (
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

//两个集合的交集
func Intersection(slice1 []string, slice2 []string) []string {
	var orgin []string
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				orgin = append(orgin, v1)
			}
		}
	}
	return orgin
}

//两个集合的差集
func Difference(slice1 []string, slice2 []string) []string {
	var diff []string
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
