package benchmark

import (
	"time"
)

func measureTiming(hashFunction func(string) string, input string) time.Duration {
	start := time.Now()
	_ = hashFunction(input)
	return time.Since(start)
}

func calculateAverage(times []time.Duration) time.Duration {
	var total time.Duration
	for _, t := range times {
		total += t
	}
	return total / time.Duration(len(times))
}

func abs(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}

	return d
}

func SideChannelAttack(hashFunction func(string) string, presision ...int) (atRisk bool, result string) {
	repeat := 10
	if len(presision) > 0 && presision[0] > repeat {
		repeat = presision[0]
	}

	shortTime := make([]time.Duration, repeat)
	for i := range repeat {
		shortTime[i] = measureTiming(hashFunction, RandomString(16))
	}
	midShortTime := calculateAverage(shortTime)

	longTime := make([]time.Duration, repeat)
	for i := range repeat {
		longTime[i] = measureTiming(hashFunction, RandomString(64))
	}
	midLongTime := calculateAverage(longTime)

	threshold := midShortTime * 5 / 100

	if abs(midLongTime-midShortTime) > threshold {
		return true, "Hash function is potentially vulnerable to timing attacks.\n"
	}

	return false, "Hash function is resistant to timing attacks.\n"
}
