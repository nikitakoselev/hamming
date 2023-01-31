package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var result = 0

	if len(a) != len(b) {
		return -1, errors.New("inputs must be of the same length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
