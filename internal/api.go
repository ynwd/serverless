package internal

import (
	"context"
	"fmt"
	"io"
	"time"
	"unicode/utf8"

	"cloud.google.com/go/storage"
	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

type apiService struct {
	db database
}

func saveToGCS(ctx context.Context, r io.Reader, bucketName, name string) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	// Next check if the bucket exists
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

	topic := req.FormValue("topic")
	title := req.FormValue("title")
	content := req.FormValue("content")
	address := req.FormValue("address")
	email := req.FormValue("email")
	phone := req.FormValue("phone")
	user := req.FormValue("user")
	video := req.FormValue("video")

	file := ""
	id := uuid.New().String()
	url := "/"
	req.ParseMultipartForm(32 << 20)
	uploadedFile, _, err := req.FormFile("file")
	if uploadedFile != nil {
		if err != nil {
			createResponsePage(err.Error(), "/", res)
			return
		}
		defer uploadedFile.Close()

		_, _, err := saveToGCS(req.Context(), uploadedFile, "fastro-images", id)
		if err != nil {
			createResponsePage(err.Error(), "/", res)
			fmt.Printf("GCS is not setup %v\n", err)
			return
		}

		file = "https://storage.googleapis.com/fastro-images/" + id
	}

	if topic == "" {
		msg = "Topic tidak boleh kosong. pilih salah satu."
		createResponsePage(msg, url, res)
		return
	}

	if title == "" {
		msg = "Judul tidak boleh kosong. lengkapi dg benar."
		createResponsePage(msg, url, res)
		return
	}

	if utf8.RuneCountInString(title) > 100 {
		msg = "Judul iklan terlalu panjang. maksimal 100 karakter."
		createResponsePage(msg, url, res)
		return
	}

	if content == "" {
		msg = "Isi iklan tidak boleh kosong. lengkapi dg benar."
		createResponsePage(msg, url, res)
		return
	}

	if utf8.RuneCountInString(content) > 280 {
		msg = "Isi iklan terlalu panjang. maksimal 280 karakter."
		createResponsePage(msg, url, res)
		return
	}

	if user == "" {
		user = "user"
	}

	if address == "" && user != "user" {
		msg = "Kota tidak boleh kosong."
		createResponsePage(msg, url, res)
		return
	}
	if phone == "" && user != "user" {
		msg = "Phone tidak boleh kosong."
		createResponsePage(msg, url, res)
		return
	}
	if email == "" && user != "user" {
		msg = "Email tidak boleh kosong."
		createResponsePage(msg, url, res)
		return
	}

	post["id"] = uuid.New().String()
	post["created"] = time.Now()
	post["topic"] = topic
	post["title"] = title
	post["content"] = content
	post["email"] = email
	post["phone"] = phone
	post["address"] = address
	post["type"] = "ads"
	post["user"] = user
	post["file"] = file
	post["video"] = video

	s.db.addPost(req.Context(), post)

	msg = "Iklan anda telah selesai disimpan. akan ditayangkan besok."
	createResponsePage(msg, url, res)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}

func (s *apiService) createUser(req fastrex.Request, res fastrex.Response) {
	user := make(map[string]interface{})
	var msg string

	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")

	user["name"] = name
	user["email"] = email
	user["password"] = password

	s.db.addUser(req.Context(), user)
	msg = "data Anda telah tersimpan."
	url := "/"
	createResponsePage(msg, url, res)
}

func (s *apiService) getUserByEmailAndPassword(req fastrex.Request, res fastrex.Response) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	post := req.FormValue("post")

	user, err := s.db.getUserDetail(req.Context(), email, password)
	if err != nil {
		url := "/signin"
		createResponsePage("user tidak ditemukan. periksa email dan password anda", url, res)
		return
	}
	c := fastrex.Cookie{}
	c.Name("__session").Value(user.ID).Path("/")

	if post != "" {
		url := "/post/" + post
		res.Cookie(c).Redirect(url, 302)
		return
	}

	res.Cookie(c).Redirect("/home", 302)
}
