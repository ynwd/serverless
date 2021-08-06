package internal

import (
	"encoding/base64"
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) signOut(req fastrex.Request, res fastrex.Response) {
	cookie, err := req.Cookie("__session")
	if err != nil {
		log.Printf("signOut:req.Cookie: %v", err.Error())
		createResponsePage(res, "Response", err.Error(), "")
		return
	}

	sessionByte, err := base64.StdEncoding.DecodeString(cookie.GetValue())
	if err != nil {
		log.Printf("signOut:base64.StdEncoding.DecodeString: %v", err.Error())
		createResponsePage(res, "Response", err.Error(), "")
		return
	}

	sessionID := string(sessionByte)

	if err == nil {
		_, err := p.db.delete(req.Context(), &Query{
			Collection: "session",
			Field:      "sessionID",
			Op:         "==",
			Value:      sessionID,
			OrderBy:    "userAgent",
		})
		if err != nil {
			log.Printf("signOut:%v", err.Error())
		}

		res.ClearCookie(cookie)
		res.Redirect("/", 302)
	}
}
