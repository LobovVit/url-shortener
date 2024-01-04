package actions

import "github.com/LobovVit/url-shortener/internal/app/domain/storage"

var store = StorInterface(storage.GetStorage())

type StorInterface interface {
	SetRelation(key string, val string) error
	GetAll() map[string]string
	GetSingle(key string) (string, error)
}
