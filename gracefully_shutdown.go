package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &hiHandler{})

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit

		if err := server.Close(); err != nil {
			log.Fatal("Close Server: ", err)
		}
	}()

	log.Println("Starting HTTP Server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server Closed under request")
		} else {
			log.Fatal("Server Closed unexpected")
		}
	}
}

type hiHandler struct {
}

func (_ *hiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
