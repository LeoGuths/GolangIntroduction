package main

import "fmt"

func main() {
	mapa := map[string][]string{
		"sobrenome1_nome1": {"hobby1", "hobby2", "hobby3"},
		"sobrenome2_nome2": {"hobby1", "hobby2", "hobby3"},
		"sobrenome3_nome3": {"hobby1", "hobby2", "hobby3"},
	}

	for index, value := range mapa {
		fmt.Println(index)
		for i, s := range value {
			fmt.Println("\t", i, " - ", s)
		}
	}
}
