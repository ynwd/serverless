package internal

import (
	"time"

	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

type apiService struct {
	db database
}

func (s *apiService) createPost(req fastrex.Request, res fastrex.Response) {
	post := make(map[string]interface{})
	ads := make(map[string]interface{})

	topic := req.FormValue("topic")
	title := req.FormValue("title")
	content := req.FormValue("content")

	if topic == "" {
		ads["message"] = "topic tidak boleh kosong. lengkapi dg benar."
		res.Json(ads)
		return
	}

	if title == "" {
		ads["message"] = "judul tidak boleh kosong. lengkapi dg benar."
		res.Json(ads)
		return
	}

	if content == "" {
		ads["message"] = "isi iklan tidak boleh kosong. lengkapi dg benar."
		res.Json(ads)
		return
	}

	post["id"] = uuid.New().String()
	post["created"] = time.Now()
	post["topic"] = topic
	post["title"] = title
	post["content"] = content
	post["type"] = "ads"
	post["user"] = "user"

	s.db.addPost(req.Context(), post)
	ads["message"] = "iklan anda telah selesai disimpan. akan ditayangkan besok."
	res.Json(ads)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}
