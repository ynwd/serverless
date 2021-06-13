package internal

import "github.com/fastrodev/fastrex"

func CreateApp() fastrex.App {
	app := fastrex.New()
	app.Static("static")
	app.Template("template/index.html")
	app.Get("/html", htmlHandler)
	app.Get("/api", handler)
	return app
}

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func htmlHandler(req fastrex.Request, res fastrex.Response) {
	data := map[string]interface{}{
		"title": "Cloud function app",
		"name":  "Agus",
	}
	res.Render(data)
}
