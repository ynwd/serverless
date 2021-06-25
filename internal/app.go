package internal

import "github.com/fastrodev/fastrex"

func CreateApp() fastrex.App {
	app := fastrex.New()
	app.Static("static")
	app.Template("template/index.html")
	app.Get("/iklan", htmlHandler)
	app.Get("/api", handler)
	return app
}

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

func htmlHandler(req fastrex.Request, res fastrex.Response) {
	data := map[string]interface{}{
		"topic":   "Motor",
		"title":   "Lorem ipsum dolor sit amet",
		"content": "consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	}

	res.Render(data)

	// type Iklan struct {
	// 	Title   string
	// 	Content string
	// }

	// type Topic struct {
	// 	Title     string
	// 	IklanList []Iklan
	// }

	// type IklanBaris struct {
	// 	Topics []Topic
	// }

	// iklanBaris := IklanBaris{
	// 	Topics: []Topic{
	// 		{
	// 			Title: "Motor",
	// 			IklanList: []Iklan{
	// 				{
	// 					Title:   "Lorem ipsum dolor sit amet",
	// 					Content: "consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 				},
	// 			},
	// 		},
	// 	},
	// }

}
