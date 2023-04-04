package infrastructure

import (
    "encoding/base64"
    "net/smtp"
    "strings"
)

type EmailSender struct {
    From            string
    Auth            smtp.Auth
    Host            string
    Port            string
}

func (sender EmailSender) SendMessage(subject, body string, emailAdresses ...string) error {
    address := sender.Host + sender.Port

    message := "To: " + strings.Join(emailAdresses, ",") + "\r\n" +
        "From: " + sender.From + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "Content-Type: text/html; charset=\"UTF-8\"\r\n" +
        "Content-Transfer-Encoding: base64\r\n" +
        "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

    return smtp.SendMail(address, sender.Auth, sender.From, emailAdresses, []byte(message))
}


