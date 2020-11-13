package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	/*http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})*/

	//http.HandleFunc("/", hello)

	//http.Handle("/", &helloHandler{})

	// 本质上就是一个带有路由层的 http.Handler 具体实现，并以此为基础提供大量便利的辅助方法。
	/*mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))*/

	/*mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})

	server := &http.Server{
		Addr: ":4000",
		Handler: mux,
	}
	log.Println("Starting HTTP server...")
	log.Fatal(server.ListenAndServe())*/

	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	mux.HandleFunc("/timeout", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(2 * time.Second)
		// (52) Empty reply from server
		writer.Write([]byte("Timeout"))
	})

	server := &http.Server{
		Addr:         ":4001",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Starting HTTP server...")
	log.Fatal(server.ListenAndServe())

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type helloHandler struct {
}

func (_ *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
