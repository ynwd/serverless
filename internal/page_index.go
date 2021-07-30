package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) idxPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
	email := ""
	if user != nil {
		email = user.Email
	}

	data := struct {
		Email       string
		Title       string
		Description string
		Domain      string
	}{email, TITLE, DESC, DOMAIN}
	err := res.Render(data)
	if err != nil {
		log.Panic(err.Error())
	}
}
