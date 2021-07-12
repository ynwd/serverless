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
	app.Template("template/search.html").
		Template("template/signin.html").
		Template("template/signup.html").
		Template("template/membership.html").
		Template("template/home.html").
		Template("template/detail.html")
	return app
}

func createPageRoute(ctx context.Context, app fastrex.App) fastrex.App {
	s := createPageService(ctx)
	app.Get("/search", s.searchPage).
		Get("/signin", s.signinPage).
		Get("/signup", s.signupPage).
		Get("/membership", s.membershipPage).
		Get("/home", s.homePage).
		Get("/post", s.createPostPage).
		Get("/detail/:post", s.detailPage)
	return app
}

func createApiRoute(ctx context.Context, app fastrex.App) fastrex.App {
	api := createApiService(ctx)
	app.Get("/api", api.getPost).
		Post("/api", api.createPost)
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
