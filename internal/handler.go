package internal

import "github.com/fastrodev/fastrex"

func searchHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("search", nil)
}

func signinHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("signin", nil)
}

func signupHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("signup", nil)
}

func homeHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("home", nil)
}

func membershipHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("membership", nil)
}

func detailHandler(req fastrex.Request, res fastrex.Response) {
	res.Render("detail", nil)
}
