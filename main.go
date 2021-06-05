package function

import (
	"net/http"

	"github.com/fastrodev/rider"
)

func handler(req rider.Request, res rider.Response) {
	res.Send("Hello, cloud!")
}

func createRouter() rider.Router {
	app := rider.New()
	app.Get("/", handler)
	return app
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	createRouter().ServeHTTP(w, r)
}
