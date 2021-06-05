package function

import (
	"net/http"

	"github.com/fastrodev/rider"
)

func handler(req rider.Request, res rider.Response) {
	data := map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Agus",
	}
	res.Render(data)
}

func createRouter() rider.Router {
	app := rider.New()
	app.Template("index.html")
	app.Get("/", handler)
	return app
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	createRouter().ServeHTTP(w, r)
}
