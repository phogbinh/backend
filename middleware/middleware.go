package middleware

import (
	"time"

	. "backend/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
NewAuthMiddleware handles jwt middleware
*/
func NewAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "jwt",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24 * 30,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Login); ok {
				return jwt.MapClaims{
					"mail":     v.Mail,
					"password": v.Password,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return &Login{
				Mail:     claims["mail"].(string),
				Password: claims["password"].(string),
			}
		},
		Authenticator: func(context *gin.Context) (interface{}, error) {
			var login Login
			// Since unlike ShouldBindWith, ShouldBindBodyWith puts back data into context after reading, it is used here so that context data can be re-used twice by the call of authMiddleware.LoginHandler, which checks whether Authenticator is null and then calls the function.
			if err := context.ShouldBindBodyWith(&login, binding.JSON); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			mail := login.Mail
			password := login.Password
			return &Login{Mail: mail, Password: password}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*Login); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

}
