package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rssProxy/handlers"
	"text/template"
)

const RSSPROXY_API_PORT_ENV_KEY = "RSSPROXY_API_PORT"
const DEFAULT_API_PORT_VALUE = "5000"

//go:embed assets views
var staticAssets embed.FS

func main() {

	var envPort = initializePort()
	var mux = http.NewServeMux()
	var fileServer = http.FileServer(http.FS(staticAssets))

	mux.Handle("/assets/", fileServer)

	mux.HandleFunc("GET /swagger/", func(w http.ResponseWriter, r *http.Request) {
		var t, _ = template.ParseFS(staticAssets, "views/swagger.html")
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
