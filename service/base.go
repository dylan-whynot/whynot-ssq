package service

import (
	"github.com/dylan_whynot/whynot_ssq/db"
	"github.com/dylan_whynot/whynot_ssq/model"
	"github.com/dylan_whynot/whynot_ssq/util"
)

// 匹配搜索到的满足条件的model.Ssq
func match(query *model.Query, matchFun func(index int, ssq model.Ssq)) {
	blues := query.Blue
	for i, v := range blues {
		if len(v) == 1 {
			blues[i] = "0" + v
		}
	}
	for i, v := range db.DATA_POOL {
		b := v.Blue
		date := v.Date
		w := v.Week
		if len(blues) > 0 && !util.ContainsString(blues, b) {
			continue
		}
		if query.Week != "" && w != query.Week {
			continue
		}
		if query.StartYear != "" && date < query.StartYear {
			continue
		}
		if query.EndYear != "" && query.EndYear < date {
			continue
		}
		if len(query.Issues) > 0 && !util.ContainsString(query.Issues, v.Id) {
			continue
		}
		// 执行自定义函数
		matchFun(i, v)
	}

}
