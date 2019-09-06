package hamming

import "errors"

// Distance calculates the Hamming difference between two Strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The two strings should be of the same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
