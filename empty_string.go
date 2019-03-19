package simple_bench

func EqualEmpty(str string) bool {
	return str == ""
}

func LenZero(str string) bool {
	return len(str) == 0
}

func NotEmpty(str string) bool {
	return str != ""
}

func LenSupZero(str string) bool {
	return len(str) > 0
}
