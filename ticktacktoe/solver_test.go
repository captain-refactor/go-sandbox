package ticktacktoe

import (
	"testing"
)

func TestNewSolver(t *testing.T) {
	NewSolver(Circle)
}

func TestSolver_RankState(t *testing.T) {
	var startState = State{nextTurn: Circle}
	type fields struct {
		side       Side
		stateRanks map[State]StateRank
	}
	type args struct {
		state State
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   StateRank
	}{
		{
			name: "Certain deatch",
			args: struct{ state State }{
				state: startState.Set(0, 0, Circle).Set(0, 1, Cross).Set(1, 1, Circle).Set(0, 2, Cross),
			},
			want: Defeat,
		},
		{
			name: "Possible win",
			args: struct{ state State }{
				state: NewState(Cross, [3][3]Side{
					{Cross, Empty, Cross},
					{Circle, Empty, Cross},
					{Circle, Empty, Circle},
				}),
			},
			want: Victory,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solver{
				side:       Cross,
				stateRanks: map[State]StateRank{},
			}
			if got := s.RankState(tt.args.state); got != tt.want {
				t.Errorf("RankState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolver_NextMove(t *testing.T) {
	state := State{
		board: [3][3]Side{
			{Circle, Cross, Empty},
			{Cross, Empty, Circle},
			{Circle, Cross, Empty},
		},
		nextTurn: Circle,
	}
	solver := NewSolver(Circle)
	x, y := solver.NextMove(state)
	if x != 1 || y != 1 {
		t.Fail()
	}
	state, err := state.CheckNWrite(x, y, Circle)
	if err != nil {
		panic(err)
	}
	state, err = state.CheckNWrite(0, 2, Cross)
	if err != nil {
		panic(err)
	}
	x, y = solver.NextMove(state)
	if x != 2 || y != 2 {
		t.Fail()
	}
	state, err = state.CheckNWrite(x, y, Circle)
	if err != nil {
		panic(err)
	}
	_, side := state.WhoWon()
	if side != Circle {
		t.Fail()
	}
}
