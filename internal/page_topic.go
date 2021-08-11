package internal

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *page) topicPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("topic")
	t := ""
	if len(params) > 0 {
		t = params[0]
	}

	user, _ := p.getUserFromSession(req, res)
	email := ""
	initial := ""
	if user != nil {
		email = user.Email
		initial = user.Username[0:1]
	}

	topic := strings.Title(t)
	td := createData(topic)
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	desc := "Hasil pencarian berdasarkan topic: " + topic
	data := struct {
		UserEmail   string
		Initial     string
		Title       string
		Data        []FlatPost
		Description string
		Date        string
		Domain      string
	}{email, initial, topic, td, desc, date.Format("2 January 2006"), DOMAIN}
	err := res.Render("result", data)
	if err != nil {
		log.Println(err)
	}
}

type FlatPost struct {
	Header string
	Post
}

func createData(topic string) []FlatPost {
	t := strings.ToLower(topic)
	body := createJsonPost(t)
	data := []FlatPost{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Printf("createData %v:", errUnmarshal.Error())
	}
	return data
}

func createJsonPost(topic string) []byte {
	d := ReadPostByTopic(topic)
	output, errMarshal := json.Marshal(groupByTopic(d))
	if errMarshal != nil {
		panic(errMarshal)
	}

	return output
}
