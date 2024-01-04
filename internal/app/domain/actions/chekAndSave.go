package actions

import (
	"errors"
	"math/rand"
	"strings"
)

func hash(text string) string {
	//hasher := md5.New()
	//hasher.Write([]byte(text))
	//return hex.EncodeToString(hasher.Sum(nil))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func CheckAndSave(value string) (string, error) {

	ret := hash(value)
	err := store.SetRelation(ret, value)
	if err != nil {
		return "", errors.New("bad request")
	}
	return ret, nil
}
