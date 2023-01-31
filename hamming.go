package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var result = 0

	if len(a) != len(b) {
		return 0, errors.New("strings' length is different")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result += 1
		}
	}

	return result, nil
}
