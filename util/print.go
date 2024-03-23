package util

import (
	"fmt"
	"slices"
)

// 根据value排序后打印
func PrintMapOrderValue(maps map[string]int, timesGreater int) {
	var newMap = make(map[int][]string)
	var keys []int
	for k, v := range maps {
		strs, ok := newMap[v]
		if ok {
			newMap[v] = append(strs, k)
		} else {
			keys = append(keys, v)
			newMap[v] = []string{k}
		}
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	for _, v := range keys {
		if timesGreater >= v {
			continue
		}
		strings := newMap[v]
		for _, str := range strings {
			fmt.Println(str, "\t", v)
		}
	}
}
