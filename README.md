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
curl -XPOST host:9000/webhook/{msgType}/{msgMedium} -d "helloworld"
```

## Config
```
export DING_ROBOT_TOKEN=xxxxxxxxxxxxxxxxx
export DING_ROBOT_SECRET=SECxxxxxxxxxxxxxxxxxx

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
docker image prune -f
```

