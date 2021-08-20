package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) homeAccountPage(req fastrex.Request, res fastrex.Response) {
	user, err := p.getUserFromSession(req, res)
	if err != nil {
		res.Redirect("/", 302)
		return
	}

	// fmt.Println("user", user)

	data := struct {
		Title           string
		Initial         string
		AccountUsername string
		AccountName     string
		AccountEmail    string
		AccountPhone    string
		AccountPassword string
		AccountID       string
	}{
		"Account", "T", user.Username, user.Name, user.Email, user.Phone, user.Password, user.ID,
	}

	err = res.Render("home_account", data)
	if err != nil {
		log.Println(err.Error())
	}
}
