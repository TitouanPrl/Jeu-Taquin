package Game

import (
	"errors"
	"fmt"
	"math/rand"
)

/* PrintPlayground displays the playground */
func PrintPlayground(tab *[3][3]int) error {
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

/* convert1Dto2D converts a 1D slice to a 2D slice */
func convert1Dto2D(tab []int) ([3][3]int, error) {
	if tab == nil {
		return [3][3]int{}, nil
	}

	const lenght int = 3 /* Size of the tab out */
	var res [lenght][lenght]int
	tmp := 0

	/* Checking if conversion is doable */
	tabLenght := len(tab)
	if (tabLenght*tabLenght)%lenght != 0 {
		return [3][3]int{}, errors.New("wrong entry slice size")
	}

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			res[i][j] = tab[tmp]
			tmp++
		}
	}

	return res, nil
}

/* SetupInitialPlayground set the playgrounds randomly when we launch the game */
func SetupInitialPlayground() ([3][3]int, error) {
	tmp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	playTab, err := convert1Dto2D(tmp)
	if err != nil {
		return [3][3]int{}, err
	}

	/* Creating a random playground by making doable moves from a sorted playground */
	for i := 0; i < 10000; i++ {
		randNb := rand.Intn(8)
		playTab, err = MoveItem(playTab, randNb)
		if err != nil {
		}
	}

	return playTab, nil

}
