package Game

import (
	"time"
)

type node struct {
	tab    [3][3]int
	cost   int
	val    int
	parent *node
}

func IAGame(playTab [3][3]int) error {
	var initialState node
	initialState.tab = playTab

	answer, err := astar(&initialState)
	if err != nil {
		return err
	}

	for i := 0; i < len(answer); i++ {
		CallClearTerminal()
		err = PrintPlayground(&answer[i].tab)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func astar(initialState *node) ([]*node, error) {
	var priorityQueue []*node /* Nodes to see */
	var alreadySeen []*node   /* Nodes already seen */
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

		for _, F := range sonsOfNode(actualState) {
			/* If the node doesn't already exist and is relevant, we create it */
			if !(contains(priorityQueue, F) || contains(alreadySeen, F)) {
				F.cost = pathCost(initialState, actualState) + pathCost(actualState, F)
				F.val = F.cost + heuristicHammingCost(F)
				F.parent = actualState

				/* Inserting and sorting the node in the tab */
				priorityQueue = insertInSorted(priorityQueue, F)
			}
		}

		/* If the queue isn't empty, we check the new head */
		if len(priorityQueue) > 0 {
			actualState = priorityQueue[len(priorityQueue)-1]
		}
		/* If the state is final, then we return it */
		win, err = checkWinCondition(actualState.tab, nbCoup)
		if err != nil {
			return nil, err
		}
		if win {
			return path(initialState, actualState), nil
		}
	}

	/* Random return */
	return nil, nil
}

/* insertInSorted insert an element at the right place in an array */
func insertInSorted(sortedArr []*node, elmt *node) []*node {
	/* Looking for the right index to insert the node */
	i := 0
	for i < len(sortedArr) && sortedArr[i].val > elmt.val {
		i++
	}

	/* Insert the node at the index */
	sortedArr = append(sortedArr[:i], append([]*node{elmt}, sortedArr[i:]...)...)

	return sortedArr
}

func heuristicHammingCost(E *node) int {
	total := len(E.tab) /* Number of cells */
	lenght := 3
	wellPlaced := 0 /* Number of well-placed cells */
	tmp := 1        /* Cursor to check if a cell is well-placed */

	/* Checking the number of cells already well-placed */
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
func contains(list []*node, item *node) bool {
	for _, a := range list {
		if a.tab == item.tab {
			return true
		}
	}
	return false
}

/* sonsOfNode returns a list of the sons of a node */
func sonsOfNode(E *node) []*node {
	var listSons []*node

	/* Checking the moves possible and adding them to the list of sons */
	for i := 0; i < 9; i++ {
		tab, err := MoveItem(E.tab, i)
		if err == nil {
			listSons = append(listSons, &node{tab: tab, parent: E})
		}
	}

	return listSons
}

/* path returns a list of the nodes between F and E */
func path(E, F *node) []*node {

	var ListRes []*node
	ListRes = append(ListRes, F)

	/* Checks if E is the parent of F, otherwise adds the parent of F to the list */
	for F.parent != E {
		ListRes = append(ListRes, F.parent)
		F = F.parent
	}

	/* Reverse the slice to return the right path */
	for i, j := 0, len(ListRes)-1; i < j; i, j = i+1, j-1 {
		ListRes[i], ListRes[j] = ListRes[j], ListRes[i]
	}

	return ListRes
}

/* pathCost returns the cost from one node to another */
func pathCost(E, F *node) int {
	if E == F {
		return 0
	}

	cost := 1

	for F.parent != E {
		cost++
		F = F.parent
	}

	return cost
}
