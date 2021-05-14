package ticktacktoe

import "fmt"

//go:generate stringer -type=Side
type Side byte

const (
	Empty Side = iota
	Circle
	Cross
)

const (
	width  = 3
	height = 3
)

type State struct {
	board    [width][height]Side
	nextTurn Side
}

func NewState(nextTurn Side, board [width][height]Side) State {
	return State{board: board, nextTurn: nextTurn}
}

func (s State) NextTurn() Side {
	return s.nextTurn
}

type WriteStateError struct {
	reason string
	cell   Side
	x, y   int
}

func (w WriteStateError) Error() string {
	return fmt.Sprintf("Cant write into cell %v Reason: %s", w.cell, w.reason)
}

func (cell Side) Flip() Side {
	switch cell {
	case Cross:
		return Circle
	case Circle:
		return Cross
	}
	panic(fmt.Sprintf("Cant Flip %+v", cell))
}

// Set doesnt check validity of state
func (s State) Set(x, y int, cell Side) State {
	s.board[x][y] = cell
	s.nextTurn = cell.Flip()
	return s
}

func (s State) CheckNWrite(x, y int, cell Side) (State, error) {
	if end, _ := s.WhoWon(); end {
		return s, WriteStateError{"Game already ended", cell, x, y}
	}
	existing := s.board[x][y]
	if existing != Empty {
		return s, WriteStateError{"Field is not empty", cell, x, y}
	}
	if cell != s.nextTurn {
		return s, WriteStateError{"Not your turn", cell, x, y}
	}
	s = s.Set(x, y, cell)
	return s, nil
}

func (s State) WhoWon() (bool, Side) {
	b := s.board
	places := [][3]Side{
		// columns
		{b[0][0], b[0][1], b[0][2]},
		{b[1][0], b[1][1], b[1][2]},
		{b[2][0], b[2][1], b[2][2]},
		// rows
		{b[0][0], b[1][0], b[2][0]},
		{b[0][1], b[1][1], b[2][1]},
		{b[0][2], b[1][2], b[2][2]},
		// diagonal
		{b[0][0], b[1][1], b[2][2]},
		{b[0][2], b[1][1], b[2][0]},
	}

	for _, place := range places {
		if place[0] == place[1] && place[1] == place[2] && place[0] != Empty {
			return true, place[0]
		}
	}
	for _, column := range s.board {
		for _, cell := range column {
			if cell == Empty {
				// there is no winner and there are still empty fields
				return false, Empty
			}
		}
	}
	// Tie
	return true, Empty
}
