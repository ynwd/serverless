package internal

import "github.com/fastrodev/fastrex"

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func CreateApp() fastrex.App {
	app := fastrex.New()
	app.Get("/api", handler)
	// add static folder
	app.Static("static")
	// add html templates
	app.Template("template/search.html")
	app.Template("template/signin.html")
	app.Template("template/signup.html")
	app.Template("template/membership.html")
	app.Template("template/home.html")
	app.Template("template/detail.html")
	// add routes
	app.Get("/search", searchHandler)
	app.Get("/signin", signinHandler)
	app.Get("/signup", signupHandler)
	app.Get("/membership", membershipHandler)
	app.Get("/home", homeHandler)
	app.Get("/detail/:post", detailHandler)
	return app
}
