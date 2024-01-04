## Alertmanager Dingtalk Webhook
Webhook for sending message.

Message Type supported: {msgType}
- Raw string: raw
- Prometheus Alert Notification: prom
- Huawei SMN: huaweismn
- AlibabaCloud CloudMonitor: raw 

Target medium supported: {msgMedium}
- DingTalk: ding

Example: 

```
go run main.go -secret
curl -XPOST host:9000/webhook/{msgType}/{msgMedium}?secret=asdfasdfasdfasdfasdf -d "helloworld"

go run main.go
curl -XPOST host:9000/webhook/{msgType}/{msgMedium} -d "helloworld"
```

### Config
```
export DING_ROBOT_TOKEN=xxxxxxxxxxxxxxxxx
export DING_ROBOT_SECRET=SECxxxxxxxxxxxxxxxxxx
export WEBHOOK_SECRET_KEY=asdasdfasdfasdfasdf

```

### Deploy to K8s cluster
```
kubectl apply -f deploy.yaml

# If you want to use secret to store Environment Variables, 
# adapt the "deploy.yaml", and create secret by below command:
kubectl create secret generic alert-webhook-receiver-secret \
    --from-literal=token=<your_token> \
    --from-literal=secret=<your_secret> \
    --from-literal=webhook-secret-key=<your_webhook_secret_key> \
    -n monitoring

```

Set Env value from secret:

```
spec:
  template:
    spec:
      containers:
        - env:
            - name: DING_ROBOT_TOKEN
              valueFrom:
                secretKeyRef:
                  name: alert-webhook-receiver-secret
                  key: token
            - name: DING_ROBOT_SECRET
              valueFrom:
                secretKeyRef:
                  name: alert-webhook-receiver-secret
                  key: secret
            - name: WEBHOOK_SECRET_KEY
              valueFrom:
                secretKeyRef:
                  name: alert-webhook-receiver-secret
                  key: webhook-secret-key
```

### Build image

```
docker build -t <NAME>/alert-webhook-receiver:<TAG> .

docker run -d -p 8000:8000 \
    -e DING_ROBOT_TOKEN=xxxxxxxxxxxxxxxxx \
    -e DING_ROBOT_SECRET=SECxxxxxxxxxxxxxxxxxx \
    -e WEBHOOK_SECRET_KEY=asdasdfasdfasdfasdf \
    --name webhook <NAME>/alert-webhook-receiver:<TAG> -port 8000 -secret

docker image prune -f
```

