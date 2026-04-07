package main

import (
	"fmt"
	"math/rand/v2"
)

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnm1234567890-_!QWERTYUIOPASDFGHJKLZXCVBNM")

type account struct {
	login    string
	password string
	url      string
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
	// password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login: login,
		// password: password,
		url: url,
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
