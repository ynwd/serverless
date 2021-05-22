package function

import (
	"net/http"

	"github.com/fastrodev/router"
	"github.com/fastrodev/router/web"
)

func createRouter() router.Router {
	app := router.New()
	app.Get("/", HelloHandler)
	return app
}

func HelloHandler(req web.Request, res web.Response) {
	res.Send("Hello")
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	createRouter().ServeHTTP(w, r)
}
