package service

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/dylan_whynot/whynot_ssq/db"
	"github.com/dylan_whynot/whynot_ssq/model"
	"github.com/dylan_whynot/whynot_ssq/util"
)

// MatchNext 查找匹配的数据
// query: 查询搜索条件
// ballColor : 匹配球的颜色
// redCount: 如果是红球 匹配红球的个数
func MatchNext(query *model.Query, condition *model.Condition) *model.MatchResult {
	// 1.查询满足条件的数据
	var matched []model.Ssq
	match(query, func(index int, ssq model.Ssq) {
		if index == 0 {
			return
		}
		value := db.DATA_POOL[index-1]
		matched = append(matched, value)
	})
	if len(matched) == 0 {
		return &model.MatchResult{}
	}
	// 生成KEY
	generatekeys := util.GenerateKeys(condition.BallColor, condition.RedCount)
	var items *[]model.MatchItem
	if "blue" == strings.ToLower(condition.BallColor) {
		// 统计蓝色球
		items = matchBlueBall(&generatekeys, matched)
	} else {
		items = matchRedBall(&generatekeys, matched)
	}
	return &model.MatchResult{Query: query, Items: items}
}

// PrintMatchNextResult 打印结果集合
func PrintMatchNextResult(query *model.Query, condition *model.Condition, print *model.PrintControl, result *model.MatchResult) {
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
	fmt.Println("########统计下一期数值出现频率###", now)
	fmt.Printf("条件蓝球%s 星期[%s] 选择年份[%s,%s)  统计颜色[%s] [%d]红球组合 统计次数大于[%d]\n", query.Blue, week, startYear, endYear, condition.BallColor, condition.RedCount, print.GranterThan)
	fmt.Println("总结果数:", len(items))
	if print.PrintIssues {
		fmt.Println("序号\t颜色\t号码\t频率\t期号")
	} else {
		fmt.Println("序号\t颜色\t号码\t频率")
	}
	fmt.Println("--------------------------------------------------")
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
