// Code generated by "stringer -type=Type"; DO NOT EDIT.

package user

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UNKNOWNT-0]
	_ = x[FREE-1]
	_ = x[PAID-2]
}

const _Type_name = "UNKNOWNTFREEPAID"

var _Type_index = [...]uint8{0, 8, 12, 16}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}