package array

func Filter[T any](array []T, condition func(T) bool) []T {
	result := make([]T, 0, len(array))
	for _, v := range array {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func AreSlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
