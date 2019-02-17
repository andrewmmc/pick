# pick

## Deploy
```
gcloud beta functions deploy install --env-vars-file .env.yaml --runtime go111 --entry-point Install --trigger-http --region asia-northeast1
gcloud beta functions deploy authCallback --env-vars-file .env.yaml --runtime go111 --entry-point AuthCallback --trigger-http --region asia-northeast1
gcloud beta functions deploy getAnswer --env-vars-file .env.yaml --runtime go111 --entry-point GetAnswer --trigger-http --region asia-northeast1
```