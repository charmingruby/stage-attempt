package helpers

func IsItemInStringMap(m map[string]string, v string) bool {
	for _, item := range m {
		if item == v {
			return true
		}
	}

	return false
}
