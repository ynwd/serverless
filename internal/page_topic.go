package internal

import (
	"strings"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) topicPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("topic")
	user, _ := p.getUserFromSession(req, res)
	email := user.Email
	topic := strings.Title(params[0])
	data := struct {
		Email  string
		Title  string
		Domain string
	}{email, topic, DOMAIN}
	res.Render(data)
}
