package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/o-ga09/line-bot-go/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Logger()
	q := r.URL.Query().Get("user")
	if q != "" {
		fmt.Fprintf(w, "<h1>Hello %s from Go!</h1>", q)
	} else {
		fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	}
	slog.Info("success")
}
