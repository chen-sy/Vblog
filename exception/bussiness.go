package exception

//404 资源不存在或者没有访问权限
func NotExistOrNotPermission(format string, a ...any) *Exception {
	return NewException(404, format, a...)
}

func ValidateError(format string, a ...any) *Exception {
	return NewException(500, format, a...)
}
