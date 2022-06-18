package utils

func Ternary[K any](condition bool, onSuccess, onFailure K) K {
	if condition {
		return onSuccess
	} else {
		return onFailure
	}
}
