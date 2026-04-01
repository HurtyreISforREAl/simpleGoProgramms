package main

import "fmt"

func main() {
	m := map[string]string{
		"PurpleSchool": "https://purpleschool.ru",
	}
	fmt.Println(m)
	fmt.Println(m["PurpleSchool"])
	m["Google"] = "https://google.com"
	m["Yandex"] = "https://yandex.ru"
	fmt.Println(m)
	delete(m, "Google")
	fmt.Println(m)
}
