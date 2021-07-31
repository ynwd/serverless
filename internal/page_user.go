package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) userPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("username")

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

	if params[0] == "search" {
		p.topicPage(req, res)
		return
	}

	user, _ := p.getUserFromSession(req, res)
	email := ""
	if user != nil {
		email = user.Email
	}
	td := createDataByUsername(params[0])
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	desc := fmt.Sprintf("Profile of %v", params[0])
	usr, _ := p.db.getUserDetailByUsername(req.Context(), params[0])
	name := "Guest"
	if usr != nil {
		name = usr.Name
	}

	data := struct {
		Email       string
		Title       string
		Data        []FlatPost
		Description string
		Date        string
		Domain      string
	}{email, name, td, desc, date.Format("2 January 2006"), DOMAIN}
	res.Render("result", data)
}

func createDataByUsername(username string) []FlatPost {
	u := strings.ToLower(username)
	body := createJsonPostByUsername(u)
	data := []FlatPost{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Printf("createDataByUsername %v:", errUnmarshal.Error())
	}
	return data
}

func createJsonPostByUsername(username string) []byte {
	d := ReadPostByUsername(username)
	output, errMarshal := json.Marshal(groupByTopic(d))
	if errMarshal != nil {
		panic(errMarshal)
	}

	return output
}
