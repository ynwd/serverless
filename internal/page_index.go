package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) idxPage(req fastrex.Request, res fastrex.Response) {
	user, err := p.getUserFromSession(req, res)
	if err != nil {
		log.Printf("idxPage:getUserFromSession: %v", err)
	}
	email := ""
	if user != nil {
		email = user.Email
	}

	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, TITLE, DOMAIN}
	res.Render(data)
}
