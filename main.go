package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"rssProxy/handlers"
)

const RSSPROXY_API_PORT_ENV_KEY = "RSSPROXY_API_PORT"
const DEFAULT_API_PORT_VALUE = "5000"

func main() {

	var envPort = initializePort()
	var mux = http.NewServeMux()
	var fs = http.FileServer(http.Dir("assets"))

	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("GET /swagger/", func(w http.ResponseWriter, r *http.Request) {
		var t, _ = template.ParseFiles("views/swagger.html")
		t.Execute(w, nil)
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	mux.HandleFunc("GET /rss", handlers.GetRss)

	log.Default().Println("Listening on :" + envPort)
	http.ListenAndServe(fmt.Sprintf(":%s", envPort), mux)
}

func initializePort() string {
	var envPort, ok = os.LookupEnv(RSSPROXY_API_PORT_ENV_KEY)
	if ok {
		return envPort
	} else {
		return DEFAULT_API_PORT_VALUE
	}
}
