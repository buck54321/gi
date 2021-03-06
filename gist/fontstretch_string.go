// Code generated by "stringer -type=FontStretch"; DO NOT EDIT.

package gist

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FontStrNormal-0]
	_ = x[FontStrUltraCondensed-1]
	_ = x[FontStrExtraCondensed-2]
	_ = x[FontStrSemiCondensed-3]
	_ = x[FontStrSemiExpanded-4]
	_ = x[FontStrExtraExpanded-5]
	_ = x[FontStrUltraExpanded-6]
	_ = x[FontStrCondensed-7]
	_ = x[FontStrExpanded-8]
	_ = x[FontStrNarrower-9]
	_ = x[FontStrWider-10]
	_ = x[FontStretchN-11]
}

const _FontStretch_name = "FontStrNormalFontStrUltraCondensedFontStrExtraCondensedFontStrSemiCondensedFontStrSemiExpandedFontStrExtraExpandedFontStrUltraExpandedFontStrCondensedFontStrExpandedFontStrNarrowerFontStrWiderFontStretchN"

var _FontStretch_index = [...]uint8{0, 13, 34, 55, 75, 94, 114, 134, 150, 165, 180, 192, 204}

func (i FontStretch) String() string {
	if i < 0 || i >= FontStretch(len(_FontStretch_index)-1) {
		return "FontStretch(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FontStretch_name[_FontStretch_index[i]:_FontStretch_index[i+1]]
}

func (i *FontStretch) FromString(s string) error {
	for j := 0; j < len(_FontStretch_index)-1; j++ {
		if s == _FontStretch_name[_FontStretch_index[j]:_FontStretch_index[j+1]] {
			*i = FontStretch(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: FontStretch")
}
