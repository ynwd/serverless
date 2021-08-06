package internal

import "github.com/fastrodev/fastrex"

type formService struct {
	db Database
}

func createForm(db Database) *formService {
	return &formService{db}
}

func createFormRoute(app fastrex.App, form *formService) fastrex.App {
	app.Get("/form/post", form.getPost).
		Post("/form/post", form.createPost).
		Post("/form/signup", form.createUser).
		Post("/form/signin", form.getUserByEmailAndPassword)
	return app
}
