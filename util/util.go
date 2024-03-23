package util

import (
	"strconv"
	"strings"
)

func ArrToString(ints []int, sep string) string {
	var strArr []string
	for _, v := range ints {
		itoa := strconv.Itoa(v)
		if len(itoa) == 1 {
			itoa = "0" + itoa
		}
		strArr = append(strArr, itoa)
	}
	return strings.Join(strArr, sep)
}
func StringArrToIntArr(arr []string) []int {
	leng := len(arr)
	var ints = make([]int, leng)
	for i, v := range arr {
		ints[i], _ = strconv.Atoi(v)
	}
	return ints
}
