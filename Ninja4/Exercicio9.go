package main

import "fmt"

func main() {
	mapa := map[string][]string{
		"sobrenome1_nome1": {"hobby1", "hobby2", "hobby3"},
		"sobrenome2_nome2": {"hobby3", "hobby4", "hobby5"},
		"sobrenome3_nome3": {"hobby6", "hobby7", "hobby8"},
	}

	mapa["sobrenome4_nome4"] = []string{"hobby9", "hobby10", "hobby11"}

	for index, value := range mapa {
		fmt.Println(index)
		for i, s := range value {
			fmt.Println("\t", i, " - ", s)
		}
	}
}
