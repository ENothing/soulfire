package jwt

import (
	"gin-init/pkg/config"
	"gin-init/pkg/logging"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Encode()(token string,err error) {

	app, _ := config.Cfg.GetSection("app")

	jwtSecret := []byte(app.Key("JWT_SECRET").String())

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       1,
			"username": "test",
			"nbf":      time.Now().Unix(),
			"iat":      time.Now().Unix(),
		})


	token,err = claims.SignedString(jwtSecret)
	logging.Logging(logging.ERR,err)//记录日志

	return token,err

}