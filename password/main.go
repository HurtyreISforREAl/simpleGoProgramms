package main

import (
	"fmt"

	"password/account"
	"password/files"
)

func main() {
	files.WriteFile("hello i am a new file", "file.txt")
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}
	myAccount.OutputPassword()
	// files.WriteFile()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
