package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

var (
	subject  = flag.String("s", "", "subject of the mail")
	body     = flag.String("b", "", "body of themail")
	reciMail = flag.String("m", "", "recipient mail address")
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {
	// Set up authentication information.
	flag.Parse()
	sub := *subject
	//content := *body
	//mailList := strings.Split( *reciMail,",")
	mailList := *reciMail

	smtpServer := "smtp.exmail.qq.com"
	auth := smtp.PlainAuth(
		"",
		"account@yuanbaohui.com",
		"password",
		smtpServer,
	)

	from := mail.Address{"元宝团队", "account@yuanbaohui.com"}
	to := mail.Address{"收件人", mailList}
	title := sub

	body := *body

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	//header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
	//println(message)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":25",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	}
}
