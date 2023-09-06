package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bagechashu/alert-webhook-receiver/pkg"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func httpHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "error: bad request")
		return
	}

	body, _ := io.ReadAll(r.Body)
	if len(body) == 0 {
		fmt.Fprint(w, "error: post body is empty ")
		return
	}
	// log.Print(string(body))

	// var notification model.Notification
	// // json to struct
	// if err := json.Unmarshal(body, &notification); err != nil {
	// 	log.Printf("unmarshal data error, %s", err.Error())
	// 	return
	// }

	// send messages to dingtalk
	if err := pkg.SendRaw(string(body)); err != nil {
		log.Printf("send messages error, %s", err.Error())
		return
	}

	fmt.Fprint(w, `{"message": "send to dingtalk successful"}`)
}

func serveHTTP(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen error, %s", err.Error())
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook/ding/raw", httpHandle)
	server := &http.Server{Addr: ":9000", Handler: mux}

	go serveHTTP(server)

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	log.Println("Service is started and listen on ", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
