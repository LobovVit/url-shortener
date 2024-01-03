package handlers

import (
	"github.com/LobovVit/url-shortener/internal/app/domain/actions"
	"io"
	"net/http"
)

func SetShortHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	res, err := actions.CheckAndSave(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write([]byte("http://" + r.Host + "/" + res))
		w.WriteHeader(http.StatusCreated)
	}
}
