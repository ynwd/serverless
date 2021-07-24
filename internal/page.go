package internal

import (
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
