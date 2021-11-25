package tool

import (
	"encoding/json"
	"math"
	"strconv"
)

func ReverseIntSlice(s []int64) []int64 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SecondsToInt64(costTime float64) (int64, int64) {
	costTimeFloat := strconv.FormatFloat(costTime, 'f', 0, 64)
	costTimeInt, _ := strconv.ParseInt(costTimeFloat, 10, 64)
	min := math.Floor(float64(costTimeInt)) / float64(60)
	sec := costTimeInt % 60
	return int64(min), sec
}

func Interface2Int64(inter interface{}) int64 {
	var temp int64
	switch inter.(type) {
	case string:
		temp, _ = strconv.ParseInt(inter.(string), 10, 64)
		break
	case int64:
		temp = inter.(int64)
		break
	case int:
		s := strconv.Itoa(inter.(int))
		temp, _ = strconv.ParseInt(s, 10, 64)
		break
	case int32:
		temp = int64(inter.(int32))
		break
	case float64:
		tempStr := strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		temp, _ = strconv.ParseInt(tempStr, 10, 64)
		break
	case json.Number:
		temp, _ = inter.(json.Number).Int64()
	}
	return temp
}
