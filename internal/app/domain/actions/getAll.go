package actions

func GetAll() (map[string]string, error) {
	return store.GetAll(), nil
}
