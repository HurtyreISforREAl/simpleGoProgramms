package main

import (
	"fmt"
)

func main() {
	transactions := [3]int{8, 9, 7}
	banks := [2]string{}

	fmt.Println(transactions[1])
	banks[1] = "Yandex"
	fmt.Println(banks)
}
