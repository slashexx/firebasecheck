package models

type BlogPost struct {
	ID      string `json:"id"`
	Headers string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
