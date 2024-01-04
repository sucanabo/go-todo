package model

type Filter struct {
	Status      string `json:"status" form:"status"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}
