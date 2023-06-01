package handler

import (
	//"api-gateway/pkg/util"
	"errors"
)

// PanicIfUserError 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService.proto--" + err.Error())
		panic(err)
	}
}
