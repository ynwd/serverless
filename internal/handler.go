package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fastrodev/fastrex"
)

func readJson(host string) Data {
	schema := "http://"
	if os.Getenv("GCP") == "true" {
		schema = "https://"
	}

	fullURLFile := schema + host + "/iklan.json"

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, errGet := client.Get(fullURLFile)
	if errGet != nil {
		log.Fatal(errGet)
	}

	body, errReadall := ioutil.ReadAll(resp.Body)
	if errReadall != nil {
		log.Fatal(errGet)
	}

	data := Data{}
	errUnmarshal := json.Unmarshal(body, &data)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	defer resp.Body.Close()
	return data
}

func htmlHandler(req fastrex.Request, res fastrex.Response) {
	json := readJson(req.Host)
	res.Render(json.Data)
}
