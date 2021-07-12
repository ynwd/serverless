package internal

import "time"

type Post struct {
	ID      string    `json:"id"`
	User    string    `json:"user"`
	Topic   string    `json:"topic"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created,omitempty"`
}

type Data struct {
	Topic string `json:"topic"`
	Posts []Post `json:"posts"`
}
