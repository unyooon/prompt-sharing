package exception

import (
	"github.com/unyooon/prompt-sharing/core/constants"
	"go.uber.org/zap"
)

type CustomException struct {
	Code       constants.Code
	Message    constants.Message
	Err        error
	StackTrace string
}

func NewCustomException(c constants.Code, m constants.Message, e error) *CustomException {
	stack := zap.Stack("").String
	return &CustomException{
		Code:       c,
		Message:    m,
		Err:        e,
		StackTrace: stack,
	}
}
