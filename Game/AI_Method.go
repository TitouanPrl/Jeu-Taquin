package Game

import (
	"container/heap"
)

type node struct {
	tab    [3][3]int
	cost   int
	rank   int
	parent *node
	index  int
}

func astar(initialState *node) ([]*node, error) {
	var priorityQueue []*node	/* Nodes to see */
	var alreadySeen []*node	/* Nodes already seen */
	var err error
	win := false
	nbCoup := 0

	priorityQueue = append(priorityQueue, initialState)
	alreadySeen = []*node{}

	initialState.cost = 0
	actualState := initialState

	for len(priorityQueue) > 0 && !win {
		/* Marking the actualState as seen */
		priorityQueue = priorityQueue[:len(priorityQueue)-1]
		alreadySeen = append(alreadySeen, actualState)

		/* Checking if the state is a win state */
		win, err = checkWinCondition(actualState.tab, nbCoup)
		if err != nil {
			return nil, err
		}

		for _, F := range sonsOfNode(actualState) {
			if !(contains(priorityQueue, F) && contains(alreadySeen, F)) || actualState.cost >  {
			}
		}
	}

}

func realCost(E *node) int {

}

func heuristicHammingCost(E *node) int {
	total := len(E.tab) /* Number of cells */
	lenght := 3
	wellPlaced := 0	/* Number of well placed cells */
	tmp := 1 /* Cursor to check if a cell is well placed */

	/* Checking the number of cells already well placed */
	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			/* Checking the 0 at the end */
			if i == 2 && j == 2 && E.tab[i][j] == 0 {
				wellPlaced++
			}
			if E.tab[i][j] == tmp {
				wellPlaced++
			}
		}
	}

	/* Calculating the heuristic cost */
	cost := total - wellPlaced

	return cost
}


/* contains checks if a node is contained in a list */
func contains(list []*node,  item *node) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}

/* sonsOfNode returns a list of the sons of a node */
func sonsOfNode(E *node) []*node {
	var listSons []*node
	var tmp node

	/* Checking the moves possible and adding them to the list of sons */
	for i := 0; i < 9; i++ {
		tab, err := MoveItem(E.tab, i)
		if err == nil {
			tmp.tab = tab
			tmp.parent = E

			listSons = append(listSons, &tmp)
		}
	}

	return listSons
}

/* path returns a list of the nodes between F and E */
func path(E node, F *node) []*node{

	var ListRes []*node
	ListRes = append(ListRes, F)

	/* Checks if E is the parent of F, otherwise adds the parent of F to the list */
	for *F.parent != E {
		ListRes = append(ListRes, F.parent)
		F = F.parent
	}

	/* Reverse the slice to return the right path */
	for i, j := 0, len(ListRes)-1; i < j; i, j = i+1, j-1 {
		ListRes[i], ListRes[j] = ListRes[j], ListRes[i]
	}

	return ListRes
}
}

