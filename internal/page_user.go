package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *page) userPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("username")
	param := ""
	if len(params) > 0 {
		param = params[0]
	}

	if param == "home" {
		p.homePage(req, res)
		return
	}

	if param == "signout" {
		p.signOut(req, res)
		return
	}

	if param == "signin" {
		p.signinPage(req, res)
		return
	}

	if param == "signup" {
		p.signupPage(req, res)
		return
	}

	if param == "post" {
		p.createPostPage(req, res)
		return
	}

	if param == "search" {
		p.queryPage(req, res)
		return
	}

	if param == "activate" {
		p.activatePage(req, res)
		return
	}

	user, _ := p.getUserFromSession(req, res)
	email := ""
	if user != nil {
		email = user.Email
	}
	td := createDataByUsername(param)
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	desc := fmt.Sprintf("Profile of %v", param)
	usr, _ := p.db.getUserDetailByUsername(req.Context(), param)
	name := "Guest"
	initial := ""
	if usr != nil {
		name = usr.Name
		initial = usr.Username[0:1]
	}

	data := struct {
		Initial     string
		Email       string
		Title       string
		Data        []FlatPost
		Description string
		Date        string
		Domain      string
	}{initial, email, name, td, desc, date.Format("2 January 2006"), DOMAIN}
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
