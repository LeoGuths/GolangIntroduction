package main

import "fmt"

func main() {
	
	esporteFavorito := "Basquete"

	switch esporteFavorito {
	case "Basquete":
		fmt.Println("Basquete!")
	case "Arqueirismo":
		fmt.Println("Erro no teu código!")
	default:
		fmt.Println("Nenhum!")
	}
}
