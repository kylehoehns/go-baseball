package utils

func Rotate(arr []string, positions int) []string {
	length := len(arr)
	newArr := make([]string, length)
	for i, value := range arr {
		newPosition := (i + positions) % length
		newArr[newPosition] = value
	}
	return newArr
}
