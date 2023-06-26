package main

import "fmt"

func main() {

	statusApi := "400"

	switch {
	case statusApi == "200":
		fmt.Println("Sucesso!")
	case statusApi == "500":
		fmt.Println("Erro no teu código!")
	default:
		fmt.Println("Bad Request!")
	}
}
