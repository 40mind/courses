package bootstrap

import (
    "courses/domain/models"
    "courses/infrastructure"
    "net/smtp"
)

func InitEmailSender(config models.Email) infrastructure.EmailSender {
    return infrastructure.EmailSender{
        From:       config.From,
        Auth:       smtp.PlainAuth("", config.From, config.Password, config.Host),
        Host:       config.Host,
        Port:       config.Port,
    }
}
