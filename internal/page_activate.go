package internal

import (
	"log"

	"github.com/fastrodev/fastrex"
)

func (p *page) activatePage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("code")
	code := ""
	if len(params) > 0 {
		code = params[0]
	}
	it, err := p.db.getUserByActivationCode(req.Context(), code)
	if err != nil {
		log.Println("activatePage:", err)
	}

	if it != nil {
		go p.db.activateUserByCode(req.Context(), code)
		createResponsePage(res, "Aktivasi", "User telah diaktifkan. silahkan masuk.", "/signin")
		return
	}

	createResponsePage(res, "Aktivasi", "Aktivasi gagal. Pastikan kode yang anda masukkan benar.", "")
}
