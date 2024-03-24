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
