package main

import "fmt"

func main() {
	matriz := [][]string{{"Nome", "Sobrenome", "Hobby favorito"}, {"Leonardo", "Güths", "Viajar de moto"}, {"Nome2", "Sobrenome3", "Hobby2"}, {"Nome3", "Sobrenome3", "Hobby2"}}

	for _, valores := range matriz {
		for _, valor := range valores {
			fmt.Printf("%v\t\t", valor)
		}
		fmt.Println("")
	}
}
