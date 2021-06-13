package serverless

import (
	"net/http"

	"github.com/fastrodev/serverless/internal"
)

func Main(w http.ResponseWriter, r *http.Request) {
	internal.CreateApp().ServeHTTP(w, r)
}
