package internal

import "time"

type Post struct {
	ID      string    `json:"id"`
	User    string    `json:"user_id"`
	Topic   string    `json:"topic"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	File    string    `json:"file"`
	Content string    `json:"content"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Created time.Time `json:"created,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Data struct {
	Topic string `json:"topic"`
	Posts []Post `json:"posts"`
}
