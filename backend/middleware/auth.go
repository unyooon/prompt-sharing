package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/core/validatation"
	"github.com/unyooon/prompt-sharing/domain/setting"
)

func ValidateToken(s setting.Setting) gin.HandlerFunc {
	return validatation.ValidateTokenWithGoogle(s)
}
