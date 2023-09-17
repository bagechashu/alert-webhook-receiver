package router

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bagechashu/alert-webhook-receiver/medium"
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
		fmt.Fprint(w, `{"message": "webhook message only support type [raw/prom/huaweismn]."}`)
		log.Printf("error: msgtype doesn't support %s\n", msgType)
		return
	}

	var med medium.Medium
	switch msgMedium {
	case "ding":
		markdown, err := msg.ConvertToDingMarkdown()
		if err != nil {
			fmt.Fprint(w, `{"message": "convert to ding markdown error"}`)
			log.Printf("error: convert data to ding markdown error, %s", err.Error())
			return
		}
		// set msgMedium to dingtalk
		token := os.Getenv("DING_ROBOT_TOKEN")
		secret := os.Getenv("DING_ROBOT_SECRET")
		med = &medium.DingRobot{
			Token:   token,
			Secret:  secret,
			ReqBody: markdown,
		}
	default:
		fmt.Fprint(w, `{"message": "webhook message medium only support [ding]."}`)
		log.Printf("error: wrong message medium: %s \n", msgMedium)
		return
	}

	err := med.Send()
	if err != nil {
		fmt.Fprint(w, `{"message": "send dingtalk error"}`)
		log.Printf("error: send messages error, %s", err.Error())
		return
	}

	fmt.Fprint(w, `{"message": "send successful"}`)
}
