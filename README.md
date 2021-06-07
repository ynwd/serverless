# Google Cloud Function and Firebase

Get the latest commit hash:
```
go get github.com/fastrodev/rider@3b9e2eb
```

> Note: *`3b9e2eb` is the latest commit. You can see it in the [rider github repository](https://github.com/fastrodev/rider).*

Entry point:
```go
package serverless

import (
  "net/http"

  "github.com/fastrodev/rider"
)

func htmlHandler(req rider.Request, res rider.Response) {
  data := map[string]interface{}{
    "title": "Cloud function app",
    "name":  "Agus",
  }
  res.Render(data)
}

func apiHandler(req rider.Request, res rider.Response) {
  res.Json(`{"message":"ok"}`)
}

func createApp() rider.Router {
  app := rider.New()
  app.Static("static")
  app.Template("index.html")
  app.Get("/html", htmlHandler)
  app.Get("/api", apiHandler)
  return app
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
  createApp().ServeHTTP(w, r)
}

```

Deploy to google cloud function:
```
gcloud functions deploy HelloHTTP --runtime go113 --trigger-http --allow-unauthenticated
```

Live demo: [https://us-central1-phonic-altar-274306.cloudfunctions.net/HelloHTTP](https://us-central1-phonic-altar-274306.cloudfunctions.net/HelloHTTP)

## Firebase

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
        "function": "HelloHTTP"
      }
    ],
    "ignore": [
      "firebase.json",
      "**/.*",
      "**/node_modules/**"
    ]
  }
}
```

Deploy to firebase:
```
firebase deploy
```

Live demo: [https://phonic-altar-274306.web.app](https://phonic-altar-274306.web.app)