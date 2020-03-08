package model

type Link struct {
	Relation string `json:"rel"`
	Method   string `json:"method"`
	Href     string `json:"href"`
}
