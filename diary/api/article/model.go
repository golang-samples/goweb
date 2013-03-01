package article

import (
	"time"
)

type Article struct {
	Id      string    `json:"_id" bson:"_id"`
	Date    time.Time `json:"-" bson:"-"`
	Title   string    `json:"title" bson:"title"`
	Content string    `json:"content" bson:"content"`
}

func NewArticle(title string) *Article {
	return &Article{"", time.Now(), title, ""}
}
