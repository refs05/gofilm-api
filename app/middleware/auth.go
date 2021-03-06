package middleware

import (
	"gofilm/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
	ExpiresDuration int
}

func (jwtConfig *ConfigJWT) Init() middleware.JWTConfig {
	
	return middleware.JWTConfig{
		Claims: &JwtCustomClaims{},
		SigningKey: []byte(jwtConfig.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusForbidden, e) 
		}),
	}
}

func (jwtConfig *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		"admin",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConfig.SecretJWT))

	return token
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}