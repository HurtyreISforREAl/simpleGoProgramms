package main

import (
	"fmt"
	"demo/app-3/actions"
	"errors"
)

func main() {
	bookmarks := map[string]string{}
	for {
		printMenu()
		userChoice, err := getUserChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}
		toDo := whatToDo(userChoice)
		buf := action(toDo, bookmarks)
		if buf == 3 {
			break
		}
	}
}

func action(toDo string, bookmarks map[string]string) int{
	switch toDo {
	case "Show":
		actions.Show(bookmarks)
	case "Add":
		status, errAdd := actions.Add(bookmarks)
		if status != 0 || errAdd != nil {
			return 1
		}
		return 0
	case "Delete":
		deletedElement := actions.Del(bookmarks)
		if deletedElement["error"] == "error"{
			return 1
		}
		return 0
	case "Complete": 
		return actions.Complete()
	}
	return 0
}

func whatToDo(value int) string {
	switch value {
	case 1:
		return "Show"
	case 2:
		return "Add"
	case 3:
		return "Delete"
	default:
		return "Complete"
	}
}

func getUserChoice() (int, error) {
	var userChoice int
	cntReaden, err := fmt.Scan(&userChoice)
	if cntReaden != 1 || err != nil {
		err = errors.New("Некорректный ввод")
		return 0, err
	}
	return userChoice, err
}

func printMenu() {
	fmt.Println("Меню: Введите номер операции")
	fmt.Println("1 - Просмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку ")
	fmt.Println("4 - Завершить программу")
}
