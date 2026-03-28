package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for {
		userHeight, userWeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			fmt.Println(err)
			continue // or panic("error_exit_comment")
		}
		bodyType := defineBodyType(IMT)
		printBodyType(bodyType)
		outputResult(IMT)

		isRepeateCalculation := checkRepeateOperation()
		if !isRepeateCalculation {
			break
		}
	}
}

func checkRepeateOperation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще расчет? (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice  == "Y" {
		return true
	}
	return false
}

func calculateIMT(weight float64, height float64) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("Введен некорректный вес или рост")
	}
	IMT := weight / math.Pow(height/100, IMTPower)

	return IMT, nil
}

func defineBodyType(IMT float64) string {
	var result string
	switch {
	case IMT < 16:
		result = "isLean"
	case IMT < 18.5:
		result = "isDisadvantage"
	case IMT < 25:
		result = "isNORMAL"
	case IMT < 30:
		result = "isOverage"
	case IMT < 35:
		result = "isFirst"
	case IMT < 40:
		result = "isSecond"
	default:
		result = "isThird"
	}
	return result
}

func printBodyType(bodyType string) {
	switch bodyType {
	case "isLean":
		fmt.Println("У вас дефицит веса")
	case "isDisadvantage":
		fmt.Println("У вас недостаток веса")
	case "isNORMAL":
		fmt.Println("У вас идеальный вес")
	case "isOverage":
		fmt.Println("У вас избыточный вес")
	case "isFirst":
		fmt.Println("У вас первая степень ожирения")
	case "isSecond":
		fmt.Println("У вас вторая степень ожирения")
	default:
		fmt.Println("У вас третья степень ожирения")
	}
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeight float64
	fmt.Println("Введите свой рост в сантиметрах:")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес:")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}

func outputResult(IMT float64) {
	res := fmt.Sprintf("Ваш индекс массы тела: %.2f\n", IMT)
	fmt.Print(res)
}
