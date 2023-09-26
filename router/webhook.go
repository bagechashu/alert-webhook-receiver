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

	if err := auth(r); err != nil {
		fmt.Fprint(w, `{"message": "bad request, auth failed"}`)
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
	msgType := vars[msgType]
	msgMedium := vars[msgMedium]

	if err := webhookController(msgType, msgMedium, body); err != nil {
		fmt.Fprintf(w, `{"message": "bad request, doesn't support %s or %s"}`, msgType, msgMedium)
		log.Printf("error: %s\n", err)
	}

	fmt.Fprint(w, `{"message": "send successful"}`)
}

// 判断是否需要安全访问
func auth(r *http.Request) (err error) {
	if config.Server.SecretRequest {
		query := r.URL.Query()
		secret := query.Get(secret)
		if secret != config.Server.SecretKey {
			err = fmt.Errorf("bad request, secret key is error")
			return
		}
	}
	return
}

func webhookController(msgType string, msgMedium string, body []byte) (err error) {
	msg, err := message.CreateInterfaceMessage(msgType, body)
	if err != nil {
		return
	}

	med, err := medium.CreateInterfaceMedium(msgMedium, msg)
	if err != nil {
		return
	}

	err = med.Send()
	if err != nil {
		err = fmt.Errorf("send messages error, %s", err.Error())
		return
	}

	return
}
