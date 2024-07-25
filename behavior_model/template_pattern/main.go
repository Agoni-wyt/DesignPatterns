package main

import (
	"fmt"
	"template_pattern/template"
)

func main() {

	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &template.Sms{}
	o := template.NewOtp(smsOTP)
	o.GenAndSendOTP(4)

	fmt.Println("")
	emailOTP := &template.Email{}
	o = template.NewOtp(emailOTP)
	o.GenAndSendOTP(4)
}
