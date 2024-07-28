package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/o-ga09/line-bot-go/api"
	"github.com/o-ga09/line-bot-go/api/callback"
	"github.com/o-ga09/line-bot-go/api/json"
	"github.com/o-ga09/line-bot-go/api/push"
	"github.com/o-ga09/line-bot-go/api/user"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{name: "case1 (GET /api)", method: "GET", path: "/api", status: http.StatusOK},
		{name: "case2 (GET /api/user/[userId])", method: "GET", path: "/api/user/0000000000", status: http.StatusOK},
		{name: "case3 (GET /api/json)", method: "GET", path: "/api/json", status: http.StatusOK},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("Handler Test %s", tt.name), func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.path, nil)

			switch tt.path {
			case "/api":
				api.Handler(w, r)
			case "/api/user/0000000000":
				user.Handler(w, r)
			case "/api/json":
				json.Handler(w, r)
			case "/api/callback":
				callback.Handler(w, r)
			case "/api/push":
				push.Handler(w, r)
			}

			t.Cleanup(func() {
				r.Body.Close()
			})

			if w.Code != http.StatusOK {
				t.Errorf("expected %d, but got %d", http.StatusOK, w.Code)
			}
		})
	}
}
