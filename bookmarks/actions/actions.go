package actions

import (
	"fmt"
	"errors"
)
const ok int = 0
const addInput int = 1
const delInput int = 2
const exitProgram int = 3

func Show(bookmarks map[string]string) {
	for key, value := range bookmarks {
		fmt.Println(key, value)
	}
}

func Add(bookmarks map[string]string) (int, error){
	userInputKey, userInputValue, err := getUserInput(addInput)
	if err != nil {
		fmt.Println(err)
		return addInput, err
	}
	bookmarks[userInputKey] = userInputValue
	return ok, err
}

func Del(bookmarks map[string]string) map[string]string {
	userDelKey, _, err := getUserInput(delInput)
	if err != nil {
		fmt.Println(err)
		return map[string]string{"error":"error"}
	}
	deletedElement := map[string]string{userDelKey : bookmarks[userDelKey]}
	delete(bookmarks, userDelKey)

	return deletedElement
}

func Complete() int{
	fmt.Println("Программа завершена :)")
	return exitProgram
}

func getUserInput(variantInput int) (string, string, error) {
	var userKey, userValue string
	if variantInput == addInput {
		fmt.Println("Введите ключ, по которому будет добавлено значение:")
		cntReaden, err := fmt.Scan(&userKey)
		if cntReaden != 1 || err != nil {
			err := errors.New("Ошибка ввода")
			return "", "", err
		} 
		fmt.Println("Введите значение:")
		cntReaden1, err1 := fmt.Scan(&userValue)
		if cntReaden1 != 1 || err1 != nil {
			err1 := errors.New("Ошибка ввода")
			return "", "", err1
		}
		return userKey, userValue, err1
	} else {
		fmt.Println("Введите ключ, по которому хотите удалить значение:")
		cntReaden, err := fmt.Scan(&userKey)
		if cntReaden != 1 || err != nil {
			err := errors.New("Ошибка ввода")
			return "", "", err
		}
		return userKey, userValue, err
	}
}