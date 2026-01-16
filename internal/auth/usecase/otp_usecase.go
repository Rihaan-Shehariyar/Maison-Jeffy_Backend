package usecase

import (
	"backend/internal/auth/entity"
	"backend/internal/auth/repository"
	"backend/pkg/email"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type OTPUsecase struct {
	otpRepo repository.OtpRepository
    userRepo repository.UserRepository
}

func NewOTPUsecase(
  otpRepo repository.OtpRepository,
  userRepo repository.UserRepository,
)*OTPUsecase{
  return &OTPUsecase{otpRepo,userRepo}
}



//  SENT OTP

func (u *OTPUsecase) SendOTP(emailaddrs string)error{
  user,err:=u.userRepo.FindByEmail(emailaddrs)

 if err != nil {
	return err
}
if user == nil {
	return errors.New("user not found")
}

 if user.IsVerified{
    return errors.New("User Already Verified")
}

 otp:=GenerateOtp()

 hash,_ := bcrypt.GenerateFromPassword([]byte(otp),bcrypt.DefaultCost)
 
 otpEntity:=entity.OTP{
  Email: emailaddrs,
  CodeHash: string(hash),
  ExpiresAt: time.Now().Add(5 *time.Minute),
}

 if err:=u.otpRepo.Create(&otpEntity);err!=nil{
   return err
}

 return email.SendOTPEmail(emailaddrs,otp)

}


// VERIFY OTP


func(u *OTPUsecase)VerifyOTP(emailaddrs string,otp string)error{
 storedOTP, err := u.otpRepo.FindLatestByEmail(emailaddrs)
if err != nil {
	return errors.New("otp not found")
}
if storedOTP == nil {
	return errors.New("otp not found")
}


 if time.Now().After(storedOTP.ExpiresAt){
   return errors.New("OTP Expired")
}

 if err:=bcrypt.CompareHashAndPassword([]byte(storedOTP.CodeHash),[]byte(otp)) ;err!=nil{
 return  errors.New("Invalid OTP")
} 

 _=u.otpRepo.MarkUsed(storedOTP.ID)
 _=u.userRepo.MarkVerfied(emailaddrs)

 return nil

} 	


func GenerateOtp()string{
  rand.Seed(time.Now().UnixNano())
  return fmt.Sprintf("%06d",rand.Intn(1000000))
}

