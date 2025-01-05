package util

func IndexOf(arr []rune, target rune) int {
	for i, r := range arr {
		if r == target {
			return i
		}
	}
	return -1
}
