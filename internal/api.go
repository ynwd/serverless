package internal

import (
	"encoding/json"
	"time"

	"github.com/fastrodev/fastrex"
	"github.com/google/uuid"
)

type apiService struct {
	db database
}

func (s *apiService) createPost(req fastrex.Request, res fastrex.Response) {
	var post map[string]interface{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&post)
	if err != nil {
		panic(err)
	}
	post["id"] = uuid.New().String()
	post["created"] = time.Now()
	s.db.addPost(req.Context(), post)
	res.Json(post)
}

func (s *apiService) getPost(req fastrex.Request, res fastrex.Response) {
	d := s.db.getPost(req.Context())
	res.Json(d)
}
