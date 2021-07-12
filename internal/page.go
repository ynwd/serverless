package internal

import "github.com/fastrodev/fastrex"

type pageService struct {
	db database
}

func (p *pageService) searchPage(req fastrex.Request, res fastrex.Response) {
	res.Render("search", nil)
}

func (p *pageService) signinPage(req fastrex.Request, res fastrex.Response) {
	res.Render("signin", nil)
}

func (p *pageService) signupPage(req fastrex.Request, res fastrex.Response) {
	res.Render("signup", nil)
}

func (p *pageService) homePage(req fastrex.Request, res fastrex.Response) {
	res.Render("home", nil)
}

func (p *pageService) membershipPage(req fastrex.Request, res fastrex.Response) {
	res.Render("membership", nil)
}

func (p *pageService) detailPage(req fastrex.Request, res fastrex.Response) {
	res.Render("detail", nil)
}

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	res.Render("detail", nil)
}
