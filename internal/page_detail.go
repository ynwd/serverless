package internal

import (
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/fastrodev/fastrex"
	"github.com/leekchan/accounting"
)

func (p *page) detailPage(req fastrex.Request, res fastrex.Response) {
	id := ""
	params := req.Params("id")
	if len(params) > 0 {
		id = params[0]
	}
	c, _ := req.Cookie("__session")

	post, err := p.svc.getPostDetail(req.Context(), id)
	if err != nil {
		msg := err.Error()
		createResponsePage(res, "Response", msg, "")
		return
	}

	userDetail, _ := p.svc.getUserDetailByID(req.Context(), post.User)

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
	view := post.View
	if post.Video != "" {
		s := strings.Split(post.Video, "=")
		video = "https://www.youtube.com/embed/" + s[1] + "?autoplay=1&mute=1"
	}

	if user == "" {
		user = "guest"
		username = "guest"
	}

	usr, _ := p.getUserFromSession(req, res)
	initial := ""
	userEmail := ""
	if usr != nil {
		initial = usr.Username[0:1]
		userEmail = usr.Email
	}

	if post.File != "" {
		file = post.File
	}
	d := strings.Title(topic)
	if address != "" {
		d = d + " di " + strings.Title(address) + ": "
	} else {
		d = d + ": "
	}
	d = d + content
	d = d + " | " + DOMAIN

	mapAddr := ""
	tmpAddr := address
	if strings.Contains(address, ",") {
		tmpAddr = strings.ReplaceAll(address, ",", "")
	}
	for idx, v := range strings.Split(tmpAddr, " ") {
		if idx == 0 {
			mapAddr = mapAddr + v
		} else {
			mapAddr = mapAddr + "+" + v
		}
	}

	ac := accounting.Accounting{Symbol: "Rp", Precision: 2}
	data := struct {
		Initial     string
		Description string
		Title       string
		Topic       string
		File        string
		Date        string
		Content     string
		Email       string
		UserEmail   string
		Phone       string
		Address     string
		Map         string
		Cookie      string
		ID          string
		User        string
		Price       string
		Video       string
		Username    string
		Domain      string
	}{initial, d, title, topic, file, date, content, email, userEmail, phone, address, mapAddr, c.GetValue(), id, user, ac.FormatMoney(post.Price), video, username, DOMAIN}

	err = res.Render("detail", data)

	if err != nil {
		log.Println(err.Error())
	} else {
		go p.updateView(req, id, view)
	}
}

func (p *page) updateView(req fastrex.Request, id string, view int64) {
	_, err := p.svc.update(req.Context(), &Query{
		Collection: "post",
		Field:      "id",
		Op:         "==",
		Value:      id,
		OrderBy:    "created",
	}, []firestore.Update{
		{
			Path:  "view",
			Value: view + 1,
		},
	})

	if err != nil {
		log.Println(err.Error())
	}
}
