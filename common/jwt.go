package common

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const jwtSecret = "cly—test"

//Claims ...
type Claims struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(username, password string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		username,
		password,
		authority,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cmall",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTSECRET)

	return token, err
}

/*认证token*/
func CheckToken(token string) (*Claims, error) {
	claims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSECRET, nil
	})
	if claims != nil {
		if c, ok := claims.Claims.(*Claims); ok && claims.Valid {
			return c, nil
		}
		return nil, err
	}
	return nil, err
}
