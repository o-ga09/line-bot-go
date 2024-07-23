package user

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/user/")
	fmt.Fprintf(w, "Hello userid: %s !", id)
}
