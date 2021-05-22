package function

import (
	std "net/http"

	"github.com/fastrodev/router"
	"github.com/fastrodev/router/http"
)

func createRouter() router.Router {
	app := router.New()
	app.Get("/", func(req http.Request, res http.Response) {
		res.Send("Hello")
	})
	return app
}

func HelloHTTP(w std.ResponseWriter, r *std.Request) {
	createRouter().ServeHTTP(w, r)
}
