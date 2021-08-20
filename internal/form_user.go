package internal

import (
	"encoding/base64"
	"log"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

func (s *form) createUser(req fastrex.Request, res fastrex.Response) {
	user := make(map[string]interface{})
	var msg string
	respTitle := "Daftar"

	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")

	if username == "" {
		createResponsePage(res, respTitle, "Username tidak boleh kosong", "")
		return
	}
	re := regexp.MustCompile("^[a-zA-Z0-9-_]+$")
	if !re.MatchString(username) {
		createResponsePage(res, respTitle, "Username harus berupa angka dan huruf. Tidak boleh ada spasi.", "")
		return
	}

	if email == "" {
		createResponsePage(res, respTitle, "Email tidak boleh kosong", "")
		return
	}
	_, errEmail := mail.ParseAddress(email)
	if errEmail != nil {
		createResponsePage(res, respTitle, "Email yang kamu masukkan tidak valid", "")
		return
	}

	if password == "" {
		createResponsePage(res, respTitle, "Password tidak boleh kosong", "")
		return
	}
	code := String(6)
	username = strings.ToLower(username)
	user["email"] = email
	user["password"] = password
	user["id"] = uuid.New().String()
	user["username"] = username
	user["name"] = username
	user["active"] = false
	user["code"] = code
	user["created"] = time.Now()

	_, _, err := s.svc.createUser(req.Context(), user)
	if err != nil {
		createResponsePage(res, respTitle, err.Error(), "")
		return
	}

	go func() {
		err := SendEmail(email, code)
		if err != nil {
			log.Println(err.Error())
		}
	}()

	msg = "Pendaftaran berhasil. Cek email dan buka link aktivasi yang telah terkirim."
	createResponsePage(res, respTitle, msg, "/signin")
}

func readUserIP(req fastrex.Request) string {
	IPAddress := req.Header.Get("X-Real-Ip")
	// fmt.Println("X-Real-Ip", IPAddress)
	if IPAddress == "" {
		IPAddress = req.Header.Get("X-Forwarded-For")
		// fmt.Println("X-Forwarded-For", IPAddress)
	}
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
		// fmt.Println("RemoteAddr", req.RemoteAddr)
	}
	return IPAddress
}

func (s *form) getUserByEmailAndPassword(req fastrex.Request, res fastrex.Response) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	post := req.FormValue("post")
	respTitle := "Masuk"

	if email == "" {
		createResponsePage(res, respTitle, "Email tidak boleh kosong", "")
		return
	}

	if password == "" {
		createResponsePage(res, respTitle, "Password tidak boleh kosong", "")
		return
	}

	user, err := s.svc.getUserDetail(req.Context(), email, password)
	if err != nil {
		createResponsePage(res, respTitle, "user tidak ditemukan. periksa email dan password anda", "")
		return
	}
	c := fastrex.Cookie{}
	userAgent := req.UserAgent()
	ip := readUserIP(req)
	ses := s.svc.createSession(req.Context(), user.ID, userAgent, ip)
	sessionID := base64.StdEncoding.EncodeToString([]byte(ses))

	c.Name("__session").Value(sessionID).Path("/")

	if post != "" {
		url := "/post/" + post
		res.Cookie(c).Redirect(url, 302)
		return
	}

	res.Cookie(c).Redirect("/home", 302)
}

func (s *form) updateAccount(req fastrex.Request, res fastrex.Response) {
	var msg string
	respTitle := "Update Account"

	userID := req.FormValue("userID")
	name := req.FormValue("name")
	phone := req.FormValue("phone")
	password := req.FormValue("password")
	if password == "" {
		createResponsePage(res, respTitle, "Password tidak boleh kosong", "")
		return
	}

	u := s.getUserByID(req, userID)
	if u.Password != password {
		createResponsePage(res, respTitle, "Password yang kamu masukkan salah", "")
		return
	}

	update := []firestore.Update{}
	if name != "" {
		update = append(update, firestore.Update{Path: "name", Value: name})
	}
	if phone != "" {
		update = append(update, firestore.Update{Path: "phone", Value: phone})
	}
	update = append(update, firestore.Update{Path: "updated", Value: time.Now()})

	_, err := s.svc.update(req.Context(), &Query{
		Collection: "user",
		Field:      "id",
		Op:         "==",
		Value:      userID,
		OrderBy:    "created",
	}, update)
	if err != nil {
		log.Println(err.Error())
	}

	msg = "Update account berhasil."
	createResponsePage(res, respTitle, msg, "")
}

func (s *form) getUserByID(req fastrex.Request, userID string) *User {
	data, err := s.svc.get(req.Context(), &Query{
		Collection: "user",
		Field:      "id",
		Value:      userID,
		Op:         "==",
		OrderBy:    "created",
	})

	if err != nil {
		log.Println("err:getUserByID", err.Error())
	}

	x := data.Data()
	p := x["password"].(string)

	return &User{
		Password: p,
	}
}
