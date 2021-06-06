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
	app.Template("index.html")
	app.Get("/", htmlHandler)
	app.Get("/api", apiHandler)
	return app
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	createApp().ServeHTTP(w, r)
}
