package handler

import (
	"log/slog"
	"net/http"

	"github.com/CleysonPH/dreampic.ai/types"
)

func getAuthenticatedUser(r *http.Request) types.AuthenticatedUser {
	user, ok := r.Context().Value(types.UserContextKey).(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{}
	}
	return user
}

func Make(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "error", err, "path", r.URL.Path)
		}
	}
}
