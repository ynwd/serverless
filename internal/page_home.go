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
	userID, err := p.svc.getUserIDWithSession(req.Context(), string(sessionID))
	if err != nil {
		log.Printf("error:homePage:getUserIDWithSession: %v", err.Error())
		return
	}

	user, err := p.svc.getUserDetailByID(req.Context(), userID)
	if err != nil {
		log.Printf("error:homePage:getUserDetailByID: %v", err.Error())
		createResponsePage(res, "Response", err.Error(), "")
	}

	path := req.URL.Query().Get("menu")
	if path == "" {
		path = "search"
	}

	initial := user.Name[0:1]
	userEmail := user.Email
	data := struct {
		Path      string
		Initial   string
		UserEmail string
		Title     string
		Email     string
		Name      string
		Date      string
		Domain    string
	}{path, initial, userEmail, "Home", initial, user.Name, time.Now().Local().Format("2 January 2006"), DOMAIN}
	err = res.Render("home", data)
	if err != nil {
		log.Println(err.Error())
	}
}
