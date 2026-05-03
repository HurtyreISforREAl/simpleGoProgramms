package menu

import (
	"fmt"

	"password/account"

	"github.com/fatih/color"
)

func GetMenu() int {
	var userChoice int
	fmt.Println("Выберите действие:")
	fmt.Println("Создать аккаунт: 1")
	fmt.Println("Найти аккаунт: 2")
	fmt.Println("Удалить аккаунт: 3")
	fmt.Println("Выход: 4")
	fmt.Scan(&userChoice)

	return userChoice
}

func FindAccount(vault *account.Vault) {
	userInputURL := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccountsByURL(userInputURL)
	if len(accounts) == 0 {
		color.Yellow("Аккаутов не найден :(")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func DeleteAccount(vault *account.Vault) {
	userInputURL := promptData("Введите URL для поиска: ")
	isDeleted := vault.DeleteAccountByURL(userInputURL)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func CreateAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
