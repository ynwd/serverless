package internal

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/fastrodev/fastrex"
)

func CreateApp() fastrex.App {
	app := fastrex.New()
	ctx := context.Background()
	app = createPageRoute(ctx, app)
	app = createFormRoute(ctx, app)
	app = createTemplate(app)
	app.Ctx(ctx).Static("public", "/public")
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

func createPageRoute(ctx context.Context, app fastrex.App) fastrex.App {
	s := createPageService(ctx)
	app.Post("/", healthChk).
		Get("/", s.idxPage).
		Get("/:username", s.userPage).
		Get("/post/:id", s.detailPage).
		Get("/topic/:topic", s.topicPage).
		Get("/search", s.queryPage).
		Post("/search", s.searchPage).
		Get("/activate/:code", s.activatePage)
	return app
}

func healthChk(req fastrex.Request, res fastrex.Response) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error on receive event from pubsub %v", err.Error())
	}
	msg := string(body)
	log.Printf("receiveEvent:%v", msg)
	res.Send(msg)
}

func createFormRoute(ctx context.Context, app fastrex.App) fastrex.App {
	api := createApiService(ctx)
	app.Get("/form/post", api.getPost).
		Post("/form/post", api.createPost).
		Post("/form/signup", api.createUser).
		Post("/form/signin", api.getUserByEmailAndPassword)
	return app
}

func createDatabase(ctx context.Context) *database {
	return &database{client: createClient(ctx)}
}

func createApiService(ctx context.Context) *apiService {
	db := createDatabase(ctx)
	return &apiService{db: *db}
}

func createPageService(ctx context.Context) *pageService {
	db := createDatabase(ctx)
	return &pageService{db: *db}
}
