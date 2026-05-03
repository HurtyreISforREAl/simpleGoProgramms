package main

import (
	"password/account"
	"password/menu"
)

func main() {
	vault := account.NewVault()
Menu:
	for {
		variant := menu.GetMenu()
		switch variant {
		case 1:
			menu.CreateAccount(vault)
		case 2:
			menu.FindAccount(vault)
		case 3:
			menu.DeleteAccount(vault)
		default:
			break Menu
		}
	}
}
