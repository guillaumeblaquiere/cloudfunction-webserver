# Overview

This project is a hack example to servce file as a webserver with a Google Cloud Functions (normally single purpose,
only one entry point). It comes in addition of that
[Medium article](https://medium.com/google-cloud/hack-use-cloud-functions-as-a-webserver-with-golang-42edc7935247)

# Deployment

To deploy the Cloud Functions, you must have the gcloud SDK with a correct authentication that allows you to deploy
Cloud Functions

*You can change the region, and the security as you wish*

Deploy with runtime v1
```
gcloud beta functions deploy --runtime=v1 \
  --region=us-central1 --allow-unauthenticated \
  --runtime=go113 --trigger-http --entry-point=Webserver webserver
```

Deploy with runtime v2
```
gcloud beta functions deploy --runtime=v2 \
  --region=us-central1 --allow-unauthenticated \
  --runtime=go113 --trigger-http --entry-point=Webserver webserver
```

# URL paths to test

4 types of paths are provided by that Cloud Functions

*Update `<ProjectId>` with your own project ID*

```
# Dynamic path
https://us-central1-<ProjectId>.cloudfunctions.net/webserver/hello
https://us-central1-<ProjectId>.cloudfunctions.net/webserver/subroute/login

# Static resources
https://us-central1-<ProjectId>.cloudfunctions.net/webserver/static/
https://us-central1-<ProjectId>.cloudfunctions.net/webserver/static/index.html
https://us-central1-<ProjectId>.cloudfunctions.net/webserver/static/subdir/login.html
```

# License

This library is licensed under Apache 2.0. Full license text is available in
[LICENSE](https://github.com/guillaumeblaquiere/cloudrun-cloudfunction-compare/tree/master/LICENSE).