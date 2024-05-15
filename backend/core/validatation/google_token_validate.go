package validatation

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unyooon/prompt-sharing/domain/setting"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type Google struct {
	Config *oauth2.Config
}

func ValidateTokenWithGoogle(s setting.Setting) gin.HandlerFunc {
	google := &Google{
		Config: &oauth2.Config{
			ClientID:     s.GoogleClientId,
			ClientSecret: s.GoogleClientSecret,
			Endpoint:     googleOAuth.Endpoint,
			Scopes:       []string{"openid"},
			RedirectURL:  "http://localhost:3000/api/auth/callback/google",
		},
	}

	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return func(c *gin.Context) {
		// jwt取得
		bearer := c.GetHeader("Authorization")
		if len(bearer) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}
		j := strings.Fields(bearer)[1]

		// Token検証準備
		ctx := context.Background()
		httpClient := google.Config.Client(ctx, &oauth2.Token{AccessToken: j})
		service, err := v2.NewService(ctx, option.WithHTTPClient(httpClient))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		// Token検証 & ユーザ情報取得
		userInfo, err := service.Tokeninfo().AccessToken(j).Context(ctx).Do()
		if err != nil {
			log.Print(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.New("invalid token").Error(),
			})
			return
		}

		c.Set("userId", userInfo.UserId)
		c.Set("mailAddress", userInfo.Email)
	}
}
