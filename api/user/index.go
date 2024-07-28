package user

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/o-ga09/line-bot-go/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Logger()
	id := strings.TrimPrefix(r.URL.Path, "/api/user/")
	slog.Info(fmt.Sprintf("user id: %s", id))
}
