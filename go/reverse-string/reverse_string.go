// Package reverse has methods to reverse a string
package reverse

// String returns a string in reverse order
func String(word string) string {

	if word == "" {
		return ""
	}
	result := []rune(word)
	length := len(result)
	for idx, c := range result {
		if idx == length/2 {
			break
		}
		temp := result[length-idx-1]
		result[length-idx-1] = c
		result[idx] = temp

	}

	return string(result)
}
