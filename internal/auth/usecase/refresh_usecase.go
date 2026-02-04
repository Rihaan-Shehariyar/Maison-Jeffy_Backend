package usecase

import (
	jwtutils "backend/pkg/jwt_utils"
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshUseCase struct{
}

func NewRefreshUseCase() *RefreshUseCase {
	return &RefreshUseCase{}
}

func (u *RefreshUseCase) Refresh(refreshToken string) (string, error) {

	token, err := jwt.Parse(
    refreshToken,func(t *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("JWT_REFRESH_SECRET")),nil
},
)

  if err!=nil || !token.Valid{
    return "",errors.New("Invalid refresh Token")
}

 claims:=token.Claims.(jwt.MapClaims)
 userId :=(claims["user_id"].(float64))
 role := claims["role"].(string)
 email := claims["email"].(string)
 return jwtutils.GenerateAccessToken(uint(userId),role,email)

 

}