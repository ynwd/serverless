package internal

import (
	"github.com/fastrodev/fastrex"
)

type page struct {
	db Database
}

func createPage(db Database) *page {
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
