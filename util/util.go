package util

func Find[T any](s []T, fn func(T) bool) T {
	for i := range s {
		if fn(s[i]) {
			return s[i]
		}
	}

	var zeroedV T
	return zeroedV
}

func Contains[T any](s []T, fn func(T) bool) bool {
	for i := range s {
		if fn(s[i]) {
			return true
		}
	}

	return false
}
