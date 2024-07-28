package json

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/o-ga09/line-bot-go/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	logger.Logger()
	message := struct {
		Code            int
		Message         string
		Updatedatetime  time.Time
		Createddatetime time.Time
	}{
		Code:            200,
		Message:         "Hello",
		Updatedatetime:  time.Now(),
		Createddatetime: time.Now(),
	}

	json, err := json.Marshal(message)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "error marshalling json"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
	slog.Info("success")
}
