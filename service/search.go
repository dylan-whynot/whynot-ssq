package service

import (
	"fmt"
	"time"

	"github.com/dylan_whynot/whynot_ssq/model"
)

func Search(query *model.Query) *[]model.Ssq {
	var matched []model.Ssq
	match(query, func(index int, s model.Ssq) {
		matched = append(matched, s)
	})
	return &matched
}
func PrintSearchResult(query *model.Query, print *model.PrintControl, list *[]model.Ssq) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("########输出结果###", now)
	fmt.Printf("查询条件:期号%s 星期[%s] 蓝球%s 年份区间[%s,%s)\n", query.Issues, query.Week, query.Blue, query.StartYear, query.EndYear)
	fmt.Println("序号\t期号\t日期\t星期\t红球\t蓝球\t销售额(元)\t奖金池(元)")
	fmt.Println("-----------------------------------------")
	for i, v := range *list {
		if i >= print.PageSize {
			break
		}
		fmt.Println(i+1, "\t", v.Id, "\t", v.Date, "\t", v.Week, "\t", v.RedNumbers, "\t", v.Blue, "\t", v.Sales, "\t", v.PoolAmount)
	}
	fmt.Println("########输出结果###", now)
}
