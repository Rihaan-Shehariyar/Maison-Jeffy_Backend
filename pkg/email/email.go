package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(to string, otp string) error {

	from := os.Getenv("SMTP_EMAIL")
    password:=os.Getenv("SMTP_PASSWORD")

  msg:=fmt.Sprintf(
   "Subject: Email Verification OTP \n\n Your OTP is : %s\n Valid For 5 minutes",otp ,
)

 auth :=smtp.PlainAuth("",from,password,"smtp.gmail.com")

 return smtp.SendMail(
  "smtp.gmail.com:587",
   auth,
   from,
   []string{to},
   []byte(msg),
)
}