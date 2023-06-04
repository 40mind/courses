package infrastructure

import (
    "encoding/base64"
    "net/smtp"
)

type EmailSender struct {
    From            string
    Auth            smtp.Auth
    Host            string
    Port            string
}

func (sender EmailSender) SendMessage(subject, course, fio, firstDate string, emailAdress ...string) error {
    address := sender.Host + sender.Port

    message := "To: " + emailAdress[0] + "\r\n" +
        "From: " + sender.From + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "Content-Type: text/html; charset=\"UTF-8\"\r\n" +
        "Content-Transfer-Encoding: base64\r\n" +
        "\r\n" + base64.StdEncoding.EncodeToString([]byte(`<!DOCTYPE html>
    <html>
    <head>
        <meta charset="UTF-8">
        <title>Подтверждение записи на курс</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                background-color: #f7f7f7;
            }
    
            h1 {
                color: #333333;
                font-size: 24px;
            }
    
            h2 {
                color: #666666;
                font-size: 20px;
            }
    
            p {
                color: #666666;
                font-size: 16px;
                line-height: 1.5;
            }
    
            .container {
                max-width: 600px;
                margin: 0 auto;
                padding: 20px;
                background-color: #ffffff;
                border-radius: 4px;
                box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            }
    
            .footer {
                margin-top: 20px;
                text-align: center;
            }
        </style>
    </head>
    <body>
    <div class="container">
        <h1>Успешная запись на курс</h1>
        <p>Уважаемый/ая ` + fio + `,</p>
        <p>Мы рады сообщить вам, что вы успешно записались на курс:</p>
        <h2>` + course + `</h2>
        <p>Первое занятие состоится ` + firstDate + `.</p>
        <p>Мы ожидаем вас с нетерпением и надеемся, что курс будет интересным и познавательным.</p>
        <p>Если у вас возникнут вопросы или требуется дополнительная информация, пожалуйста, свяжитесь с нами.</p>
        <div class="footer">
        <p>С уважением,</p>
        <p>Команда сайта energy education</p>
    </div>
    </div>
    </body>
    </html>`))

    return smtp.SendMail(address, sender.Auth, sender.From, emailAdress, []byte(message))
}
