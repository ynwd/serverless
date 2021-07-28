package internal

import (
	"github.com/fastrodev/fastrex"
)

func (p *pageService) idxPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
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
