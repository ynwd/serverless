package internal

import (
	"log"
	"strings"

	"github.com/fastrodev/fastrex"
)

func (p *page) homeDashboardPage(req fastrex.Request, res fastrex.Response) {
	user, _ := p.getUserFromSession(req, res)
	if user == nil {
		res.Redirect("/", 302)
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
		initial, name, "Posts", td,
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
		PostFile    string
	}{
		initial, user.Name, "Create New Post", user.Email, user.ID,
		"",
		"",
		"",
		"",
		"",
		"",
		[]byte(""),
		"",
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

	video := ""
	if post.Video != "" {
		s := strings.Split(post.Video, "=")
		video = "https://www.youtube.com/embed/" + s[1] + "?autoplay=1&mute=1"
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
		PostFile    string
		Video       string
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
		post.File,
		video,
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
