package utils

func ExtractOneFromStringMap(m map[string]string) (string, string, bool) {
	if len(m) != 1 {
		return "", "", false
	}
	for k, v := range m {
		return k, v, true
	}
	return "", "", false
}
