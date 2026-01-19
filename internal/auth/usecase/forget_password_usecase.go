package usecase

import (
	"backend/internal/auth/entity"
	"backend/internal/auth/repository"
	"backend/pkg/email"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ForgetPasswordUsecase struct {
	userRepo repository.UserRepository
    otpRepo repository.OtpRepository
}

func NewForgetPasswordUseCase(userRepo repository.UserRepository,otpRepo repository.OtpRepository,) *ForgetPasswordUsecase{
   return &ForgetPasswordUsecase{userRepo: userRepo,otpRepo: otpRepo}
}


func (u *ForgetPasswordUsecase) SendResendOTP(emailAddr string)error{
    user,err := u.userRepo.FindByEmail(emailAddr)
   if err!=nil || user==nil {
	return errors.New("user Not Found")
   }

  otp:=GenerateOtp()

hash,_:= bcrypt.GenerateFromPassword([]byte(otp),bcrypt.DefaultCost)

otpEntity:= entity.OTP{
  Email: emailAddr,
  CodeHash: string(hash),
  ExpiresAt: time.Now().Add(5*time.Minute),
}

 if err:= u.otpRepo.Create(&otpEntity);err!=nil{
   return err
}

 return email.SendOTPEmail(emailAddr,otp)
 
 
}