package controller

//使用gomail实现
import (
	"Test/config"
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEamil(EmailAddress string) {
	m := gomail.NewMessage()

	mailConfig := config.GetMailConfig()

	m.SetAddressHeader("From", mailConfig["username"].(string), "lxtbolog")
	m.SetHeader("To", EmailAddress)
	m.SetHeader("Subject", "你已回复")
	m.SetBody("text/html", "回复成功")
	send := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))
	err := send.DialAndSend(m)
	if err != nil {
		fmt.Println("error:", err)
	}

} //发送邮件以此表明已经回复对方邮件，使用go同时进行
