package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
)

func QueryBindHelper(c *gin.Context, obj interface{}) (interface{}, *exception.CustomException) {
	o := obj
	if err := c.ShouldBindQuery(&o); err != nil {
		e := &exception.CustomException{
			Code:    constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err:     err,
		}
		return o, e
	}

	return o, nil
}
