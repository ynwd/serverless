# Fastrex Serverless

Google provides a very efficient serverless service. A function will serve if someone uses it and will iddle if no one is using it. You pay nothing when your function is idle. *So cheap, right?*

With [fastrex](https://github.com/fastrodev/fastrex) you can create golang based fullstack applications and install them in cloud functions. You can create multiple routes, handle static files, and render html templates -- then associate them with an entry point in the cloud function. The entry point is then redirected to firebase hosting according to the domain and url you want.

Live demo: [https://phonic-altar-274306.web.app](https://phonic-altar-274306.web.app)

## Getting start
Create serverless webapp folder: `serverless`
```
mkdir serverless && cd serverless
go mod init github.com/fastrodev/serverless
go get github.com/fastrodev/fastrex
```
You can rename the module `github.com/fastrodev/serverless` to anything.

## Create web application
create an app folder: `internal`
```
mkdir internal
```

create a webapp file: `internal/app.go`
```go
package internal

import "github.com/fastrodev/fastrex"

func CreateApp() fastrex.App {
	app := fastrex.New()
	app.Get("/", handler)
	return app
}

func handler(req fastrex.Request, res fastrex.Response) {
	res.Json(`{"message":"ok"}`)
}

```

> *You can add new routes, handlers, templates and static files later. See the full source code at: [internal/app.go](internal/app.go)*
## Localhost entrypoint
To test locally, you can create a localhost webapp entrypoint file: `cmd/main.go`
```go
package main

import (
	"github.com/fastrodev/serverless/internal"
)

func main() {
	internal.CreateApp().Listen(9000)
}

```

> *You can see the full source code at: [cmd/main.go](cmd/main.go)*

You can run by this command:
```
go run cmd/main.go
```
You can access it via url:
```
http://localhost:9000
```

## Serverless entrypoint
To see serverless endpoint, you must create a serverless webapp entrypoint file: `serverless.go`
```go
package serverless

import (
	"net/http"

	"github.com/fastrodev/serverless/internal"
)

func Main(w http.ResponseWriter, r *http.Request) {
  internal.CreateApp().Serverless(true).ServeHTTP(w, r)
}
```

> *You can see the full source code at: [serverless.go](serverless.go)*

## Cloud function deployment
> *Prerequisite: [Cloud SDK](https://cloud.google.com/sdk/docs/quickstart)*

To see live serverless endpoint, you must deploy to google cloud function:
```
gcloud functions deploy Main --runtime go113 --trigger-http --allow-unauthenticated
```

Live demo: [https://us-central1-phonic-altar-274306.cloudfunctions.net/Main](https://us-central1-phonic-altar-274306.cloudfunctions.net/Entrypoint)

## Firebase domain setup

We will redirect above url to firebase domain. You can change it with your own from firebase dashboard.

Install firebase:
```
npm install -g firebase-tools
```

Init hosting files:
```
firebase init hosting
```

Remove public folder:
```
rm -rf public
```

Change firebase config:
```json
{
  "hosting": {
    "rewrites": [
      {
        "source": "**",
        "function": "Main"
      }
    ],
    "ignore": [
      "**/cmd/**",
      "**/internal/**",
      "**/static/**",
      "**/template/**",
      ".gitignore",
      "cloudbuild.yaml",
      "firebase.json",
      "go.mod",
      "go.sum",
      "README.md",
      "serverless.go",
      "**/.*"
    ]
  }
}
```

Deploy to firebase:
```
firebase deploy
```

Live demo: [https://phonic-altar-274306.web.app](https://phonic-altar-274306.web.app)