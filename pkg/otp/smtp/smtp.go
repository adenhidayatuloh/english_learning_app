package smtp

import (
	"english_app/pkg/errs"
	"net/smtp"
)

func SendEmail(recipient, subject, body string) errs.MessageErr {

	SMTPServer := "smtp.gmail.com"
	SMTPPort := "587"
	Username := "learnlingo.id@gmail.com"
	Password := "hwstyjafzvxonxad"
	from := Username
	to := []string{recipient}
	message := []byte("Subject: " + subject + "\n" + body)

	auth := smtp.PlainAuth("", Username, Password, SMTPServer)
	err := smtp.SendMail(SMTPServer+":"+SMTPPort, auth, from, to, message)

	if err != nil {
		return errs.NewBadRequest(err.Error())

	}

	return nil
}

// func sendEmail(recipient, otp string) error {
// 	smtpServer := "smtp.gmail.com"
// 	smtpPort := "587"
// 	username := "learnlingo.id@gmail.com" // Ganti dengan email Anda
// 	password := "hwstyjafzvxonxad"        // Gunakan App Password

// 	from := username
// 	to := []string{recipient}
// 	subject := "Subject: Your One-Time Password (OTP)\n"
// 	body := fmt.Sprintf(""+
// 		"Hello!\n\n"+
// 		"We are excited to have you on board. Your One-Time Password (OTP) is:\n\n"+
// 		"\t\t\t\t\t\t**%s**\n\n"+
// 		"Use this code to complete your registration.\n"+
// 		"Please note: This code is valid for a limited time only!\n\n"+
// 		"If you did not request this, please ignore this email.\n\n"+
// 		"Best regards,\nThe Team", otp)
// 	message := []byte(subject + "\n" + body)

// 	auth := smtp.PlainAuth("", username, password, smtpServer)
// 	return smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
// }
