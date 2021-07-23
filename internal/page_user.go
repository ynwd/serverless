package internal

import (
	"strings"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) userPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("id")
	if params[0] == "home" {
		p.homePage(req, res)
		return
	}

	if params[0] == "signout" {
		p.signOut(req, res)
		return
	}

	if params[0] == "signin" {
		p.signinPage(req, res)
		return
	}

	if params[0] == "signup" {
		p.signupPage(req, res)
		return
	}

	if params[0] == "post" {
		p.createPostPage(req, res)
		return
	}
	user, _ := p.getUserFromSession(req, res)
	email := user.Email
	title := strings.Title(params[0])

	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, title, DOMAIN}
	res.Render(data)
}
