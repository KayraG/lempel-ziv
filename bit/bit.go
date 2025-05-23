package app

import (
	"fmt"
	"strings"
)

type Bit = int

func AsString(b Bit) string {
	return fmt.Sprintf("%d", b)
}

type BitList []Bit

func NewBitList() BitList {
	return BitList{}
}

func (s BitList) ConvertToString() string {
	var b strings.Builder
	for _, value := range s {
		b.WriteString(AsString(value))
	}

	return b.String()
}

func (s1 BitList) EqualTo(s2 BitList) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

type List struct {
	_ctnt []BitList
}

func NewEmptyList() List {
	return List{[]BitList{}}
}

func Add(s List, e BitList) List {
	return List{append(s._ctnt, e)}
}

func (s List) Has(e BitList) bool {
	l := len(s._ctnt)
	for i := 0; i < l; i++ {
		if e.EqualTo(s._ctnt[i]) {
			return true
		}
	}
	return false
}

func (s List) Get(i int) BitList {
	l := len(s._ctnt)
	if i >= l {
		return nil
	}
	return s._ctnt[l]
}

func (s List) IndexOf(e BitList) int {
	l := len(s._ctnt)
	for i := 0; i < l; i++ {
		if e.EqualTo(s._ctnt[i]) {
			return i
		}
	}
	return -1
}
