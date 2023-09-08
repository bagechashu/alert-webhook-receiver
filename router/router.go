package router

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/webhook/{msgType}/{msgMedium}", webhookHandler)
	return r
}
