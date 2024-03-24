package util

import (
	"strconv"
	"strings"
)

func GenerateKeys(ballColor string, redCount int) []string {
	if "red" == strings.ToLower(ballColor) {
		return generateRedKeys(redCount)
	} else {
		return generateBlueKeys()
	}
}

func generateBlueKeys() []string {
	var result = make([]string, 16)
	for i := 0; i < 16; i++ {
		itoa := strconv.Itoa(i + 1)
		if len(itoa) == 2 {
			result[i] = itoa
		} else {
			result[i] = "0" + itoa
		}
	}
	return result
}

// 根据数据生成Keys
func generateRedKeys(numbers int) []string {
	var list [][]int
	var result []string
	generate(numbers, 1, nil, &list)
	if len(list) > 0 {
		for _, v := range list {
			str := ArrToString(v, "_")
			result = append(result, str)
		}
	}
	return result
}

// 递归生成需要的数据
func generate(numbers int, start int, list []int, result *[][]int) {
	if numbers == 0 {
		var aa = make([]int, len(list))
		for i, v := range list {
			aa[i] = v
		}
		*result = append(*result, aa)
		return
	}
	for i := start; i <= 33; i++ {
		if ContainsInt(list, i) {
			continue
		}
		list2 := list[0:]
		list2 = append(list2, i)
		generate(numbers-1, i+1, list2, result)
	}

}
