package model

type Book struct {
	ID      string `json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
}

type NewBook struct {
	Title  string `json:"title"`
	Content string `json:"content"`
}