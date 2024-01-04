package handlers

import (
	"github.com/LobovVit/url-shortener/internal/app/config"
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
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://" + config.Instance.HostB + "/" + res))
	}
}
