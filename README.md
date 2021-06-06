# Cloud Function and Firebase

Get the latest commit hash
```
go get github.com/fastrodev/rider@c6dfb2d
```

Deploy to google cloud function
```
gcloud functions deploy HelloHTTP --runtime go113 --trigger-http --allow-unauthenticated
```

Live demo: [https://us-central1-phonic-altar-274306.cloudfunctions.net/HelloHTTP](https://us-central1-phonic-altar-274306.cloudfunctions.net/HelloHTTP)

## Setup domain

Install firebase:
```
npm install -g firebase-tools
```

Init hosting:
```
firebase init hosting
```

Change firebase config:
```json
{
  "hosting": {
    "public": "public",
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

Change public/index.html to public/app.html:
```
mv public/index.html public/app.html
```

Deploy:
```
firebase deploy
```

Live demo: [https://phonic-altar-274306.web.app/](https://phonic-altar-274306.web.app/)