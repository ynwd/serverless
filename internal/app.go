package internal

import (
	"context"
	"io/ioutil"

	"github.com/fastrodev/fastrex"
)

func CreateApp() fastrex.App {
	ctx := context.Background()
	db := &database{client: createClient(ctx)}
	page := &pageService{db: *db}
	form := &formService{db: *db}

	app := fastrex.New()
	app.Ctx(ctx).Static("public", "/public")
	app = createPageRoute(app, page)
	app = createFormRoute(app, form)
	app = createTemplate(app)
	return app
}

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
		Template("template/footer.gohtml").
		Template("template/meta.gohtml").
		Template("template/style.gohtml")
	return app
}

func createPageRoute(app fastrex.App, page *pageService) fastrex.App {
	app.Post("/", healthChk).
		Get("/", page.idxPage).
		Get("/:username", page.userPage).
		Get("/post/:id", page.detailPage).
		Get("/topic/:topic", page.topicPage).
		Get("/search", page.queryPage).
		Post("/search", page.searchPage).
		Get("/activate/:code", page.activatePage)
	return app
}

func createFormRoute(app fastrex.App, form *formService) fastrex.App {
	app.Get("/form/post", form.getPost).
		Post("/form/post", form.createPost).
		Post("/form/signup", form.createUser).
		Post("/form/signin", form.getUserByEmailAndPassword)
	return app
}

func healthChk(req fastrex.Request, res fastrex.Response) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Send(err.Error())
		return
	}
	msg := string(body)
	res.Send(msg)
}
