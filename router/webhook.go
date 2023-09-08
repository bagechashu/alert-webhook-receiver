package router

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bagechashu/alert-webhook-receiver/medium/dingtalk"
	"github.com/bagechashu/alert-webhook-receiver/message"
	"github.com/gorilla/mux"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, `{"message": "bad request, only allow POST method"}`)
		log.Printf("error: bad request: [ %+v ]\n", r)
		return
	}

	body, _ := io.ReadAll(r.Body)
	if len(body) == 0 {
		fmt.Fprint(w, `{"message": "post body is empty "}`)
		log.Println("error: post body is empty")
		return
	}
	// log.Print(string(body))

	vars := mux.Vars(r)
	msgType := vars["msgType"]
	msgMedium := vars["msgMedium"]

	var msg message.Message
	switch msgType {
	case "raw":
		msg = message.Raw{Body: body}
	case "prom":
		msg = message.Prom{Body: body}
	case "huaweismn":
		msg = message.HuaweiSMN{Body: body}
	}

	if msg == nil {
		fmt.Fprint(w, `{"message": "webhook message only support type raw/prom/huaweismn"}`)
		log.Printf("error: msgtype doesn't support %s\n", msgType)
		return
	}

	if msgMedium != "ding" {
		fmt.Fprint(w, `{"message": "only support dingtalk"}`)
		log.Printf("error: wrong message medium: %s \n", msgMedium)
		return
	} else {
		markdown := msg.ConvertToDingMarkdown()
		// send messages to dingtalk
		err := dingtalk.Send(markdown)
		if err != nil {
			fmt.Fprint(w, `{"message": "send dingtalk error"}`)
			log.Printf("error: send messages error, %s", err.Error())
			return
		}
	}

	fmt.Fprint(w, `{"message": "send successful"}`)
}
