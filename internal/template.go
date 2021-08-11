package internal

import "github.com/fastrodev/fastrex"

func createTemplate(app fastrex.App) fastrex.App {
	app.Template("template/index.gohtml").
		Template("template/arsip.gohtml").
		Template("template/signin.gohtml").
		Template("template/signup.gohtml").
		Template("template/membership.gohtml").
		Template("template/home.gohtml").
		Template("template/detail.gohtml").
		Template("template/create.gohtml").
		Template("template/response.gohtml").
		Template("template/result.gohtml").
		Template("template/header.gohtml").
		Template("template/headline.gohtml").
		Template("template/footer.gohtml").
		Template("template/meta.gohtml").
		Template("template/style.gohtml").
		Template("template/script.gohtml").
		Template("template/navigation.gohtml")
	return app
}
