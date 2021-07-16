package internal

import (
	"time"
	"unicode/utf8"

	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

type apiService struct {
	db database
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

	if topic == "" {
		msg = "topic tidak boleh kosong. lengkapi dg benar."
		createResponsePage(msg, res)
		return
	}

	if title == "" {
		msg = "judul tidak boleh kosong. lengkapi dg benar."
		createResponsePage(msg, res)
		return
	}

	if utf8.RuneCountInString(title) > 280 {
		msg = "judul iklan terlalu panjang. maksimal 280 karakter."
		createResponsePage(msg, res)
		return
	}

	if content == "" {
		msg = "isi iklan tidak boleh kosong. lengkapi dg benar."
		createResponsePage(msg, res)
		return
	}

	if utf8.RuneCountInString(content) > 280 {
		msg = "isi iklan terlalu panjang. maksimal 280 karakter."
		createResponsePage(msg, res)
		return
	}

	if user == "" {
		user = "user"
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

	s.db.addPost(req.Context(), post)
	msg = "iklan anda telah selesai disimpan. akan ditayangkan besok."
	createResponsePage(msg, res)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}

func (s *apiService) createUser(req fastrex.Request, res fastrex.Response) {
	user := make(map[string]interface{})
	var msg string

	email := req.FormValue("email")
	phone := req.FormValue("phone")
	password := req.FormValue("password")
	user["email"] = email
	user["phone"] = phone
	user["password"] = password

	s.db.addUser(req.Context(), user)
	msg = "data Anda telah tersimpan."
	createResponsePage(msg, res)
}

func (s *apiService) getUserByEmailAndPassword(req fastrex.Request, res fastrex.Response) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	post := req.FormValue("post")

	user, err := s.db.getUserDetail(req.Context(), email, password)
	if err != nil {
		createResponsePage("user tidak ditemukan. periksa email dan password anda", res)
	}
	c := fastrex.Cookie{}
	c.Name("__session").Value(user.Email).Path("/")

	if post != "" {
		url := "/post/" + post
		res.Cookie(c).Redirect(url, 302)
		return
	}

	res.Cookie(c).Redirect("/home", 302)
}
