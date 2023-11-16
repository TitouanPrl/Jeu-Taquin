package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/* pickANumberInATab returns a random number in a slice */
func pickANumberInATab(tab []int) ([]int, int, error) {
	if tab == nil {
		return nil, 0, errors.New("there is no tab in")
	}

	/* Getting the number of  numbers still available */
	tabLenght := len(tab)
	if tabLenght == 0 {
		return nil, 0, errors.New("tab empty")
	}

	/* Choosing a random num between those left */
	rand := rand.Intn(tabLenght)
	res := tab[rand]

	/* Deleting this value from the slice and returning res */
	if tabLenght > 1 {
		tab[rand] = tab[tabLenght-1]
		return tab, res, nil
	} else {
		return nil, res, nil
	}
}

/* Displays the playground */
func printPlayground(tab [][]int) error {
	if tab == nil {
		return errors.New("no slice specified")
	}

	lenght := 4
	/* On affiche la barre du haut */
	fmt.Printf("┌")
	for i := 0; i < lenght-2; i++ {
		fmt.Printf("───┬")
	}
	fmt.Printf("───┐\n")

	for i := 0; i < lenght-1; i++ {

		/* Affichage de la ligne */
		for j := 0; j < lenght-1; j++ {
			if tab[i][j] == 0 {
				fmt.Printf("│   ")
			} else {
				fmt.Printf("│ %v ", tab[i][j])
			}
		}
		fmt.Printf("│\n")

		/* Affichage de la ligne séparatrice */
		if i < lenght-2 {
			fmt.Printf("├")
			for k := 0; k < lenght-2; k++ {
				fmt.Printf("───┼")
			}
			fmt.Printf("───┤\n")
		}
	}

	/* Affichage de la dernière ligne */
	fmt.Printf("└")
	for i := 0; i < lenght-2; i++ {
		fmt.Printf("───┴")
	}
	fmt.Printf("───┘\n")

	return nil
}

func convert1Dto2D(tab []int) (*[3][3]int, error) {
	if tab == nil {
		return nil, nil
	}

	const lenght int = 3 /* Size of the tab out */
	var res [lenght][lenght]int
	tmp := 0

	/* Checking if conversion is doable */
	tabLenght := len(tab)
	if (tabLenght*tabLenght)%lenght != 0 {
		return nil, errors.New("wrong entry slice size")
	}

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			res[i][j] = tab[tmp]
			tmp++
		}
	}

	return &res, nil
}
