package utils

func FirstSliceItem[T any](slice []T) (result T) {
	if len(slice) == 0 {
		return
	}
	return slice[0]
}
