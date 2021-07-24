package internal

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

func (s *apiService) createUser(req fastrex.Request, res fastrex.Response) {
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
	username = strings.ToLower(username)
	user["email"] = email
	user["password"] = password
	user["id"] = uuid.New().String()
	user["username"] = username
	user["name"] = username

	_, _, err := s.db.createUser(req.Context(), user)
	if err != nil {
		createResponsePage(res, respTitle, err.Error(), "")
		return
	}

	msg = "Akun Anda telah tersimpan."
	createResponsePage(res, respTitle, msg, "/signin")
}

func readUserIP(req fastrex.Request) string {
	IPAddress := req.Header.Get("X-Real-Ip")
	fmt.Println("X-Real-Ip", IPAddress)
	if IPAddress == "" {
		IPAddress = req.Header.Get("X-Forwarded-For")
		fmt.Println("X-Forwarded-For", IPAddress)
	}
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
		fmt.Println("RemoteAddr", req.RemoteAddr)
	}
	return IPAddress
}

func (s *apiService) getUserByEmailAndPassword(req fastrex.Request, res fastrex.Response) {
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

	user, err := s.db.getUserDetail(req.Context(), email, password)
	if err != nil {
		createResponsePage(res, respTitle, "user tidak ditemukan. periksa email dan password anda", "")
		return
	}
	c := fastrex.Cookie{}
	userAgent := req.UserAgent()
	ip := readUserIP(req)
	ses := s.db.createSession(req.Context(), user.ID, userAgent, ip)
	sessionID := base64.StdEncoding.EncodeToString([]byte(ses))

	c.Name("__session").Value(sessionID).Path("/")

	// dapetin ID
	// sid, _ := base64.StdEncoding.DecodeString(sessionID)
	// userID, err := s.db.getUserIDWithSession(req.Context(), string(sid))
	// fmt.Println(userID)

	if post != "" {
		url := "/post/" + post
		res.Cookie(c).Redirect(url, 302)
		return
	}

	res.Cookie(c).Redirect("/home", 302)
}
