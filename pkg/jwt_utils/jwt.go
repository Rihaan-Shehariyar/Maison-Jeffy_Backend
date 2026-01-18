package jwtutils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}


func GenerateAccessToken(userId uint,email string)(string,error){

  secret:=os.Getenv("JWT_SECRET")
 if secret ==""{
   return "",errors.New("JWT Secret Not set")
}

 claims := Claims{
  UserID: userId,
  Email: email,
 RegisteredClaims: jwt.RegisteredClaims{
   IssuedAt: jwt.NewNumericDate(time.Now()),
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(15*time.Minute)),
},
}

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString([]byte(secret))

} 


func GenerateRefreshToken(email string,userID uint)(string,error){
 
 secret := os.Getenv("JWT_REFRESH_SECRET")
 if secret ==""{
  return "", errors.New("JWT Secret Not Set")
}

 claims:=Claims{
  UserID: userID,
  Email: email,
  RegisteredClaims:jwt.RegisteredClaims{
    IssuedAt: jwt.NewNumericDate(time.Now()),
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*24*time.Hour)),
},
}

 token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString([]byte(secret))

}
