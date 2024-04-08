package benchmark

import "fmt"

func BirthdayAttack(hashFunction func(string) string, numTrials int, stringLength ...int) (isFound bool, result string) {
	length := 16
	if len(stringLength) > 0 && stringLength[0] > length {
		length = stringLength[0]
	}

	hashSet := make(map[string]string)
	isFound = false
	for i := 0; i < numTrials; i++ {
		input := RandomString(length)
		hash := hashFunction(input)

		if _, ok := hashSet[hash]; ok {
			isFound = true
			return isFound, fmt.Sprintf("Collision found after %d trials\nInput 1: %s\nInput 2: %s\n", i+1, input, hashSet[hash])
		}

		hashSet[hash] = input
	}

	return false, "No collision found within the given number of trials.\n"
}
