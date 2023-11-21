package Game

import "container/list"

type node struct {
	cost   int
	rank   int
	parent *node
	open   bool
	closed bool
	index  int
}

// func astar(tuile int) list.List {
/* Définition des variables */
//	fileAvoir := list.New()
// listeVus := list.New()

//	fileAvoir.PushBack(tuile)
/* idk */

// }

/* path returns a list of the nodes between F and E */
func path(E node, F node) list.List {
	listRes := list.New()

	listRes.PushFront(F)

	/* Checks if E is the parent of F, otherwise adds the parent of F to the list */
	for check := true; check; check = *F.parent != E {
		listRes.PushFront(F.parent)
		F = *F.parent
	}

	return *listRes
}
