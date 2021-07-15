package internal

import "time"

type Post struct {
	ID      string    `json:"id"`
	User    string    `json:"user"`
	Topic   string    `json:"topic"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Created time.Time `json:"created,omitempty"`
}

type User struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Data struct {
	Topic string `json:"topic"`
	Posts []Post `json:"posts"`
}
