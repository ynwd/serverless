package internal

import "github.com/fastrodev/fastrex"

func createTemplate(app fastrex.App) fastrex.App {
	return app.Template("template/index.gohtml",
		"template/arsip.gohtml",
		"template/signin.gohtml",
		"template/signup.gohtml",
		"template/membership.gohtml",
		"template/detail.gohtml",
		"template/response.gohtml",
		"template/result.gohtml",
		"template/header.gohtml",
		"template/headline.gohtml",
		"template/footer.gohtml",
		"template/meta.gohtml",
		"template/style.gohtml",
		"template/script.gohtml",
		"template/navigation.gohtml",
		"template/style_navigation.gohtml",
		"template/home.gohtml",
		"template/home_header.gohtml",
		"template/home_wrapper.gohtml",
		"template/home_post.gohtml",
		"template/home_post_new.gohtml",
		"template/home_topic.gohtml",
		"template/home_account.gohtml",
		"template/home_setting.gohtml")
}
