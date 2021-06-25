package internal

import (
	"github.com/fastrodev/fastrex"
)

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func CreateApp() fastrex.App {
	root := Handler{data: readJson()}
	app := fastrex.New()
	app.Static("static")
	app.Template("template/index.html")
	app.Get("/api", handler)
	app.Get("/", root.htmlHandler)
	return app
}
