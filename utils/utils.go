package utils

func Contains(source []string, tg string) bool {
	for _, s := range source {
		if s == tg {
			return true
		}
	}
	return false
}
