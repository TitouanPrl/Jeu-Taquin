package main

import "fmt"

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

	if input == 1 {
		/* EXECUTE MAN */
	} else {
		/* EXECUTE IA */
	}
	/*test := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	} */

	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	res, err := convert1Dto2D(test)
	if err != nil {
		panic(err)
	}

	fmt.Println(*res)
}
