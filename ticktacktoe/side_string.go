// Code generated by "stringer -type=Side"; DO NOT EDIT.

package ticktacktoe

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Empty-0]
	_ = x[Circle-1]
	_ = x[Cross-2]
}

const _Side_name = "EmptyCircleCross"

var _Side_index = [...]uint8{0, 5, 11, 16}

func (i Side) String() string {
	if i >= Side(len(_Side_index)-1) {
		return "Side(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Side_name[_Side_index[i]:_Side_index[i+1]]
}