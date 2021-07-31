package internal

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/fastrodev/fastrex"
)

func (p *pageService) activatePage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("code")
	code := ""
	if len(params) > 0 {
		code = params[0]
	}

	iter := p.db.client.CollectionGroup("user").
		Where("code", "==", code).
		Where("active", "==", false).
		OrderBy("created", firestore.Desc).
		Documents(context.Background())

	it, err := getDocumentSnapshot(iter)

	if err != nil {
		log.Println("activatePage:", err)
	}

	if it != nil {
		go p.activateUserByCode(req.Context(), code)
		createResponsePage(res, "Aktivasi", "User telah diaktifkan. silahkan masuk.", "/signin")
		return
	}

	createResponsePage(res, "Aktivasi", "Aktivasi gagal. Pastikan kode yang anda masukkan benar.", "")
}

func (p *pageService) activateUserByCode(ctx context.Context, code string) {
	iter := p.db.client.CollectionGroup("user").
		Where("code", "==", code).
		Where("active", "==", false).
		Documents(ctx)
	it, err := getDocumentSnapshot(iter)
	if err != nil {
		log.Println(err)
	}

	p.db.client.Collection("user").Doc(it.Ref.ID).Update(ctx, []firestore.Update{
		{
			Path:  "active",
			Value: true,
		},
	})
}
