package files

import (
	"fmt"
	"os"
)

func WriteFile(content string, name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Println("Запись прошла успешно")
	return nil
}

func ReadFile(name string) {
	fmt.Println("ok")
}
