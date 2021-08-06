package internal

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *page) homePage(req fastrex.Request, res fastrex.Response) {
	c, err := req.Cookie("__session")
	if err != nil {
		res.Redirect("/", 302)
		return
	}
	sessionByte, err := base64.StdEncoding.DecodeString(c.GetValue())
	if err != nil {
		log.Printf("error:homePage:base64.StdEncoding.DecodeString: %v", err.Error())
		createResponsePage(res, "Response", err.Error(), "")
		return
	}
	sessionID := string(sessionByte)
	userID, err := p.db.getUserIDWithSession(req.Context(), string(sessionID))
	if err != nil {
		log.Printf("error:homePage:getUserIDWithSession: %v", err.Error())
		return
	}

	user, err := p.db.getUserDetailByID(req.Context(), userID)
	if err != nil {
		log.Printf("error:homePage:getUserDetailByID: %v", err.Error())
		createResponsePage(res, "Response", err.Error(), "")
	}
	initial := user.Name[0:1]
	// email := user.Email
	data := struct {
		Title  string
		Email  string
		Name   string
		Date   string
		Domain string
	}{"Home", initial, user.Name, time.Now().Local().Format("2 January 2006"), DOMAIN}
	res.Render("home", data)
}
