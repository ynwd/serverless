package serverless

import (
	"net/http"

	"github.com/fastrodev/rider"
)

func htmlHandler(req rider.Request, res rider.Response) {
	data := map[string]interface{}{
		"title": "Cloud function app",
		"name":  "Agus",
	}
	res.Render(data)
}

func apiHandler(req rider.Request, res rider.Response) {
	res.Json(`{"message":"ok"}`)
}

func createApp() rider.Router {
	app := rider.New()
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

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	createApp().ServeHTTP(w, r)
}
