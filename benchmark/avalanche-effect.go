package benchmark

import "fmt"

func flipBit(input string, bitPosition int) string {
	bytes := []byte(input)
	byteIndex := bitPosition / 8
	bitIndex := bitPosition % 8
	bytes[byteIndex] ^= 1 << bitIndex
	return string(bytes)
}

func bitDifference(originalHash, modifiedHash string) int {
	diff := 0
	for i := 0; i < len(originalHash); i++ {
		originalByte := originalHash[i]
		modifiedByte := modifiedHash[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << j)
			if (originalByte & mask) != (modifiedByte & mask) {
				diff++
			}
		}
	}
	return diff
}

func measureAvalancheEffect(hashFunction func(string) string, input string) float64 {
	var totalDifference int
	originalHash := hashFunction(input)
	for i := 0; i < len(input)*8; i++ {
		bitFlippedInput := flipBit(input, i)
		modifiedHash := hashFunction(bitFlippedInput)
		totalDifference += bitDifference(originalHash, modifiedHash)
	}

	avalancheEffect := (float64(totalDifference) / (float64(len(input)*8) * float64(len(originalHash)*4))) * 100
	return avalancheEffect
}

func measureAvalancheEffect2(hashFunction func(string) string, input string) float64 {
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

func AvalancheEffect(hashFunction func(string) string, tryCount int, stringLength ...int) (measure float64, result string) {
	length := 16
	if len(stringLength) > 0 && stringLength[0] > length {
		length = stringLength[0]
	}

	var sum = 0.0
	for i := 0; i < tryCount; i++ {
		input := RandomString(length)
		sum += measureAvalancheEffect(hashFunction, input)
	}

	avalancheEffect := sum / float64(tryCount)
	return avalancheEffect, fmt.Sprintf("Avalanche effect of hash algorithm for %d try: %.2f%%\n", tryCount, avalancheEffect)
}
