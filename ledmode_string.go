// Code generated by "stringer -type=ledMode"; DO NOT EDIT

package xethru

import "fmt"

const _ledMode_name = "LEDOffLEDSimpleLEDFull"

var _ledMode_index = [...]uint8{0, 6, 15, 22}

func (i ledMode) String() string {
	if i >= ledMode(len(_ledMode_index)-1) {
		return fmt.Sprintf("ledMode(%d)", i)
	}
	return _ledMode_name[_ledMode_index[i]:_ledMode_index[i+1]]
}
