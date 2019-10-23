package jwt

import (
	"gin-init/pkg/config"
	"gin-init/pkg/logging"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type UserToken struct {
	Id         int64
	Username   string
	Openid     string
	SessionKey string
}

var (
	app, _    = config.Cfg.GetSection("app")
	jwtSecret = []byte(app.Key("JWT_SECRET").String())
)

func Encode(userToken UserToken) (token string, err error) {

	expriesTime, _ := strconv.ParseInt(app.Key("TOKEN_EXPIRESAT").String(), 10, 64)
	expriesTime = time.Now().Unix() + expriesTime

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":          userToken.Id,
			"aud":         userToken.Username,
			"openid":      userToken.Openid,
			"session_key": userToken.SessionKey,
			"iat":         time.Now().Unix(),
			"nbf":         time.Now().Unix(),
			"exp":         expriesTime,
			"iss":         "soulfire",
		})

	token, err = claims.SignedString(jwtSecret)

	logging.Logging(logging.ERR, err) //记录日志

	return token, err

}

func Parse(token string) (*UserToken,error) {

	userToken := &UserToken{}

	keyFunc := func(t *jwt.Token) (i interface{}, e error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret),nil
	}


	tokenParse, err := jwt.Parse(token,keyFunc)

	if err != nil {

		return userToken,err

	}else if claims,ok := tokenParse.Claims.(jwt.MapClaims);ok && tokenParse.Valid {

		userToken.Id = int64(claims["id"].(float64))
		userToken.Username = claims["aud"].(string)
		userToken.Openid = claims["openid"].(string)
		userToken.SessionKey = claims["session_key"].(string)

		return userToken,nil

	}else{

		logging.Logging(logging.ERR, err) //记录日志

		return userToken,err

	}

}
