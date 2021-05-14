package ticktacktoe

import (
	"reflect"
	"testing"
)

func TestPredictPossibleStates(t *testing.T) {
	var startState = State{nextTurn: Circle}
	type args struct {
		state State
	}

	midGame := startState.Set(2, 0, Circle).Set(2, 1, Cross).Set(1, 1, Circle)

	tests := []struct {
		name              string
		args              args
		wantPossibilities []State
	}{
		{
			name: "Start",
			args: struct{ state State }{state: startState},
			wantPossibilities: []State{
				startState.Set(0, 0, Circle),
				startState.Set(0, 1, Circle),
				startState.Set(0, 2, Circle),
				startState.Set(1, 0, Circle),
				startState.Set(1, 1, Circle),
				startState.Set(1, 2, Circle),
				startState.Set(2, 0, Circle),
				startState.Set(2, 1, Circle),
				startState.Set(2, 2, Circle),
			},
		},
		{
			name: "Mid Game",
			args: struct{ state State }{state: midGame},
			wantPossibilities: []State{
				midGame.Set(0, 0, Cross),
				midGame.Set(0, 1, Cross),
				midGame.Set(0, 2, Cross),
				midGame.Set(1, 0, Cross),
				midGame.Set(1, 2, Cross),
				midGame.Set(2, 2, Cross),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPossibilities := PredictPossibleStates(tt.args.state); !reflect.DeepEqual(gotPossibilities, tt.wantPossibilities) {
				t.Errorf("PredictPossibleStates() = %v, want %v", gotPossibilities, tt.wantPossibilities)
			}
		})
	}
}
