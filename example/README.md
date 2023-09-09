# Test
```
curl -XPOST http://127.0.0.1:9000/webhook/raw/ding -d "raw string send"

# https://support.huaweicloud.com/intl/en-us/usermanual-smn/smn_ug_a9002.html
curl -XPOST http://127.0.0.1:9000/webhook/huaweismn/ding --data @example/huaweismn-notification.json
curl -XPOST http://127.0.0.1:9000/webhook/huaweismn/ding --data @example/huaweismn-subscription.json

# https://www.puppeteers.net/blog/testing-alertmanager-webhooks-with-curl/
curl -XPOST http://127.0.0.1:9000/webhook/prom/ding --data @example/prom-notification.json

```
