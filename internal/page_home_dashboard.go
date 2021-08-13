package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) homeDashboardPage(req fastrex.Request, res fastrex.Response) {
	err := res.Render("home_dashboard", "")
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homePostPage(req fastrex.Request, res fastrex.Response) {
	err := res.Render("home_post", "")
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeTopicPage(req fastrex.Request, res fastrex.Response) {
	err := res.Render("home_topic", "")
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeAccountPage(req fastrex.Request, res fastrex.Response) {
	err := res.Render("home_account", "")
	if err != nil {
		log.Println(err.Error())
	}
}

func (p *page) homeSettingPage(req fastrex.Request, res fastrex.Response) {
	err := res.Render("home_setting", "")
	if err != nil {
		log.Println(err.Error())
	}
}
