package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) idxPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
	email := ""
	initial := ""
	if user != nil {
		email = user.Email
		initial = user.Username[0:1]
	}

	data := struct {
		Initial     string
		Email       string
		Title       string
		Description string
		Domain      string
		Path        string
	}{initial, email, TITLE, DESC, DOMAIN, req.URL.Path}
	err := res.Render(data)
	if err != nil {
		log.Panic(err.Error())
	}
}
