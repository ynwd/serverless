package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/mail"
	"regexp"
	"strconv"
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
			createResponsePage(respTitle, err.Error(), "", res)
			fmt.Printf("GCS is not setup %v\n", err)
			return
		}

		file = GCS_URL + id
	}

	if topic == "" {
		msg = "Topic tidak boleh kosong. Pilih salah satu."
		createResponsePage(respTitle, msg, "", res)
		return
	}

	if title == "" {
		msg = "Judul tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
		return
	}

	if utf8.RuneCountInString(title) > 100 {
		msg = "Judul iklan maksimal 100 karakter."
		createResponsePage(respTitle, msg, "", res)
		return
	}

	if content == "" {
		msg = "Isi iklan tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
		return
	}

	if priceStr == "" {
		msg = "Harga tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
		return
	}
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		createResponsePage(respTitle, err.Error(), "", res)
		return
	}

	if utf8.RuneCountInString(content) > 280 {
		msg = "Isi iklan maksimal 280 karakter."
		createResponsePage(respTitle, msg, "", res)
		return
	}

	if user == "" {
		user = "user"
	}

	if address == "" && user != "user" {
		msg = "Kota tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
		return
	}
	if phone == "" && user != "user" {
		msg = "Phone tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
		return
	}
	if email == "" && user != "user" {
		msg = "Email tidak boleh kosong."
		createResponsePage(respTitle, msg, "", res)
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
	createResponsePage(respTitle, msg, url, res)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}

func (s *apiService) createUser(req fastrex.Request, res fastrex.Response) {
	user := make(map[string]interface{})
	var msg string
	respTitle := "Daftar"

	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")

	if username == "" {
		createResponsePage(respTitle, "Username tidak boleh kosong", "", res)
		return
	}
	re := regexp.MustCompile("^[a-zA-Z0-9-_]+$")
	if !re.MatchString(username) {
		createResponsePage(respTitle, "Username harus berupa angka dan huruf. Tidak boleh ada spasi.", "", res)
		return
	}

	if email == "" {
		createResponsePage(respTitle, "Email tidak boleh kosong", "", res)
		return
	}
	_, errEmail := mail.ParseAddress(email)
	if errEmail != nil {
		createResponsePage(respTitle, "Email yang kamu masukkan tidak valid", "", res)
		return
	}

	if password == "" {
		createResponsePage(respTitle, "Password tidak boleh kosong", "", res)
		return
	}

	user["email"] = email
	user["password"] = password
	user["id"] = uuid.New().String()
	user["username"] = username
	user["name"] = username

	_, _, err := s.db.addUser(req.Context(), user)
	if err != nil {
		createResponsePage(respTitle, err.Error(), "", res)
		return
	}

	msg = "Akun Anda telah tersimpan."
	createResponsePage(respTitle, msg, "/signin", res)
}

func (s *apiService) getUserByEmailAndPassword(req fastrex.Request, res fastrex.Response) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	post := req.FormValue("post")
	respTitle := "Masuk"

	if email == "" {
		createResponsePage(respTitle, "Email tidak boleh kosong", "", res)
		return
	}

	if password == "" {
		createResponsePage(respTitle, "Password tidak boleh kosong", "", res)
		return
	}

	user, err := s.db.getUserDetail(req.Context(), email, password)
	if err != nil {
		createResponsePage(respTitle, "user tidak ditemukan. periksa email dan password anda", "", res)
		return
	}
	c := fastrex.Cookie{}
	userID := base64.StdEncoding.EncodeToString([]byte(user.ID))
	c.Name("__session").Value(userID).Path("/")

	if post != "" {
		url := "/post/" + post
		res.Cookie(c).Redirect(url, 302)
		return
	}

	res.Cookie(c).Redirect("/home", 302)
}
