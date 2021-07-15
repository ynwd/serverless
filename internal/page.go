package internal

import (
	"github.com/fastrodev/fastrex"
)

type pageService struct {
	db database
}

func (p *pageService) homePage(req fastrex.Request, res fastrex.Response) {
	c, _ := req.Cookie("user")
	data := struct {
		Email string
	}{c.GetValue()}
	res.Render(data)
}

func (p *pageService) arsipPage(req fastrex.Request, res fastrex.Response) {
	res.Render("arsip", nil)
}

func (p *pageService) signinPage(req fastrex.Request, res fastrex.Response) {
	res.Render("signin", nil)
}

func (p *pageService) signupPage(req fastrex.Request, res fastrex.Response) {
	res.Render("signup", nil)
}

func (p *pageService) membershipPage(req fastrex.Request, res fastrex.Response) {
	res.Render("membership", nil)
}

func (p *pageService) detailPage(req fastrex.Request, res fastrex.Response) {
	id := req.Params("id")
	c, _ := req.Cookie("user")
	// cookie := struct {
	// 	Email string
	// }{c.GetValue()}

	post, err := p.db.getPostDetail(req.Context(), id[0])
	if err != nil {
		msg := err.Error()
		createResponsePage(msg, res)
		return
	}

	title := post.Title
	date := post.Created.Format("2 January 2006")
	topic := post.Topic
	content := post.Content
	email := post.Email
	phone := post.Phone
	address := post.Address

	data := struct {
		Title     string
		Topic     string
		Date      string
		Content   string
		Email     string
		Phone     string
		Address   string
		UserEmail string
	}{title, topic, date, content, email, phone, address, c.GetValue()}

	res.Render("detail", data)
}

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	res.Render("create", nil)
}
