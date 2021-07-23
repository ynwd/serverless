package internal

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"
	"unicode/utf8"

	"cloud.google.com/go/storage"
	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

func saveToGCS(ctx context.Context, r io.Reader, bucketName, name string) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	if _, err = bucket.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bucket.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, nil, err
	}

	attrs, err := obj.Attrs(ctx)
	fmt.Printf("Post is saved to GCS: %s\n", attrs.MediaLink)
	return obj, attrs, err
}

func (s *apiService) createPost(req fastrex.Request, res fastrex.Response) {
	post := make(map[string]interface{})
	var msg string
	respTitle := "Pasang Iklan"

	topic := req.FormValue("topic")
	title := req.FormValue("title")
	content := req.FormValue("content")
	priceStr := req.FormValue("price")
	address := req.FormValue("address")
	email := req.FormValue("email")
	phone := req.FormValue("phone")
	user := req.FormValue("user")
	video := req.FormValue("video")

	file := ""
	id := uuid.New().String()
	req.ParseMultipartForm(32 << 20)
	uploadedFile, _, _ := req.FormFile("file")
	if uploadedFile != nil {
		defer uploadedFile.Close()

		_, _, err := saveToGCS(req.Context(), uploadedFile, BUCKET_NAME, id)
		if err != nil {
			createResponsePage(res, respTitle, err.Error(), "")
			fmt.Printf("GCS is not setup %v\n", err)
			return
		}

		file = GCS_URL + id
	}

	if topic == "" {
		msg = "Topic tidak boleh kosong. Pilih salah satu."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	if title == "" {
		msg = "Judul tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	if utf8.RuneCountInString(title) > 100 {
		msg = "Judul iklan maksimal 100 karakter."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	if content == "" {
		msg = "Isi iklan tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	if priceStr == "" {
		msg = "Harga tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		createResponsePage(res, respTitle, err.Error(), "")
		return
	}

	if utf8.RuneCountInString(content) > 280 {
		msg = "Isi iklan maksimal 280 karakter."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	if user == "" {
		user = "user"
	}

	if address == "" && user != "user" {
		msg = "Alamat tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}
	if phone == "" && user != "user" {
		msg = "WhatsApp tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}
	if email == "" && user != "user" {
		msg = "Email tidak boleh kosong."
		createResponsePage(res, respTitle, msg, "")
		return
	}

	postID := uuid.New().String()
	post["id"] = postID
	post["created"] = time.Now()
	post["topic"] = topic
	post["title"] = title
	post["content"] = content
	post["price"] = price
	post["email"] = email
	post["phone"] = phone
	post["address"] = address
	post["type"] = "ads"
	post["user"] = user
	post["file"] = file
	post["video"] = video

	s.db.addPost(req.Context(), post)

	msg = "Iklan telah selesai disimpan. Klik tombol berikut untuk melihatnya."
	url := "/post/" + postID
	createResponsePage(res, respTitle, msg, url)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}
