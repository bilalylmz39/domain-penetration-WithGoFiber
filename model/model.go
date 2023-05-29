package model

type DomainResult struct {
	StatusCode      int     `json:"status_code"`
	ErrorCode       string  `json:"error"`
	PageRankInteger int     `json:"page_rank_integer"`
	PageRankDecimal float64 `json:"page_rank_decimal"`
	Rank            string  `json:"rank"`
	Domain          string  `json:"domain"`
}

type Data struct {
	StatusCode  int            `json:"status_code"`
	Response    []DomainResult `json:"response"`
	LastUpdated string         `json:"last_updated"`
}

const OpenPageRankURL = "https://openpagerank.com/api/v1.0/getPageRank"
