package usecase

// import "github.com/golang-jwt/jwt/v5"

type RefreshUseCase struct{}

func NewRefreshUseCase() *RefreshUseCase {
	return &RefreshUseCase{}
}

// func (u *RefreshUseCase) Refresh(refreshToken string) (string, error) {

// 	token, err := jwt.Parse()

// }