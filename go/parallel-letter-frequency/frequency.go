// Package letter has function to get frequency of characters
package letter

type FreqMap map[rune]int

// Frequency function to get frequency of characters in string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency function to get frequency of characters in a string concurrently
func ConcurrentFrequency(streams []string) FreqMap {
	c := make(chan FreqMap)
	result := FreqMap{}
	for _, stream := range streams {
		go computeFrequency(c, stream)
	}

	for range streams {
		for k, v := range <-c {
			result[k] += v
		}
	}

	return result
}

// computeFrequency function to feed FreqMap to channel
func computeFrequency(c chan FreqMap, text string) {
	c <- Frequency(text)
}
