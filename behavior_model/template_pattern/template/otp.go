package template

// 模板方法 一次性密码功能(OTP)

type IOtp interface {
	genRandomOTP(int) string       // 生成随机密码
	saveOTPCache(string)           // 缓存该密码以便后续验证
	getMessage(string) string      // 获取消息
	sendNotification(string) error // 发送通知
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  return nil
// }

type Otp struct {
	iOtp IOtp
}

func NewOtp(iOtp IOtp) *Otp {
	return &Otp{
		iOtp: iOtp,
	}
}

func (o *Otp) GenAndSendOTP(otpLength int) error { // 通用的生成和发送OTP方法
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
