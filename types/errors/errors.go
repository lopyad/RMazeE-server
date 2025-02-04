package errors

import "fmt"

const (
	NotFoundUser  = iota
	InvalidParams = iota
)

var errMessage = map[int64]string{
	NotFoundUser:  "not found user",
	InvalidParams: "invalid params",
}

func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
