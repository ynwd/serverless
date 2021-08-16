package internal

import (
	"github.com/fastrodev/fastrex"
)

type page struct {
	svc Service
}

func createPage(db Service) *page {
	return &page{db}
}

func createPageRoute(app fastrex.App, p *page) fastrex.App {
	app.Get("/", p.idxPage).
		Get("/:username", p.userPage).
		Get("/post/:id", p.detailPage).
		Get("/topic/:topic", p.topicPage).
		Get("/search", p.queryPage).
		Post("/search", p.searchPage).
		Get("/activate/:code", p.activatePage)
	return app
}

func createHomePageRoute(app fastrex.App, p *page) fastrex.App {
	app.Get("/home", p.homePage).
		Get("/home/dashboard", p.homeDashboardPage).
		Get("/home/post", p.homePostPage).
		Get("/home/post/:id", p.homeUpdatePost).
		Get("/home/topic", p.homeTopicPage).
		Get("/home/account", p.homeAccountPage).
		Get("/home/setting", p.homeSettingPage).
		Get("/home/search", p.homeSearchPage)
	return app
}
