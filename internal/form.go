package internal

import "github.com/fastrodev/fastrex"

type form struct {
	svc Service
}

func createForm(db Service) *form {
	return &form{db}
}

func createFormRoute(app fastrex.App, f *form) fastrex.App {
	app.Get("/form/post", f.getPost).
		Post("/form/post", f.createPost).
		Post("/form/update", f.updatePost).
		Post("/form/signup", f.createUser).
		Post("/form/account", f.updateAccount).
		Post("/form/signin", f.getUserByEmailAndPassword)
	return app
}
