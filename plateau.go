package main

import (
	"errors"
	"math/rand"
)

func pickANumberInATab(tab []int) (int, error) {
	if tab == nil {
		return 0, errors.New("there is no tab in")
	}

	/* Getting the number still available */
	tabLenght := len(tab)
	if tabLenght == 0 {
		return 0, errors.New("tab empty")
	}

	/* Choosing a random num between those still available */
	res := rand.Intn(tabLenght)
}
