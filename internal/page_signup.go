package internal

import "github.com/fastrodev/fastrex"

func (p *page) signupPage(req fastrex.Request, res fastrex.Response) {
	data := struct {
		Title  string
		Domain string
	}{"Daftar", DOMAIN}
	res.Render("signup", data)
}
