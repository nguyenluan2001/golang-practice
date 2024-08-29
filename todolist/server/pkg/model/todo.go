package model

type Todo struct {
	UID    string `json:"uid"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
