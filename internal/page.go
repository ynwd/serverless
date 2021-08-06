package internal

import (
	"github.com/fastrodev/fastrex"
)

type pageService struct {
	db Database
}

func createPageRoute(app fastrex.App, page *pageService) fastrex.App {
	app.Post("/", healthChk).
		Get("/", page.idxPage).
		Get("/:username", page.userPage).
		Get("/post/:id", page.detailPage).
		Get("/topic/:topic", page.topicPage).
		Get("/search", page.queryPage).
		Post("/search", page.searchPage).
		Get("/activate/:code", page.activatePage)
	return app
}
