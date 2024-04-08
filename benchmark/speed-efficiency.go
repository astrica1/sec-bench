package benchmark

import (
	"fmt"
	"runtime"
	"time"
)

func SpeedEfficiency(hashFunction func(string) string, stringLength ...int) (executionTime time.Duration, memoryUsage runtime.MemStats, hashedValue, result string) {
	length := 1000000
	if len(stringLength) > 0 && stringLength[0] > 100 {
		length = stringLength[0]
	}

	input := RandomString(length)
	startTime := time.Now()
	hash := hashFunction(input)
	elapsedTime := time.Since(startTime)
	computedTime := fmt.Sprintf("Hash computed in %v\n", elapsedTime)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memStatus := fmt.Sprintf("Memory usage: %v bytes\n", m.Alloc)
	hashValue := fmt.Sprintf("Hash value: %s\n", hash)
	return elapsedTime, m, hash, computedTime + memStatus + hashValue
}
