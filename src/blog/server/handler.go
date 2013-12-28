package server

import (
	"net/http"
	"encoding/json"
)

func DoHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	r.ParseForm()

	out := map[string]string{"name": r.RequestURI}
	b,_ := json.Marshal(out)
	w.Write(b)
}

