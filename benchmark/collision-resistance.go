package benchmark

import "fmt"

func CollisionResistance(hashFunction func(string) string, stringLength int, tryingCount ...int) (isFound bool, result string) {
	numStrings := 10000
	if len(tryingCount) > 0 && tryingCount[0] > numStrings {
		numStrings = tryingCount[0]
	}

	hashMap := make(map[string]string)

	for i := 0; i < numStrings; i++ {
		inputString := RandomString(stringLength)
		hash := hashFunction(inputString)
		if val, ok := hashMap[hash]; ok {
			return true, fmt.Sprintf("Collision found for hashes %s and %s\n", inputString, val)
		} else {
			hashMap[hash] = inputString
		}
	}

	return false, "Collision resistance evaluation complete.\n"
}
