package main

import (
	"fmt"
	"log"
)

// Run - функция создает и запускает экземпляр приложения
func Run() error {
	fmt.Println("Старт экземпляра приложения")
	return nil
}

func main() {
	fmt.Println("Курс Go REST API")
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
