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
	Price   int64     `json:"price"`
	Video   string    `json:"video"`
	View    int64     `json:"view"`
	Created time.Time `json:"created,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Data struct {
	Topic string `json:"topic"`
	Posts []Post `json:"posts"`
}

type Session struct {
	SessionID string    `json:"sessionID"`
	UserID    string    `json:"userID"`
	Created   time.Time `json:"created,omitempty"`
	Device    string
}
