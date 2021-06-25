package internal

import (
	"github.com/fastrodev/fastrex"
)

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func CreateApp(gcp bool) fastrex.App {
	app := fastrex.New()
	app.Static("static")
	app.Get("/api", handler)
	return app
}
