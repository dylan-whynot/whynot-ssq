package service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dylan_whynot/whynot_ssq/model"
	"github.com/dylan_whynot/whynot_ssq/util"
)

type rangeKey struct {
	Key    string
	Values []string
}

func MatchRange(query *model.Query, condition *model.Condition) *model.MatchRangeResult {
	var ssqs []model.Ssq
	match(query, func(index int, ssq model.Ssq) {
		ssqs = append(ssqs, ssq)
	})
	if len(ssqs) == 0 {
		return &model.MatchRangeResult{}
	}

	rangeKeys := GenerateRangeKey(condition.RangeSize)
	rangeItems := generateRangeItems(rangeKeys)
	goroutines := 16
	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutines)
	for i := 0; i < 16; i++ {
		item := rangeItems[i]
		go func(item *model.MatchRangeItem) {
			defer waitGroup.Done()
			blue := item.Blue
			for ri, v := range rangeKeys {
				values := v.Values
				for _, ssq := range ssqs {
					if ssq.Blue != blue {
						continue
					}
					// 红球
					redInts := util.StringArrToIntArr(ssq.RedNumbers)
					// 目标数
					targetInts := util.StringArrToIntArr(values)
					// 满足条件
					for _, s := range redInts {
						if util.ContainsInt(targetInts, s) {
							rangeCount := *item.RangeCount
							rangeCount[ri] = rangeCount[ri] + 1
						}
					}
				}
			}

		}(&item)
	}
	waitGroup.Wait()
	var rangKeyStr = make([]string, len(rangeKeys))
	for i, v := range rangeKeys {
		rangKeyStr[i] = v.Key
	}
	return &model.MatchRangeResult{Query: query, RangeKey: rangKeyStr, Items: &rangeItems}
}
func PrintMatchRangeResult(query *model.Query, condition *model.Condition, print *model.PrintControl, result *model.MatchRangeResult) {
	now := time.Now().Format("2006-01-02 15:04:05")
	startYear := query.StartYear
	if query.StartYear == "" {
		startYear = "-"
	}
	endYear := query.EndYear
	if query.EndYear == "" {
		endYear = "-"
	}
	week := query.Week
	if query.Week == "" {
		week = "-"
	}
	// 排序
	items := *result.Items
	fmt.Println("########统计范围数值出现频率###", now)
	fmt.Printf("条件蓝球%s 星期[%s] 选择年份[%s,%s)  统计颜色[%s] [%d]红球组合 统计次数大于[%d]\n", query.Blue, week, startYear, endYear, condition.BallColor, condition.RedCount, print.GranterThan)

	join := strings.Join(result.RangeKey, "\t")
	fmt.Println("蓝球号\t", join)
	fmt.Println("--------------------------------------------------")
	for _, v := range items {
		var console = v.Blue + "\t"
		for _, num := range *v.RangeCount {
			console = console + strconv.Itoa(num) + "\t"
		}
		fmt.Println(console)
	}
	fmt.Println("########输出结果###", now)
}
func generateRangeItems(rangeKeys []rangeKey) []model.MatchRangeItem {
	var result = make([]model.MatchRangeItem, 16)
	for i := 0; i < 16; i++ {
		itoa := strconv.Itoa(i + 1)
		if len(itoa) != 2 {
			itoa = "0" + itoa
		}

		ints := make([]int, len(rangeKeys))
		result[i] = model.MatchRangeItem{Blue: itoa, RangeCount: &ints}
	}
	return result
}
func GenerateRangeKey(size int) []rangeKey {
	var group int
	num1 := 33 / size
	if 33%size == 0 {
		group = num1
	} else {
		group = num1 + 1
	}

	var reds = make([]string, 33)
	for i := 0; i < 33; i++ {
		if i < 9 {
			reds[i] = "0" + strconv.Itoa(i+1)
		} else {
			reds[i] = strconv.Itoa(i + 1)
		}
	}

	genKey := func(values []string) string {
		if len(values) == 1 {
			return "[" + values[0] + "]"
		}
		s := values[0]
		e := values[len(values)-1]
		return "[" + s + "-" + e + "]"
	}

	var rangeKeys = make([]rangeKey, group)
	for i := 0; i < group; i++ {
		start := i * size
		var end int
		if end = (i + 1) * size; end > 33 {
			end = 33
		}
		values := reds[start:end]
		key := genKey(values)
		rangeKeys[i] = rangeKey{Key: key, Values: values}
	}
	return rangeKeys

}
