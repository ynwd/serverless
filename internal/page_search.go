package internal

import (
	"log"
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) searchPage(req fastrex.Request, res fastrex.Response) {
	query := req.FormValue("query")
	res.Redirect("/search?q="+query, REDIRECT_CODE)
}

func (p *pageService) queryPage(req fastrex.Request, res fastrex.Response) {
	q := req.URL.Query()
	query := q.Get("q")
	t := query
	user, _ := p.getUserFromSession(req, res)
	email := ""
	if user != nil {
		email = user.Email
	}

	topic := strings.Title(t)
	td := createData(topic)
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	desc := "Hasil pencarian berdasarkan topic: " + topic
	data := struct {
		Email       string
		Title       string
		Data        []FlatPost
		Description string
		Date        string
		Domain      string
	}{email, topic, td, desc, date.Format("2 January 2006"), DOMAIN}

	err := res.Render("result", data)
	if err != nil {
		log.Println(err)
	}
}
