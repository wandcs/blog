package main

import (
	"net/http"
	"blog/server"
	"os"
)

func main() {
	http.HandleFunc("/do/", server.DoHttp)

	http.ListenAndServe(":9000", nil)
	os.Exit(0);
}
