package internal

import (
	"encoding/base64"
	"log"

	"github.com/fastrodev/fastrex"
)

type pageService struct {
	db database
}

func (p *pageService) signinPage(req fastrex.Request, res fastrex.Response) {
	post := req.URL.Query().Get("post")
	data := struct {
		Post   string
		Title  string
		Domain string
	}{post, "Masuk", DOMAIN}
	res.Render("signin", data)
}

func (p *pageService) signupPage(req fastrex.Request, res fastrex.Response) {
	data := struct {
		Title  string
		Domain string
	}{"Daftar", DOMAIN}
	res.Render("signup", data)
}

func (p *pageService) getUserFromSession(req fastrex.Request, res fastrex.Response) (*User, error) {
	c, _ := req.Cookie("__session")
	sessionByte, err := base64.StdEncoding.DecodeString(c.GetValue())
	if err != nil {
		log.Printf("getUserFromSession:base64.StdEncoding.DecodeString: %v", err.Error())
	}
	userAgent := req.UserAgent()
	sessionID := string(sessionByte)
	userID, err := p.db.getUserIDWithSession(req.Context(), string(sessionID), userAgent)
	if err != nil {
		log.Printf("getUserFromSession:getUserIDWithSession: %v", err.Error())
	}

	user, err := p.db.getUserDetailByID(req.Context(), userID)
	if err != nil {
		log.Printf("getUserFromSession:getUserDetailByID: %v", err.Error())
		return nil, err
	}

	return &user, nil
}
