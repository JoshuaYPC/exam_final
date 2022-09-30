package main

import (
	"eval"
	"fmt"
)

func main() {

	tab := []int{5, 6, 8, 1, 2, 9, 4, 3} // 8 ints
	taille_tab := len(tab)

	missing_nb := eval.Ft_missing(tab, taille_tab)
	fmt.Println("Le nombre manquant dans le tableau est : ", missing_nb)
}
