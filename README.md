# Google Cloud Function and Firebase

Get the latest commit hash:
```
go get github.com/fastrodev/rider@c6dfb2d
```

*`c6dfb2d` is the latest commit. You can see it in the rider github repository.*

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