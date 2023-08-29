## Alertmanager Dingtalk Webhook

Alertmanager Webhook for sending alarm message.

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

