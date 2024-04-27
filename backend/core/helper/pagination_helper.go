package helper

import (
	"github.com/unyooon/prompt-sharing/core/constants"
	"github.com/unyooon/prompt-sharing/core/types"
	"github.com/unyooon/prompt-sharing/domain/exception"

	"github.com/gin-gonic/gin"
)

func PaginationHelper[T types.QueryInterface](c *gin.Context, q T, sizeMax int) (T, *exception.CustomException) {
	query := q
	if err := c.ShouldBindQuery(&query); err != nil {
		e := &exception.CustomException{
			Code:    constants.BadRequestCode,
			Message: constants.BadRequestMessage,
			Err:     err,
		}
		return query, e
	}

	if query.GetSize() <= 0 || query.GetSize() > sizeMax {
		query.SetSize(sizeMax)
	}
	if query.GetPage() <= 0 {
		query.SetPage(1)
	}

	return query, nil
}
