package main


func GetHeight(letter string) int {
	var letterAsRune rune
	switch letter {
	case "a", "S":
		letterAsRune = []rune("a")[0]
	case "z", "E":
		letterAsRune = []rune("z")[0]
	default:
		letterAsRune = []rune(letter)[0]
	}
	return int(letterAsRune) - 96
}

func SetNeighbours(current *Node, nodeList *[]Node, inputMatrix [][]string) {
	// top
	if current.Row-1 >= 0 {
		for i := range *nodeList {
			if ((*nodeList)[i].Row == current.Row-1 && (*nodeList)[i].Col == current.Col) &&
				((*nodeList)[i].Height-current.Height <= 1) {
				current.Neighbours = append(current.Neighbours, &(*nodeList)[i])
				break
			}
		}
	}
	// bottom
	if current.Row+1 <= len(inputMatrix) {
		for i := range *nodeList {
			if ((*nodeList)[i].Row == current.Row+1 && (*nodeList)[i].Col == current.Col) &&
				((*nodeList)[i].Height-current.Height <= 1) {
				current.Neighbours = append(current.Neighbours, &(*nodeList)[i])
				break
			}
		}
	}
	// left
	if current.Col-1 >= 0 {
		for i := range *nodeList {
			if ((*nodeList)[i].Col == current.Col-1 && (*nodeList)[i].Row == current.Row) &&
				((*nodeList)[i].Height-current.Height <= 1) {
				current.Neighbours = append(current.Neighbours, &(*nodeList)[i])
				break
			}
		}
	}
	// right
	if current.Col+1 <= len(inputMatrix[0]) {
		for i := range *nodeList {
			if ((*nodeList)[i].Col == current.Col+1 && (*nodeList)[i].Row == current.Row) &&
				((*nodeList)[i].Height-current.Height <= 1) {
				current.Neighbours = append(current.Neighbours, &(*nodeList)[i])
				break
			}
		}
	}
}
