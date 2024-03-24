package service

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dylan_whynot/whynot_ssq/model"

	"github.com/dylan_whynot/whynot_ssq/util"
)

const defaultSplitSize = 5000

// MatchCount 满足条件数据
// 根据蓝色球号码统计那些组合出现的次数最多
func MatchCount(query *model.Query, condition *model.Condition) *model.MatchResult {
	var ssqs []model.Ssq
	match(query, func(index int, ssq model.Ssq) {
		ssqs = append(ssqs, ssq)
	})
	if len(ssqs) == 0 {
		return &model.MatchResult{Query: query}
	}
	stringkeys := util.GenerateKeys(condition.BallColor, condition.RedCount)

	var items *[]model.MatchItem
	if "blue" == strings.ToLower(condition.BallColor) {
		// 统计蓝色球
		items = matchBlueBall(&stringkeys, ssqs)
	} else {
		items = matchRedBall(&stringkeys, ssqs)
	}
	return &model.MatchResult{Query: query, Items: items}
}

func PrintMatchCountResult(query *model.Query, condition *model.Condition, print *model.PrintControl, result *model.MatchResult) {
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
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})
	fmt.Println("########统计出现次数###", now)

	fmt.Printf("统计条件蓝球%s 星期[%s] 选择年份[%s,%s), 统计颜色[%s] [%d]红球组合 出现频率大于[%d]\n", query.Blue, week, startYear, endYear, condition.BallColor, condition.RedCount, print.GranterThan)
	fmt.Println("总结果数:", len(items))
	if print.PrintIssues {
		fmt.Println("序号\t颜色\t号码\t频率\t期号")
	} else {
		fmt.Println("序号\t颜色\t号码\t频率")
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
			fmt.Println(i+1, "\t", condition.BallColor, "\t", v.Key, "\t", v.Count, "\t", v.Codes)
		} else {
			fmt.Println(i+1, "\t", condition.BallColor, "\t", v.Key, "\t", v.Count)
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
