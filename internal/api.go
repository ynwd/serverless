package internal

import (
	"io/ioutil"

	"github.com/fastrodev/fastrex"
)

func createApiRoute(app fastrex.App) fastrex.App {
	app.Post("/", healthChk)
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
