package main

import "fmt"

func main() {
	statusApi := "400"
	if statusApi == "200" {
		fmt.Println("Sucesso!")
	} else if statusApi == "500" {
		fmt.Println("Erro no teu código!")
	} else {
		fmt.Println("Bad Request!")
	}
}
