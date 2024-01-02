package services

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/vinitparekh17/project-x/config"
	"github.com/vinitparekh17/project-x/handler"
	"github.com/vinitparekh17/project-x/models"
)

type EmailData struct {
	Body    string `json:"body"`
	To      string `json:"destinationEmail"`
	Subject string `json:"subject"`
}

type User models.IdentityModel

func (u *User) Mail(emailData EmailData) error {
	err := make(chan error)
	go func() {
		err <- RequestMail(emailData)
	}()
	return <-err
}

func RequestMail(emailData EmailData) error {
	res, err :=
		http.Post(
			fmt.Sprintf("%s/sendmail", config.K.String("notification_url")),
			"application/json",
			strings.NewReader(`{
			"body": "`+emailData.Body+`",
			"destinationEmail": "`+emailData.To+`",
			"subject": "`+emailData.Subject+`
		"}`))

	handler.ErrorHandler(err)
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

	if res.StatusCode != 200 {
		return errors.New("error occured while sending email")
	}

	return nil
}

func GenerateOTP() string {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	otpLength := 6
	var otp string
	for i := 0; i < otpLength; i++ {
		otp += fmt.Sprint(rand.Intn(10))
	}
	return otp
}
