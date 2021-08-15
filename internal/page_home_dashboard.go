package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) homeDashboardPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
	if user == nil {
		createResponsePage(res, "user not found", "user not found", "/")
		return
	}
	name := user.Name
	td := createDataByUsername(user.Username)

	initial := user.Name[0:1]
	data := struct {
		Initial string
		Name    string
		Title   string
		Data    []FlatPost
	}{
		initial, name, "Dashboard", td,
	}

	err := res.Render("home_dashboard", data)
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homePostPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)

	if user == nil {
		createResponsePage(res, "user not found", "user not found", "/")
		return
	}

	initial := user.Name[0:1]
	data := struct {
		Initial     string
		Name        string
		Title       string
		Email       string
		User        string
		PostTitle   string
		PostTopic   string
		PostAddress string
		PostPrice   string
		PostPhone   string
		PostVideo   string
		PostContent []byte
		PostID      string
	}{
		initial, user.Name, "New Post", user.Email, user.ID,
		"",
		"",
		"",
		"",
		"",
		"",
		[]byte(""),
		"",
	}

	err := res.Render("home_post", data)
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeUpdatePost(req fastrex.Request, res fastrex.Response) {
	params := req.Params("id")
	id := ""

	if len(params) > 0 {
		id = params[0]
	}
	user, _ := p.getUserFromSession(req, res)

	if user == nil {
		createResponsePage(res, "user not found", "user not found", "/")
		return
	}

	post, err := p.svc.getPostDetail(req.Context(), id)
	if err != nil {
		log.Println(err.Error())
	}

	initial := user.Name[0:1]
	data := struct {
		Initial     string
		Name        string
		Title       string
		Email       string
		User        string
		PostTitle   string
		PostTopic   string
		PostAddress string
		PostPrice   int64
		PostPhone   string
		PostVideo   string
		PostContent []byte
		PostID      string
	}{
		initial,
		user.Name,
		"Update Post",
		user.Email,
		user.ID,
		post.Title,
		post.Topic,
		post.Address,
		post.Price,
		post.Phone,
		post.Video,
		[]byte(post.Content),
		post.ID,
	}

	err = res.Render("home_post", data)
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeTopicPage(req fastrex.Request, res fastrex.Response) {
	data := struct {
		Initial string
		Name    string
		Title   string
		Content string
	}{
		"T", "Testing User", "Topic", "<h1>KOntent</h1>",
	}

	err := res.Render("home_topic", data)
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeAccountPage(req fastrex.Request, res fastrex.Response) {
	data := struct {
		Initial string
		Name    string
		Title   string
		Content string
	}{
		"T", "Testing User", "Account", "<h1>KOntent</h1>",
	}

	err := res.Render("home_account", data)
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeSettingPage(req fastrex.Request, res fastrex.Response) {
	data := struct {
		Initial string
		Name    string
		Title   string
		Content string
	}{
		"T", "Testing User", "Setting", "<h1>KOntent</h1>",
	}

	err := res.Render("home_setting", data)
	if err != nil {
		log.Println(err.Error())
	}
}
