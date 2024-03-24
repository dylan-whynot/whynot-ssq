package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dylan_whynot/whynot_ssq/model"
)

// 定义返回结果数据
type Response struct {
	State    int      `json:"state"`
	Message  string   `json:"message"`
	Total    int      `json:"total"`
	PageNum  int      `json:"pageNum"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Tflag    int      `json:"Tflag"`
	Result   []Detail `json:"result"`
}
type Detail struct {
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	DetailsLink string        `json:"detailsLink"`
	VideoLink   string        `json:"videoLink"`
	Date        string        `json:"date"`
	Week        string        `json:"week"`
	Red         string        `json:"red"`
	Blue        string        `json:"blue"`
	Blue2       string        `json:"blue_2"`
	Sales       string        `json:"sales"`
	Poolmoney   string        `json:"poolmoney"`
	Content     string        `json:"content"`
	Addmoney    string        `json:"addmoney"`
	Addmoney2   string        `json:"addmoney2"`
	Msg         string        `json:"msg"`
	Z2add       string        `json:"z2add"`
	M2add       string        `json:"m2add"`
	Prizegrades []Prizegrades `json:"prizegrades"`
}
type Prizegrades struct {
	Type      int    `json:"type"`
	Typenum   string `json:"typenum"`
	Typemoney string `json:"typemoney"`
}

var historyDatas []model.Ssq

const fileName = "ssq.json"

func loadHistoryDatas() {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalln("open file [ssq.json]", err)
	}
	defer file.Close()
	byte, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("read file [ssq.json]", err)
	}
	if len(byte) > 0 {
		err := json.Unmarshal(byte, &historyDatas)
		if err != nil {
			log.Fatalln("json.unmarshal error", err)
		}
	}
}
func getMaxCodeId() string {
	maxCodeId := "0"
	for _, v := range historyDatas {
		if v.Id > maxCodeId {
			maxCodeId = v.Id
		}
	}
	return maxCodeId
}

var cookie string

// 更新双色球历史数据信息
func Update(incookie string) {
	cookie = incookie
	log.Println("开始更新")
	loadHistoryDatas()
	maxCodeId := getMaxCodeId()
	var container []model.Ssq
	pageNumber := 1
	num, result := crawlerPage(pageNumber)
	over := addDatas(maxCodeId, result, &container)
	if !over {
		for pageNumber = 2; pageNumber <= num; pageNumber++ {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Millisecond)
			_, result2 := crawlerPage(pageNumber)
			over = addDatas(maxCodeId, result2, &container)
			if over {
				break
			}
		}
	}
	// 写入文件
	writeToFile(container)

}
func addDatas(maxCodeId string, list []model.Ssq, container *[]model.Ssq) (over bool) {
	for _, v := range list {
		if v.Id > maxCodeId {
			*container = append(*container, v)
		} else {
			return true
		}
	}
	return false
}
func writeToFile(list []model.Ssq) {
	if len(list) == 0 {
		log.Println("已经是最新")
		return
	}
	count := len(list)
	marshal, err := json.Marshal(list)
	if err != nil {
		log.Fatalln(err)
	}
	var buffer bytes.Buffer
	err = json.Compact(&buffer, marshal)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(fileName, buffer.Bytes(), 0666)
	if err != nil {
		log.Fatalln("写入文件失败", err)
	}
	log.Println("更新完成 ", count, "条数据")
}

var baseUrl = "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=&dayEnd=&pageNo=@@&pageSize=30&week=&systemType=PC"
var placeholder = "@@"

// 根据页码爬取数据
func crawlerPage(pageNumber int) (pageNum int, pageResult []model.Ssq) {
	log.Println("crawler page num ", pageNumber)
	newUrl := strings.Replace(baseUrl, placeholder, strconv.Itoa(pageNumber), 1)
	response, err := sendRequest(newUrl)
	if err != nil {
		log.Fatalln(err)
	}
	ssqs := parserJson(response)
	return response.PageNum, ssqs
}
func sendRequest(url string) (Response, error) {
	client := http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Cookie", cookie)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	request.Header.Add("Host", "www.cwl.gov.cn")
	request.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
		//return nil,errors.New("异常")
	}
	if resp.StatusCode == 200 {
		// 正确
		defer resp.Body.Close()
		bytes, _ := io.ReadAll(resp.Body)
		response := Response{}
		err := json.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatalln(err)
		}
		return response, nil
	} else {
		//异常
		log.Fatalln("response code", resp.StatusCode)
	}
	return Response{}, errors.New("异常")
}
func parserJson(response Response) []model.Ssq {
	var ssqs []model.Ssq
	if "查询成功" == response.Message {
		details := response.Result
		for _, v := range details {
			ssq := model.Ssq{}
			ssq.Id = v.Code
			ssq.Week = v.Week
			index := strings.Index(v.Date, "(")
			ssq.Date = v.Date[0:index]
			redArr := strings.Split(v.Red, ",")
			ssq.RedNumbers = redArr
			ssq.Blue = v.Blue
			ssq.Sales = v.Sales
			ssq.PoolAmount = v.Poolmoney
			prizes := []model.Prize{}
			for _, p := range v.Prizegrades {
				if p.Typenum != "" {
					prize := model.Prize{Code: v.Code, Number: p.Type, PeopleNumber: p.Typenum, Money: p.Typemoney}
					prizes = append(prizes, prize)
				}

			}
			ssq.Prizegrades = prizes
			ssqs = append(ssqs, ssq)
		}
	}
	return ssqs
}
