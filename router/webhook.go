package router

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bagechashu/alert-webhook-receiver/config"
	"github.com/bagechashu/alert-webhook-receiver/medium"
	"github.com/bagechashu/alert-webhook-receiver/message"
	"github.com/gorilla/mux"
)

var (
	secret    string = "secret"
	msgType   string = "msgType"
	msgMedium string = "msgMedium"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, `{"message": "bad request, only allow POST method"}`)
		log.Printf("error: bad request: [ %+v ]\n", r)
		return
	}

	// 判断是否需要安全访问
	if config.Server.SecretRequest {
		query := r.URL.Query()
		secret := query.Get(secret)
		if secret != config.Server.SecretKey {
			fmt.Fprint(w, `{"message": "bad request, secret key is error"}`)
			log.Printf("error: bad request: [ %+v ]\n", r)
			return
		}
	}

	body, _ := io.ReadAll(r.Body)
	if len(body) == 0 {
		fmt.Fprint(w, `{"message": "post body is empty "}`)
		log.Println("error: post body is empty")
		return
	}
	// log.Print(string(body))

	vars := mux.Vars(r)
	msgType := vars[msgType]
	msgMedium := vars[msgMedium]

	result, err := webhookController(msgType, msgMedium, body)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

	fmt.Fprint(w, result)
}

func webhookController(msgType string, msgMedium string, body []byte) (result string, err error) {
	// 根据 msgType 初始化报警源结构体
	var msg message.Message
	switch msgType {
	case "raw":
		msg = message.Raw{Body: body}
	case "prom":
		msg = message.Prom{Body: body}
	case "huaweismn":
		msg = message.HuaweiSMN{Body: body}
	default:
		result = `{"message": "webhook message only support type [raw/prom/huaweismn]."}`
		err = fmt.Errorf("msgtype type error: %s", msgType)
		return
	}

	// 根据 msgMedium 初始化媒体结构体
	var med medium.Medium
	switch msgMedium {
	case "ding":
		var markdown medium.DingTalkMarkdown
		markdown, err = msg.ConvertToDingMarkdown()
		if err != nil {
			result = `{"message": "convert to ding markdown error"}`
			return
		}
		// set msgMedium to dingtalk
		med = &medium.DingRobot{
			Token:   config.DingRobot.Token,
			Secret:  config.DingRobot.Secret,
			ReqBody: markdown,
		}
	default:
		result = `{"message": "webhook message medium only support [ding]."}`
		err = fmt.Errorf("message medium error: %s", msgMedium)
		return
	}

	err = med.Send()
	if err != nil {
		result = `{"message": "send dingtalk error"}`
		err = fmt.Errorf("send messages error, %s", err.Error())
		return
	}

	result = `{"message": "send successful"}`
	return
}
