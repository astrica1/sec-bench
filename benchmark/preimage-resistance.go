package benchmark

import (
	"fmt"
	"strings"
)

func PreimageResistance(hashFunction func(string) string, targetHash string, stringLength ...int) (isFound bool, result string) {
	if targetHash == "" {
		l := len(hashFunction(RandomString(10)))
		targetHash = strings.Repeat("0", l)
	}
	length := 10
	if len(stringLength) > 0 && stringLength[0] > length {
		length = stringLength[0]
	}

	isFound = false
	attempts := 0
	for !isFound {
		attempts++
		inputString := RandomString(length)
		hash := hashFunction(inputString)
		if hash == targetHash {
			isFound = true
			return isFound, fmt.Sprintf("Preimage found after %d attempts.\nInput string: %s\n", attempts, inputString)
		}
	}

	return false, "Preimage resistance evaluation complete.\n"
}
