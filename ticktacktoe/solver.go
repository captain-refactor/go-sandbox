package ticktacktoe

import "fmt"

//go:generate stringer -type=StateRank
type StateRank byte

const (
	Unknown StateRank = iota
	Defeat
	Draw
	Victory
)

type Solver struct {
	side       Side
	stateRanks map[State]StateRank
}

func (s *Solver) Side() Side {
	return s.side
}

func NewSolver(side Side) *Solver {
	return &Solver{side: side, stateRanks: map[State]StateRank{}}
}

func (s *Solver) GetRankForState(state State) StateRank {
	return s.stateRanks[state]
}

func (s *Solver) setRank(state State, rank StateRank) {
	s.stateRanks[state] = rank
}

func (s *Solver) RankState(state State) StateRank {
	rank := s.GetRankForState(state)
	if rank != Unknown {
		return rank
	}

	end, side := state.WhoWon()
	if end {
		switch side {
		case s.side:
			s.setRank(state, Victory)
			return Victory
		case s.side.Flip():
			s.setRank(state, Defeat)
			return Defeat
		case Empty:
			s.setRank(state, Draw)
			return Draw
		}
	}

	possibilities := PredictPossibleStates(state)

	if state.nextTurn == s.side {
		rank = Defeat
	} else {
		rank = Victory
	}
	for _, possibility := range possibilities {
		childRank := s.RankState(possibility)
		if s.side == state.nextTurn {
			if childRank > rank {
				rank = childRank
			}
		} else {
			if childRank < rank {
				rank = childRank
			}
		}
	}
	return rank
}

func (s *Solver) NextMove(state State) (x, y int) {
	s.RankState(state)
	var bestFuture State
	bestRank := Unknown
	fmt.Print("My choices: ")
	for _, possibility := range PredictPossibleStates(state) {
		rank := s.RankState(possibility)
		fmt.Print(rank, " ")
		if rank > bestRank {
			bestFuture = possibility
			bestRank = rank
		}
	}
	fmt.Println()
	return stateDelta(state, bestFuture)
}

func stateDelta(a, b State) (x, y int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if a.board[x][y] != b.board[x][y] {
				return x, y
			}
		}
	}
	return -1, -1
}
