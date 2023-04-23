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

func FindAll[T any](s []T, fn func(T) bool) []T {
	items := make([]T, 0)

	for i := range s {
		if fn(s[i]) {
			items = append(items, s[i])
		}
	}

	return items
}

func Contains[T any](s []T, fn func(T) bool) bool {
	for i := range s {
		if fn(s[i]) {
			return true
		}
	}

	return false
}

func IndexOf[T any](s []T, fn func(T) bool) int {
	for i := range s {
		if fn(s[i]) {
			return i
		}
	}

	return -1
}
