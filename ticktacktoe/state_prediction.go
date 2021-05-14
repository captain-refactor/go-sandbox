package ticktacktoe

func PredictPossibleStates(state State) (possibilities []State) {
	for x, column := range state.board {
		for y, cell := range column {
			if cell == Empty {
				write, err := state.CheckNWrite(x, y, state.nextTurn)
				if err != nil {
					panic(err)
				}
				possibilities = append(possibilities, write)
			}
		}
	}
	return possibilities
}
