package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bagechashu/alert-webhook-receiver/config"
	"github.com/bagechashu/alert-webhook-receiver/router"
)

func serveHTTP(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: server error [ %s ]\n", err.Error())
	}
}

func main() {
	router := router.InitRouter()

	addr := ":" + strconv.Itoa(config.Server.Port)
	server := &http.Server{Addr: addr, Handler: router}

	go serveHTTP(server)

	// Give the server some time to start
	time.Sleep(1 * time.Second)

	log.Println("info: service is started and listen on ", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
}
