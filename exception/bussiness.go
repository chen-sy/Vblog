package exception

func ValidateError(format string, a ...any) *Exception {
	return NewException(1001, format, a...)
}
