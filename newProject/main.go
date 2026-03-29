package main

import (
	"errors"
	"fmt"
)

func main() {
	// make(array, len, cap)
	q := make([]string, 0, 2)
	q[0] = "q"
	q[1] = "w"
	q = append(q, "1")
	q = append(q, "2")
	fmt.Println(q)




	transactions := []float32{}
	for {
		value, err := scanTransaction()
		if err != nil {
			fmt.Println(err)
			continue
		} else if value == 0 {
			break
		}
		transactions = append(transactions, value)
		fmt.Println(transactions)
	}
	userBalance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", userBalance)

}

func calculateBalance(transactions []float32) float32{
	var userBalance float32 = 0
	for _, value := range transactions {
		userBalance += value
	}
	return userBalance
}

func scanTransaction() (transactionsValue float32, err error) {
	var userTransaction float32
	fmt.Println("Введите сумму транзакции: (0 для выхода)")
	cntScanned, err := fmt.Scan(&userTransaction)
	if cntScanned != 1 || err != nil {
		return 0, errors.New("Введены некорректные данные")
	}
	return userTransaction, err
}
