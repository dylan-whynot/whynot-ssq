package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dylan_whynot/whynot_ssq/model"
)

// 定义双色球基本类信息

// 定义数据
var DATA_POOL []model.Ssq

func init() {
	fmt.Println("db init")
}
func LoadDatas() {
	bytes, err := os.ReadFile("data/ssq.json")
	if err != nil {
		log.Fatalln("加载数据异常", err)
		return
	}
	var datas []model.Ssq
	err = json.Unmarshal(bytes, &datas)
	if err != nil {
		log.Fatalln("解析json异常", err)
		return
	}
	for _, v := range datas {
		DATA_POOL = append(DATA_POOL, v)
	}
	fmt.Println("已经加载 ", len(DATA_POOL), "条数据")
}
