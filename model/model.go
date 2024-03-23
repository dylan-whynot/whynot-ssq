package model

// 定义双色球信息
type Ssq struct {
	Id          string   `json:"id"`
	Date        string   `json:"date"`
	Week        string   `json:"week"`
	RedNumbers  []string `json:"red_numbers"`
	Blue        string   `json:"blue"`
	Sales       string   `json:"sales"`
	PoolAmount  string   `json:"pool_amount"`
	Prizegrades []Prize  `json:"prizegrades"`
}

// 定义奖项信息
type Prize struct {
	Code         string `json:"code"`
	Number       int    `json:"number"`
	PeopleNumber string `json:"people_number"`
	Money        string `json:"money"`
}
type Query struct {
	// 期号
	Issues    []string
	Week      string
	Blue      []string
	StartYear string
	EndYear   string
}

// 打印控制项
type PrintControl struct {
	PageSize    int
	GranterThan int
	RedCount    int
	// 是否打印期号
	PrintIssues bool
}
