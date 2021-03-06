package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"qcode/models"
	"time"
)

var jwtKey =[]byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user models.User)(string ,error)  {
	expirationTime := time.Now().Add(30*time.Minute)
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer: "sjy",
			IssuedAt: time.Now().Unix(),
			Subject: "user_token",

		},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstring ,err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenstring ,nil

}

func ParseToken(tokenstring string)(*jwt.Token,*Claims,error) {
	claims := &Claims{}

	token ,err := jwt.ParseWithClaims(tokenstring,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey ,nil
	})
	return token,claims,err
}

func ReleaseTokenRoot(user models.User)(string ,error)  {
	expirationTime := time.Now().Add(30*time.Minute)
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer: "sjy_root",
			IssuedAt: time.Now().Unix(),
			Subject: "root_user_token",

		},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenstring ,err := token.SignedString(jwtKey)
	if err != nil {
		return "",err
	}
	return tokenstring ,nil

}