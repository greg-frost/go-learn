package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"text/template"
)

// Структура "сообщение"
type Message struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// Структура "данные для подключения"
type Credentials struct {
	Username string
	Password string
	Server   string
	Port     int
}

// Объект шаблона
var t *template.Template

// Текст шаблона
const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}
`

// Инициализация
func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	fmt.Println(" \n[ EMAIL ]\n ")

	// Сообщение
	message := &Message{
		From:    "from@example.com",
		To:      []string{"to@example.com"},
		Subject: "Проверка",
		Body:    "Тестовое письмо, которое никто не получит...",
	}

	// Обработка шаблона
	var body bytes.Buffer
	t.Execute(&body, message)

	// Данные для подключения
	credentials := &Credentials{
		Username: "login",
		Password: "password",
		Server:   "smtp.example.com",
		Port:     25,
	}

	// Подключение к SMTP-серверу
	auth := smtp.PlainAuth("",
		credentials.Username,
		credentials.Password,
		credentials.Server,
	)

	// Текст письма
	fmt.Println(body.String())

	// Отправка письма
	if err := smtp.SendMail(
		credentials.Server+":"+strconv.Itoa(credentials.Port),
		auth,
		message.From,
		message.To,
		body.Bytes(),
	); err != nil {
		log.Fatal(err)
	}

	// Отчет
	log.Print("Сообщение успешно отправлено!")
}
