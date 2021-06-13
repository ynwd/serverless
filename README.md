# Google Cloud Function and Firebase

Get the latest commit hash:
```
go get github.com/fastrodev/fastrex@68d585f
```

> Note: *`68d585f` is the latest commit. You can see it in the [rider github repository](https://github.com/fastrodev/fastrex).*

Deploy to google cloud function:
```
gcloud functions deploy Entrypoint --runtime go113 --trigger-http --allow-unauthenticated
```

Live demo: [https://us-central1-phonic-altar-274306.cloudfunctions.net/Entrypoint](https://us-central1-phonic-altar-274306.cloudfunctions.net/Entrypoint)

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
        "function": "Entrypoint"
      }
    ],
    "ignore": [
      "static/**",
      ".gitignore",
      "firebase.json",
      "go.mod",
      "go.sum",
      "hello.go",
      "index.html",
      "README.md",
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