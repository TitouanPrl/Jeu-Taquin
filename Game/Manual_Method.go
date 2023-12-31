package Game

import (
	"errors"
	"fmt"
)

func ManualGame(playTab [3][3]int) error {

	nbCoup := 1
	var win bool
	var err error

	for {
		CallClearTerminal()

		err = PrintPlayground(&playTab)
		if err != nil {
			return err
		}

		fmt.Println("Coup numéro ", nbCoup)

		playTab, win, err = roundManual(playTab, nbCoup)
		if err != nil {
			return err
		}

		if win == true {
			err := PrintPlayground(&playTab)
			if err != nil {
				return err
			}
			break
		}

		nbCoup++
	}

	return nil
}

/* roundManual asks the player to make a move until the move is correct */
func roundManual(playTab [3][3]int, nbCoup int) ([3][3]int, bool, error) {
	nextTab := [3][3]int{}

	/* Asks which cell to move until it's a valid play */
	for next := true; next; next = nextTab == [3][3]int{} {

		/* Asking the cell */
		cellToPlay, err := askCellToPlay()
		if err != nil {
			return [3][3]int{}, false, err
		}

		/* Making the move */
		nextTab, err = MoveItem(playTab, cellToPlay)
		if nextTab == [3][3]int{} {
			fmt.Println("La case choisie n'est pas un coup possible, veuillez réessayer")
			fmt.Println("Erreur : ", err)
		}
	}

	/* Checking win condition */
	win, err := checkWinCondition(nextTab, nbCoup)
	if err != nil {
		return [3][3]int{}, false, err
	}

	return nextTab, win, nil
}

/* askCellToPlay ask which cell the player wants to move */
func askCellToPlay() (int, error) {
	var input int

	fmt.Printf("\nQuel tuile souhaitez-vous déplacer ? \n")

	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	} else if input <= 0 || input > 8 {
		return 0, errors.New("erreur de saisie")
	}
	fmt.Printf("Vous avez choisi la tuile %v \n\n", input)

	return input, nil
}
