package exception

import "fmt"

func NewException(code int, format string, a ...any) *Exception {
	return &Exception{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

type Exception struct {
	Code    int
	Message string
	Data    any
}

// 实现Error接口
func (e *Exception) Error() string {
	return e.Message
}
