package api

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("user")
	if q != "" {
		fmt.Fprintf(w, "<h1>Hello %s from Go!</h1>", q)
	} else {
		fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	}
}
