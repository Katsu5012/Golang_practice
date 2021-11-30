package main

import "fmt"

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生", ErrCode: 401}
}
func main() {
	err := RaiseError()
	fmt.Println(err)
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.Message)
	}

}
