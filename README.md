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

## Config
```
export DING_ROBOT_TOKEN=xxxxxxxxxxxxxxxxx
export DING_ROBOT_SECRET=SECxxxxxxxxxxxxxxxxxx
export WEBHOOK_SECRET_KEY=asdasdfasdfasdfasdf

```

## Secret
```
kubectl create secret generic alert-webhook-receiver-secret \
--from-literal=token=<your_token> \
--from-literal=secret=<your_secret> -n monitoring
```

## Install
```
kubectl apply -f deploy.yaml
```

### Build image

```
docker build -t <NAME>/alert-webhook-receiver:<TAG> .

docker run -d -p 8000:8000 --name webhook <NAME>/alert-webhook-receiver:<TAG> -port 8000 -secret

docker image prune -f
```

