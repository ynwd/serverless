package internal

import (
	"time"

	"github.com/fastrodev/fastrex"
)

type pageService struct {
	db database
}

func (p *pageService) rootPage(req fastrex.Request, res fastrex.Response) {
	c, _ := req.Cookie("__session")
	email := c.GetValue()
	data := struct {
		Email string
	}{email}
	res.Render(data)
}

func (p *pageService) homePage(req fastrex.Request, res fastrex.Response) {
	c, err := req.Cookie("__session")
	if err != nil {
		res.Redirect("/", 302)
		return
	}

	email := c.GetValue()
	data := struct {
		Title string
		Email string
		Date  string
	}{"home", email, time.Now().Local().Format("2 January 2006")}
	res.Render("home", data)
}

func (p *pageService) arsipPage(req fastrex.Request, res fastrex.Response) {
	res.Render("arsip", nil)
}

func (p *pageService) signinPage(req fastrex.Request, res fastrex.Response) {
	post := req.URL.Query().Get("post")
	data := struct {
		Post  string
		Title string
	}{post, "Masuk"}
	res.Render("signin", data)
}

func (p *pageService) signupPage(req fastrex.Request, res fastrex.Response) {
	res.Render("signup", nil)
}

func (p *pageService) signOut(req fastrex.Request, res fastrex.Response) {
	cookie, err := req.Cookie("__session")
	if err == nil {
		res.ClearCookie(cookie)
		res.Redirect("/", 302)
	}
}

func (p *pageService) membershipPage(req fastrex.Request, res fastrex.Response) {
	res.Render("membership", nil)
}

func (p *pageService) detailPage(req fastrex.Request, res fastrex.Response) {
	id := req.Params("id")
	c, _ := req.Cookie("__session")

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
	user := post.User

	if user == "user" {
		user = "anonim"
	}

	data := struct {
		Title     string
		Topic     string
		Date      string
		Content   string
		Email     string
		Phone     string
		Address   string
		UserEmail string
		ID        string
		User      string
	}{title, topic, date, content, email, phone, address, c.GetValue(), id[0], user}
	res.Render("detail", data)
}

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	c, err := req.Cookie("__session")
	if err == nil {
		data := struct {
			User  string
			Title string
		}{c.GetValue(), "Pasang Iklan"}
		res.Render("create", data)
		return
	}
	res.Render("create", nil)
}
