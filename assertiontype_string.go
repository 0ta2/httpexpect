// Code generated by "stringer -type=AssertionType"; DO NOT EDIT.

package httpexpect

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AssertUsage-0]
	_ = x[AssertOperation-1]
	_ = x[AssertValid-2]
	_ = x[AssertNotValid-3]
	_ = x[AssertNil-4]
	_ = x[AssertNotNil-5]
	_ = x[AssertEmpty-6]
	_ = x[AssertNotEmpty-7]
	_ = x[AssertEqual-8]
	_ = x[AssertNotEqual-9]
	_ = x[AssertLt-10]
	_ = x[AssertLe-11]
	_ = x[AssertGt-12]
	_ = x[AssertGe-13]
	_ = x[AssertInRange-14]
	_ = x[AssertNotInRange-15]
	_ = x[AssertMatchSchema-16]
	_ = x[AssertNotMatchSchema-17]
	_ = x[AssertMatchPath-18]
	_ = x[AssertNotMatchPath-19]
	_ = x[AssertMatchRegexp-20]
	_ = x[AssertNotMatchRegexp-21]
	_ = x[AssertContainsKey-22]
	_ = x[AssertNotContainsKey-23]
	_ = x[AssertContainsElement-24]
	_ = x[AssertNotContainsElement-25]
	_ = x[AssertContainsSubset-26]
	_ = x[AssertNotContainsSubset-27]
	_ = x[AssertBelongs-28]
	_ = x[AssertNotBelongs-29]
}

const _AssertionType_name = "AssertUsageAssertOperationAssertValidAssertNotValidAssertNilAssertNotNilAssertEmptyAssertNotEmptyAssertEqualAssertNotEqualAssertLtAssertLeAssertGtAssertGeAssertInRangeAssertNotInRangeAssertMatchSchemaAssertNotMatchSchemaAssertMatchPathAssertNotMatchPathAssertMatchRegexpAssertNotMatchRegexpAssertContainsKeyAssertNotContainsKeyAssertContainsElementAssertNotContainsElementAssertContainsSubsetAssertNotContainsSubsetAssertBelongsAssertNotBelongs"

var _AssertionType_index = [...]uint16{0, 11, 26, 37, 51, 60, 72, 83, 97, 108, 122, 130, 138, 146, 154, 167, 183, 200, 220, 235, 253, 270, 290, 307, 327, 348, 372, 392, 415, 428, 444}

func (i AssertionType) String() string {
	if i >= AssertionType(len(_AssertionType_index)-1) {
		return "AssertionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AssertionType_name[_AssertionType_index[i]:_AssertionType_index[i+1]]
}
