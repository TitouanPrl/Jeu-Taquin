package main

import (
	"JeuTaquin/Game"
	"JeuTaquin/Initialization"
	"fmt"
)

func main() {
	/* Ask which solving method to use */
	var input int

	fmt.Println("Choisissez la méthode de résolution :")
	fmt.Println("1 - Manuelle")
	fmt.Println("2 - IA")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	} else if input != 1 && input != 2 {
		panic(fmt.Errorf("Erreur de saisie"))
	}
	fmt.Println("Vous avez choisi la méthode de résolution :", input)

	/* Initializing the first tab and the game configuration */
	playTab, err := Initialization.SetupInitialPlayground()
	if err != nil {
		panic(err)
	}

	if input == 1 {
		err = Game.ManualGame(playTab)
		if err != nil {
			panic(err)
		}
	} else {
		/* EXECUTE IA */
	}

}
