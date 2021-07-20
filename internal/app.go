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
	app = createApiRoute(ctx, app)
	app.Ctx(ctx).Static("public")
	app = createTemplate(app)
	return app
}

func createTemplate(app fastrex.App) fastrex.App {
	app.Template("public/index.html").
		Template("template/arsip.gohtml").
		Template("template/signin.gohtml").
		Template("template/signup.gohtml").
		Template("template/membership.gohtml").
		Template("template/home.gohtml").
		Template("template/detail.gohtml").
		Template("template/create.gohtml").
		Template("template/response.gohtml")

	return app
}

func createPageRoute(ctx context.Context, app fastrex.App) fastrex.App {
	s := createPageService(ctx)
	app.Get("/", s.idxPage).
		Post("/", receiveEvent).
		Get("/:id", s.userPage).
		Get("/post/:id", s.detailPage).
		Get("/post/topic/:topic", s.topicPage)
	return app
}

func receiveEvent(req fastrex.Request, res fastrex.Response) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error on receive event from pubsub %v", err.Error())
	}
	msg := string(body)
	log.Printf("receiveEvent:%v", msg)
	res.Send(msg)
}

func createApiRoute(ctx context.Context, app fastrex.App) fastrex.App {
	api := createApiService(ctx)
	app.Get("/api", api.getPost).
		Post("/api", api.createPost).
		Post("/api/signup", api.createUser).
		Post("/api/signin", api.getUserByEmailAndPassword)
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
