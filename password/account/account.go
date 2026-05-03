package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnm1234567890-_!QWERTYUIOPASDFGHJKLZXCVBNM")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if newAcc.Password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

func (acc Account) Output() {
	color.Cyan(acc.Login)
	color.Red(acc.Password)
	color.Magenta(acc.Url)
}

func (acc *Account) generatePassword(sizePassword int) {
	userPassword := make([]rune, sizePassword)
	lengthLetter := len(letterRunes)
	for i := range userPassword {
		userPassword[i] = letterRunes[rand.IntN(lengthLetter)]
	}
	acc.Password = string(userPassword)
}
