// Code generated by "stringer -type DispatchRuleConflictReason -trimprefix DispatchRuleConflict"; DO NOT EDIT.

package sip

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DispatchRuleConflictGeneric-0]
}

const _DispatchRuleConflictReason_name = "Generic"

var _DispatchRuleConflictReason_index = [...]uint8{0, 7}

func (i DispatchRuleConflictReason) String() string {
	if i < 0 || i >= DispatchRuleConflictReason(len(_DispatchRuleConflictReason_index)-1) {
		return "DispatchRuleConflictReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DispatchRuleConflictReason_name[_DispatchRuleConflictReason_index[i]:_DispatchRuleConflictReason_index[i+1]]
}
