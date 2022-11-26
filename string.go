package aocutil

func IsHexChar[T byte | rune](c T) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'F') ||
		(c >= 'a' && c <= 'f')
}

func IsHexString(s string) bool {
	for _, c := range s {
		if !IsHexChar(c) {
			return false
		}
	}
	return true
}
