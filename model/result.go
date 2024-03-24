package model

type MatchItem struct {
	Key   string
	Count int
	Codes []string
}

type MatchResult struct {
	Query *Query
	Items *[]MatchItem
}

type MatchRangeItem struct {
	Blue       string
	RangeCount *[]int
}
type MatchRangeResult struct {
	Query    *Query
	RangeKey []string
	Items    *[]MatchRangeItem
}
