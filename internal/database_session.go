package internal

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

func (d *client) createSession(ctx context.Context, userID string, userAgent string, ip string) string {
	sessionID := uuid.New().String()
	if userID == "" {
		log.Printf("createSession error: %v", errors.New("createSession error: userID empty"))
		return ""
	}

	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	data := make(map[string]interface{})
	data["created"] = date
	data["sessionID"] = sessionID
	data["userID"] = userID
	data["userAgent"] = userAgent
	data["ip"] = ip
	d.add(ctx, "session", data)
	return sessionID
}
