package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnm1234567890-_!QWERTYUIOPASDFGHJKLZXCVBNM")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	if login == "" {
		return nil, errors.New("Invalid login")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if newAcc.password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(sizePassword int) {
	userPassword := make([]rune, sizePassword)
	lengthLetter := len(letterRunes)
	for i := range userPassword {
		userPassword[i] = letterRunes[rand.IntN(lengthLetter)]
	}
	acc.password = string(userPassword)
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
