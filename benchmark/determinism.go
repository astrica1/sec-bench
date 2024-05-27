package benchmark

func Determinism(hashFunction func(string) string, stringLength ...int) (isDetermined bool, result string) {
	length := 16
	if len(stringLength) > 0 && stringLength[0] > 0 {
		length = stringLength[0]
	}

	input := RandomString(length)
	initialHash := hashFunction(input)
	expectedHash := hashFunction(input)

	if initialHash == expectedHash {
		return true, "The output of function is determined.\n"
	}

	return false, "The output of function is determined.\n"
}
