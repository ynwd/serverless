package internal

import (
	"context"

	"github.com/fastrodev/fastrex"
)

func CreateApp() fastrex.App {
	ctx := context.Background()
	db := createDatabase(ctx)
	page := &pageService{db}
	form := &formService{db}

	app := fastrex.New()
	app.Ctx(ctx).Static("public", "/public")
	app = createTemplate(app)
	app = createPageRoute(app, page)
	app = createFormRoute(app, form)
	app = createApiRoute(app)
	return app
}
