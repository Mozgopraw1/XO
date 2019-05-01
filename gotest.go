package main

import (
	"fmt"
	"log"
	"os"
)

const (
	addWord = 1
	addTwo  = 2
	exit    = 3
)

// askVariant asks user that where to go next?
func askVariant() (int, error) {
	fmt.Printf("%d - Добавить слово\n", addWord)
	fmt.Printf("%d - Добавить два :)\n", addTwo)
	fmt.Printf("%d - Свалить\n\n", exit)

	fmt.Print("> ")

	var v int
	_, err := fmt.Scanln(&v)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func main() {
	for {
		v, err := askVariant()
		if err != nil {
			log.Fatalln(err)
		}

		switch v {
		case addWord:
			fmt.Print("> Добавить слово!\n\n")
		case addTwo:
			fmt.Print("> Добавить 2 слова!\n\n")
		case exit:
			fmt.Println("> Пока! 👋")
			os.Exit(0)
		default:
			fmt.Print("> Неизвестнаяя команда...\n\n")
		}
	}
}
