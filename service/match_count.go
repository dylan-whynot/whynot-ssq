package service

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dylan_whynot/whynot_ssq/model"

	"github.com/dylan_whynot/whynot_ssq/util"
)

const defaultSplitSize = 5000

// 满足条件数据
type MatchCountItem struct {
	Key   string
	Count int
	Codes []string
}

type MatchCountResult struct {
	Query *model.Query
	Items *[]MatchCountItem
}

// 根据蓝色球号码统计那些组合出现的次数最多
func MatchCountRedMultiTimes(query *model.Query, redCount int) *MatchCountResult {
	stringkeys := generateKeys(redCount)
	goroutines := goroutinesSize(len(stringkeys))
	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutines)
	var ssqs []model.Ssq
	match(query, func(index int, ssq model.Ssq) {
		ssqs = append(ssqs, ssq)
	})
	if len(ssqs) == 0 {
		return &MatchCountResult{Query: query}
	}
	var list = make([][]MatchCountItem, goroutines)
	for i := 0; i < goroutines; i++ {
		go func(keys *[]string, i int) {
			defer waitGroup.Done()
			start := i * defaultSplitSize
			end := (i + 1) * defaultSplitSize
			if i == (goroutines - 1) {
				end = len(*keys)
			}
			keyValues := *keys
			subArr := keyValues[start:end]
			var result = make(map[string]*MatchCountItem)
			for _, key := range subArr {
				for _, ssq := range ssqs {
					redInts := util.StringArrToIntArr(ssq.RedNumbers)
					split := strings.Split(key, "_")
					targetInts := util.StringArrToIntArr(split)
					containsSize := 0
					// 满足条件
					for _, s := range targetInts {
						if util.ContainsInt(redInts, s) {
							containsSize++
						}
					}
					if containsSize == len(split) {
						v, ok := result[key]
						if ok {
							v.Count = v.Count + 1
							v.Codes = append(v.Codes, ssq.Id)
						} else {
							result[key] = &MatchCountItem{Key: key, Count: 1, Codes: []string{ssq.Id}}
						}
					}
				}
			}
			var l = make([]MatchCountItem, len(result))
			index := 0
			for _, v := range result {
				l[index] = *v
				index++
			}
			list[i] = l
			log.Println("goroutines [", i, "] finished")
		}(&stringkeys, i)
	}
	waitGroup.Wait()
	var items []MatchCountItem
	for _, m := range list {
		items = append(items, m...)
	}
	return &MatchCountResult{Query: query, Items: &items}
}
func PrintMatchCountResult(query *model.Query, print *model.PrintControl, result *MatchCountResult) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("########输出结果###", now)
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
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})
	fmt.Printf("查询条件蓝球%s 星期[%s] 选择年份[%s,%s) %d个红球出现频率大于[%d]\n", query.Blue, week, startYear, endYear, print.RedCount, print.GranterThan)
	fmt.Println("总结果数:", len(items))
	if print.PrintIssues {
		fmt.Println("序号\t红球号码\t频率\t期号")
	} else {
		fmt.Println("序号\t红球号码\t频率")
	}
	fmt.Println("---------------------------")
	for i, v := range items {
		if print.PageSize != -1 && print.PageSize == i {
			break
		}
		if print.GranterThan >= v.Count {
			continue
		}
		if print.PrintIssues {
			fmt.Println(i+1, "\t", v.Key, "\t", v.Count, "\t", v.Codes)
		} else {
			fmt.Println(i+1, "\t", v.Key, "\t", v.Count)
		}
	}
	fmt.Println("########输出结果###", now)
}
func goroutinesSize(length int) int {
	var goroutines int
	num1 := length / defaultSplitSize
	num2 := length % defaultSplitSize
	if num2 == 0 {
		goroutines = num1
	} else {
		goroutines = num1 + 1
	}
	return goroutines
}

// 根据数据生成Keys
func generateKeys(numbers int) []string {
	var list [][]int
	var result []string
	generate(numbers, 1, nil, &list)
	if len(list) > 0 {
		for _, v := range list {
			str := util.ArrToString(v, "_")
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
		if util.ContainsInt(list, i) {
			continue
		}
		list2 := list[0:]
		list2 = append(list2, i)
		generate(numbers-1, i+1, list2, result)
	}

}
