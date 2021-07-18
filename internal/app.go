package internal

import (
	"context"

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
		Template("template/arsip.html").
		Template("template/signin.html").
		Template("template/signup.html").
		Template("template/membership.html").
		Template("template/home.html").
		Template("template/detail.html").
		Template("template/create.html").
		Template("template/response.html")

	return app
}

func createPageRoute(ctx context.Context, app fastrex.App) fastrex.App {
	s := createPageService(ctx)
	app.Get("/", s.rootPage).
		Get("/u/:id", s.userPage).
		Get("/home", s.homePage).
		Get("/arsip", s.arsipPage).
		Get("/signin", s.signinPage).
		Get("/signup", s.signupPage).
		Get("/signout", s.signOut).
		Get("/membership", s.membershipPage).
		Get("/post", s.createPostPage).
		Get("/post/:id", s.detailPage).
		Get("/post/topic/:topic", s.topicPage)
	return app
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
