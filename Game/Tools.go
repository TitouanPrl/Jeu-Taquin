package Game

import (
	"errors"
)

/* abs returns the absolute value of x */
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/* moveItem swap an item with the null item if they're neighbours */
func moveItem(tab [3][3]int, item int) ([3][3]int, error) {
	if item <= 0 {
		return [3][3]int{}, errors.New("impossible de trouver une case négative")
	}
	lenght := 3
	itemX := -1
	itemY := -1
	zeroX := -1
	zeroY := -1

	/* Looking for the item and the empty space in the slice */
	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if tab[i][j] == item {
				itemX = j
				itemY = i
			} else if tab[i][j] == 0 {
				zeroX = j
				zeroY = i
			}
		}
	}

	/* Checking if values were found */
	if itemX == -1 || itemY == -1 || zeroX == -1 || zeroY == -1 {
		return [3][3]int{}, errors.New("les valeurs recherchées n'ont pas été trouvées")
	}

	/* Checking if they're neighbours */
	deltaX := abs(itemX - zeroX)
	deltaY := abs(itemY - zeroY)

	if !(deltaX == 1 && deltaY == 0) || !(deltaX == 0 && deltaY == 1) {
		return [3][3]int{}, errors.New("les cases ne sont pas voisines")
	}

	/* Swapping items */
	tab[itemY][itemX] = tab[zeroY][zeroX]
	tab[zeroY][zeroX] = item

	return tab, nil
}
