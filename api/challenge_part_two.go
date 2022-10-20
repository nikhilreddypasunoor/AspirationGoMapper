package api

import (
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

type ModString struct {
	inputStr string
	resStr   string
	idx      int
	i        int
}

func (str *ModString) TransformRune(pos int) {
	c := str.inputStr[pos]
	if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) {
		str.i++
		if str.i%str.idx == 0 {
			str.resStr += strings.ToUpper(string(c))
		} else {
			str.resStr += strings.ToLower(string(c))
		}
	} else {
		str.resStr += string(c)
	}
}

func (st ModString) GetValueAsRuneSlice() []rune {
	return []rune(st.inputStr)
}

func NewSkipString(idx int, inputStr string) ModString {
	return ModString{
		idx:      idx,
		inputStr: inputStr,
	}
}

func (s ModString) String() string {
	return s.resStr
}
