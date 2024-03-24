package service

import (
	"log"
	"strings"
	"sync"

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

func matchRedBall(generatekeys *[]string, matched []model.Ssq) *[]model.MatchItem {
	goroutines := goroutinesSize(len(*generatekeys))
	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutines)
	var list = make([][]model.MatchItem, goroutines)
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
			var result = make(map[string]*model.MatchItem)
			for _, key := range subArr {
				for _, ssq := range matched {
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
							result[key] = &model.MatchItem{Key: key, Count: 1, Codes: []string{ssq.Id}}
						}
					}
				}
			}
			var l = make([]model.MatchItem, len(result))
			index := 0
			for _, v := range result {
				l[index] = *v
				index++
			}
			list[i] = l
			log.Println("goroutines [", i, "] finished")
		}(generatekeys, i)
	}
	waitGroup.Wait()
	// 合并结果集合
	var items []model.MatchItem
	for _, m := range list {
		items = append(items, m...)
	}
	return &items
}

func matchBlueBall(generatekeys *[]string, matched []model.Ssq) *[]model.MatchItem {
	var result = make(map[string]*model.MatchItem)
	for _, key := range *generatekeys {
		for _, ssq := range matched {

			if ssq.Blue == key {
				v, ok := result[key]
				if ok {
					v.Count = v.Count + 1
					v.Codes = append(v.Codes, ssq.Id)
				} else {
					result[key] = &model.MatchItem{Key: key, Count: 1, Codes: []string{ssq.Id}}
				}
			}
		}
	}
	var l = make([]model.MatchItem, len(result))
	index := 0
	for _, v := range result {
		l[index] = *v
		index++
	}
	return &l
}
