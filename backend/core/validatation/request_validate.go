package validatation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/domain/exception"
)

func RequestValidate(obj interface{}, c *gin.Context) *exception.CustomException {
	if err := c.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		e := &exception.CustomException{
			Code:    constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err:     err,
		}
		return e
	}
	return nil
}
