package benchmark

import "fmt"

func measureAvalancheEffect(hashFunction func(string) string, input string) float64 {
	initialHash := hashFunction(input)

	numBits := len(input) * 8
	totalChanges := 0
	for i := 0; i < numBits; i++ {
		modifiedInput := toggleBit(input, i)
		modifiedHash := hashFunction(modifiedInput)
		if modifiedHash != initialHash {
			totalChanges++
		}
	}

	// Calculate and return the avalanche effect percentage
	avalancheEffect := float64(totalChanges) / float64(numBits) * 100
	return avalancheEffect
}

func toggleBit(input string, bitIndex int) string {
	bitIndexInByte := bitIndex % 8
	byteIndex := bitIndex / 8
	bitMask := byte(1 << uint(7-bitIndexInByte))
	byteSlice := []byte(input)
	byteSlice[byteIndex] ^= bitMask
	return string(byteSlice)
}

func AvalancheEffect(hashFunction func(string) string, stringLength ...int) (measure float64, result string) {
	length := 16
	if len(stringLength) > 0 && stringLength[0] > length {
		length = stringLength[0]
	}

	input := RandomString(length)
	avalancheEffect := measureAvalancheEffect(hashFunction, input)
	return avalancheEffect, fmt.Sprintf("Avalanche effect of hash algorithm for input '%s': %.2f%%\n", input, avalancheEffect)
}
