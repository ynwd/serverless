package internal

import "github.com/fastrodev/fastrex"

type form struct {
	db Database
}

func createForm(db Database) *form {
	return &form{db}
}

func createFormRoute(app fastrex.App, f *form) fastrex.App {
	app.Get("/form/post", f.getPost).
		Post("/form/post", f.createPost).
		Post("/form/signup", f.createUser).
		Post("/form/signin", f.getUserByEmailAndPassword)
	return app
}
