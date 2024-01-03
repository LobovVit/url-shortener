package storage

import (
	"errors"
)

type Storage struct {
	Relation map[string]string
}

func GetStorage() *Storage {
	return &Storage{make(map[string]string)} //Storage
}

func (ms *Storage) SetRelation(key string, val string) error {
	ms.Relation[key] = val
	return nil
}

func (ms *Storage) GetAll() map[string]string {
	return ms.Relation
}

func (ms *Storage) GetSingle(key string) (string, error) {
	res, ok := ms.Relation[key]
	if ok {
		return res, nil
	} else {
		return "", errors.New("not found")
	}
}
