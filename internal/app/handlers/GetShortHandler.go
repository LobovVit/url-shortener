package handlers

import (
	"github.com/LobovVit/url-shortener/internal/app/domain/actions"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetShortHandler(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	w.Header().Set("Content-Type", "text/plain")
	res, err := actions.GetSingleVal(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}
}
