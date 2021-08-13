package internal

import (
	"context"

	"github.com/fastrodev/fastrex"
)

func CreateApp() fastrex.App {
	ctx := context.Background()
	db := createDatabase(ctx)
	page := createPage(db)
	form := createForm(db)

	app := fastrex.New()
	app.Ctx(ctx).Static("public", "/public")
	app = createTemplate(app)
	app = createPageRoute(app, page)
	app = createHomePageRoute(app, page)
	app = createFormRoute(app, form)
	app = createApiRoute(app)
	return app
}
