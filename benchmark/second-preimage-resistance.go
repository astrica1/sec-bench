package benchmark

import "fmt"

func SecondPreimageResistance(hashFunction func(string) string, stringLength ...int) (isFound bool, result string) {
	length := 10
	if len(stringLength) > 0 && stringLength[0] > length {
		length = stringLength[0]
	}

	initialInput := RandomString(length)
	initialHash := hashFunction(initialInput)

	isFound = false
	attempts := 0
	for !isFound {
		attempts++
		newInput := RandomString(length)
		newHash := hashFunction(newInput)
		if newHash == initialHash && newInput != initialInput {
			isFound = true
			return isFound, fmt.Sprintf("Second preimage found after %d attempts.\nInitial input string: %s\nNew input string: %s\n", attempts, initialInput, newInput)
		}
	}

	return false, "Preimage resistance evaluation complete.\n"
}
