// Code generated by "stringer -type=StateRank"; DO NOT EDIT.

package ticktacktoe

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[Defeat-1]
	_ = x[Draw-2]
	_ = x[Victory-3]
}

const _StateRank_name = "UnknownDefeatDrawVictory"

var _StateRank_index = [...]uint8{0, 7, 13, 17, 24}

func (i StateRank) String() string {
	if i >= StateRank(len(_StateRank_index)-1) {
		return "StateRank(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StateRank_name[_StateRank_index[i]:_StateRank_index[i+1]]
}
