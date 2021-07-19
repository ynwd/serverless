package internal

import (
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
	"github.com/leekchan/accounting"
)

type pageService struct {
	db database
}

func (p *pageService) idxPage(req fastrex.Request, res fastrex.Response) {
	c, _ := req.Cookie("__session")
	userID := c.GetValue()
	user, _ := p.db.getUserDetailByID(req.Context(), userID)
	email := user.Email

	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, "Iklan Baris", DOMAIN}
	res.Render(data)
}

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

	c, _ := req.Cookie("__session")
	userID := c.GetValue()
	user, _ := p.db.getUserDetailByID(req.Context(), userID)
	email := user.Email
	title := strings.Title(params[0])

	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, title, DOMAIN}
	res.Render(data)
}

func (p *pageService) topicPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("topic")
	c, _ := req.Cookie("__session")
	userID := c.GetValue()
	user, _ := p.db.getUserDetailByID(req.Context(), userID)
	email := user.Email
	topic := strings.Title(params[0])
	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, topic, DOMAIN}
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
		Title  string
		Email  string
		Date   string
		Domain string
	}{"Home", email, time.Now().Local().Format("2 January 2006"), DOMAIN}
	res.Render("home", data)
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

func (p *pageService) signOut(req fastrex.Request, res fastrex.Response) {
	cookie, err := req.Cookie("__session")
	if err == nil {
		res.ClearCookie(cookie)
		res.Redirect("/", 302)
	}
}

func (p *pageService) detailPage(req fastrex.Request, res fastrex.Response) {
	id := ""
	params := req.Params("id")
	if len(params) > 0 {
		id = params[0]
	}
	c, _ := req.Cookie("__session")

	post, err := p.db.getPostDetail(req.Context(), id)
	if err != nil {
		msg := err.Error()
		createResponsePage("Response", msg, "", res)
		return
	}

	userDetail, _ := p.db.getUserDetailByID(req.Context(), post.User)

	file := ""
	video := ""
	title := post.Title
	date := post.Created.Format("2 January 2006")
	topic := post.Topic
	content := post.Content
	email := post.Email
	phone := post.Phone
	address := post.Address
	user := userDetail.Name
	username := userDetail.Username
	if post.Video != "" {
		s := strings.Split(post.Video, "=")
		video = "https://www.youtube.com/embed/" + s[1] + "?autoplay=1&mute=1"
	}

	if user == "" {
		user = "guest"
		username = "guest"
	}

	if post.File != "" {
		file = post.File
	}
	d := strings.Title(topic)
	d = d + " di " + strings.Title(address) + ": "
	d = d + strings.Title(title) + ". "
	d = d + content
	d = d + " | " + DOMAIN

	addr := ""
	for idx, v := range strings.Split(address, " ") {
		if idx == 0 {
			addr = addr + v
		} else {
			addr = addr + "+" + v
		}
	}

	ac := accounting.Accounting{Symbol: "Rp", Precision: 2}
	data := struct {
		Description string
		Title       string
		Topic       string
		File        string
		Date        string
		Content     string
		Email       string
		Phone       string
		Address     string
		Map         string
		UserEmail   string
		ID          string
		User        string
		Price       string
		Video       string
		Username    string
	}{d, title, topic, file, date, content, email, phone, address, addr, c.GetValue(), id, user, ac.FormatMoney(post.Price), video, username}
	res.Render("detail", data)
}

func (p *pageService) createPostPage(req fastrex.Request, res fastrex.Response) {
	c, err := req.Cookie("__session")
	userDetail, _ := p.db.getUserDetailByID(req.Context(), c.GetValue())

	if err == nil {
		data := struct {
			User   string
			Email  string
			Title  string
			Domain string
		}{userDetail.ID, userDetail.Email, "Pasang Iklan", DOMAIN}
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
