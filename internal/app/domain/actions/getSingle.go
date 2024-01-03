package actions

func GetSingleVal(key string) (string, error) {
	return store.GetSingle(key)
}
