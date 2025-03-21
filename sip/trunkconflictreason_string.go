// Code generated by "stringer -type TrunkConflictReason -trimprefix TrunkConflict"; DO NOT EDIT.

package sip

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TrunkConflictDefault-0]
	_ = x[TrunkConflictCalledNumber-1]
	_ = x[TrunkConflictCallingNumber-2]
}

const _TrunkConflictReason_name = "DefaultCalledNumberCallingNumber"

var _TrunkConflictReason_index = [...]uint8{0, 7, 19, 32}

func (i TrunkConflictReason) String() string {
	if i < 0 || i >= TrunkConflictReason(len(_TrunkConflictReason_index)-1) {
		return "TrunkConflictReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TrunkConflictReason_name[_TrunkConflictReason_index[i]:_TrunkConflictReason_index[i+1]]
}
