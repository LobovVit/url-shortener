package storage

import (
	"errors"
)

type Storage struct {
	relation map[string]string
}

func GetStorage() *Storage {
	return &Storage{make(map[string]string)} //Storage
}

func (ms *Storage) SetRelation(key string, val string) error {
	ms.relation[key] = val
	return nil
}

func (ms *Storage) GetAll() map[string]string {
	return ms.relation
}

func (ms *Storage) GetSingle(key string) (string, error) {
	res, ok := ms.relation[key]
	if ok {
		return res, nil
	} else {
		return "", errors.New("not found")
	}
}
