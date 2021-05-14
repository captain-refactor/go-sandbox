package robot_boat

import (
	"bytes"
)

type DeadEnds [][]byte

func (m DeadEnds) Add(path []State) DeadEnds {
	serialized := Serialize(path)
	for _, p := range m {
		if bytes.Equal(p, serialized) {
			return m
		}
	}
	return append(m, serialized)
}

var desiredState = State{Far, Far, Far, Far}

func (m DeadEnds) Includes(path []State) bool {
	serialized := Serialize(path)
	for _, p := range m {
		if bytes.Equal(p, serialized) {
			return true
		}
	}
	return false
}

func Serialize(path []State) []byte {
	var result []byte
	for _, state := range path {
		result = append(result, byte(state.Robot), byte(state.Fox), byte(state.Hen), byte(state.Grain))
	}
	return result
}

func includes(path []State, state State) bool {
	for _, s := range path {
		if s == state {
			return true
		}
	}
	return false
}

func getNextStep(path []State, possibilities []State, deadEnds DeadEnds) (found bool, state State) {
	for _, possibility := range possibilities {
		isDeadEnd := deadEnds.Includes(append(path, possibility))
		isCircling := includes(path, possibility)
		if isDeadEnd || isCircling {
			continue
		}
		found = true
		state = possibility
		if possibility == desiredState {
			return
		}
	}
	return
}

func SolvePuzzle() ([][]State, int) {
	initialState := State{Near, Near, Near, Near}
	states := CreateValidStates()
	transitions := CreateValidTransitions(states)
	var solutions [][]State
	var deadEnds DeadEnds
	path := []State{initialState}
	iterations := 0
	for {
		iterations++

		if len(path) == 0 {
			//there is no possible result
			return solutions, iterations
		}
		last := path[len(path)-1]
		possibilities := transitions.GetTransitions(last)
		found, next := getNextStep(path, possibilities, deadEnds)
		if !found {
			deadEnds = deadEnds.Add(path)
			path = path[:len(path)-1]
			continue
		}
		path = append(path, next)
		if next == desiredState {
			solutions = append(solutions, path)
			deadEnds = deadEnds.Add(path)
		}
	}
}
