package robot_boat

//go:generate stringer -type=Position
type Position byte

const (
	Near Position = 0
	Far  Position = 1
)

type State struct {
	Robot, Fox, Hen, Grain Position
}

func validate(s State) bool {
	if s.Hen == s.Fox && s.Hen != s.Robot {
		return false
	}
	if s.Hen == s.Grain && s.Hen != s.Robot {
		return false
	}
	return true
}

func CreateValidStates() []State {
	var states []State
	for i := 0; i < 1<<4; i++ {
		s := State{Position(i & 1), Position(i & 2 >> 1), Position(i & 4 >> 2), Position(i & 8 >> 3)}
		if validate(s) {
			states = append(states, s)
		}
	}
	return states
}
