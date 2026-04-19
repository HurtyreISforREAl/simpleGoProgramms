package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	
	"github.com/fatih/color"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnm1234567890-_!QWERTYUIOPASDFGHJKLZXCVBNM")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	if login == "" {
		return nil, errors.New("Invalid login")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if newAcc.password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(sizePassword int) {
	userPassword := make([]rune, sizePassword)
	lengthLetter := len(letterRunes)
	for i := range userPassword {
		userPassword[i] = letterRunes[rand.IntN(lengthLetter)]
	}
	acc.password = string(userPassword)
}
