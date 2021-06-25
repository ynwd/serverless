package internal

import (
	"github.com/fastrodev/fastrex"
)

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func CreateApp(gcp bool) fastrex.App {
	// read json file
	root := Handler{serverless: gcp}
	root.data = root.readJson()

	// create fastrex instance
	app := fastrex.New()
	app.Static("static")
	app.Template("template/index.html")
	app.Get("/api", handler)
	app.Get("/", root.htmlHandler)
	return app
}
