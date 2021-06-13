package serverless

import (
	"net/http"

	fastrex "github.com/fastrodev/fastrex"
)

func htmlHandler(req fastrex.Request, res fastrex.Response) {
	data := map[string]interface{}{
		"title": "Cloud function app",
		"name":  "Agus",
	}
	res.Render(data)
}

func apiHandler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func createApp() fastrex.App {
	app := fastrex.New()
	// setup static files
	app.Static("static")
	// setup html template
	app.Template("index.html")
	// setup route for /html
	app.Get("/html", htmlHandler)
	// setup route for /api
	app.Get("/api", apiHandler)
	return app
}

func Entrypoint(w http.ResponseWriter, r *http.Request) {
	createApp().ServeHTTP(w, r)
}
