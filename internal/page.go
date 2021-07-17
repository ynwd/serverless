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
	userID := c.GetValue()
	user, _ := p.db.getUserDetailByID(req.Context(), userID)
	email := user.Email

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

	userID := c.GetValue()
	user, _ := p.db.getUserDetailByID(req.Context(), userID)
	email := user.Email

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
		createResponsePage(msg, "/", res)
		return
	}

	userDetail, _ := p.db.getUserDetailByID(req.Context(), post.User)

	file := ""
	title := post.Title
	date := post.Created.Format("2 January 2006")
	topic := post.Topic
	content := post.Content
	email := post.Email
	phone := post.Phone
	address := post.Address
	user := userDetail.Name

	if user == "user" {
		user = "guest"
	}
	if post.File != "" {
		file = post.File
	}

	data := struct {
		Title     string
		Topic     string
		File      string
		Date      string
		Content   string
		Email     string
		Phone     string
		Address   string
		UserEmail string
		ID        string
		User      string
	}{title, topic, file, date, content, email, phone, address, c.GetValue(), id[0], user}
	res.Render("detail", data)
}

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	c, err := req.Cookie("__session")
	userDetail, _ := p.db.getUserDetailByID(req.Context(), c.GetValue())

	if err == nil {
		data := struct {
			User  string
			Email string
			Title string
		}{userDetail.ID, userDetail.Email, "Pasang Iklan | Fastro.app"}
		res.Render("create", data)
		return
	}
	res.Render("create", nil)
}
