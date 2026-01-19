package usecase

import (
	"backend/internal/auth/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ResetPasswordUsecase struct {
	userRepo repository.UserRepository
    otpRepo repository.OtpRepository
}

func NewResetPasswordUsecase (userRepo repository.UserRepository,otpRepo repository.OtpRepository) *ResetPasswordUsecase{
  return &ResetPasswordUsecase{userRepo: userRepo,otpRepo: otpRepo}
}

func(u *ResetPasswordUsecase)ResetPassword(email,otp,newPassword string)error{

   
  storedOTP,err:=u.otpRepo.FindLatestByEmail(email)
  if err!=nil || storedOTP == nil {
   return errors.New("Invalid Otp")
}

 if time.Now().After(storedOTP.ExpiresAt){
   return errors.New("OTP Expired")
}

if err:=bcrypt.CompareHashAndPassword([]byte(storedOTP.CodeHash),[]byte(otp));err!=nil{
    return errors.New("Invalid Otp")
}

hash,_:=bcrypt.GenerateFromPassword([]byte(newPassword),bcrypt.DefaultCost)

 if err:=u.userRepo.UpdatePassword(email,string(hash));err!=nil{
   return err
}

 _ =u.otpRepo.MarkUsed(storedOTP.ID)

 return nil 
 

}