package internal

import (
	"github.com/fastrodev/fastrex"
)

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
	if user != nil {
		data := struct {
			User   string
			Email  string
			Title  string
			Domain string
		}{user.ID, user.Email, "Pasang Iklan", DOMAIN}
		res.Render("create", data)
		return
	}
	guest := struct {
		User   string
		Email  string
		Title  string
		Domain string
	}{"", "", "Pasang Iklan", DOMAIN}
	res.Render("create", guest)
}
