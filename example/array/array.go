package array

import "fmt"

type Array struct {
	Arry []int
	Len  int
}

func NewArray(len uint) *Array {
	return &Array{(make([]int, len, len)), 0}
}

func (a *Array) AppendInTail(value int) {
	a.Arry[a.Len] = value
	a.Len++
}

func (a *Array) Length() int {
	return a.Len
}

func (a *Array) RemoveAt(index int) {
	if !a.isOutofRange(index) {
		for index < a.Len-1 {
			a.Arry[index] = a.Arry[index+1]
			index++
		}
		a.Len--
	}
}

func (a *Array) AppendAt(index, value int) {
	if !a.isOutofRange(index) {
		for i := a.Len; i > index; i-- {
			a.Arry[i] = a.Arry[i-1]
		}
		a.Arry[index] = value
		a.Len++
	}
}

func (a *Array) isOutofRange(index int) bool {
	if index < 0 || index > cap(a.Arry) {
		return true
	}

	return false
}

//打印数列
func (a *Array) Print() {
	var format string
	for i := 0; i < a.Length(); i++ {
		format += fmt.Sprintf("|%+v", a.Arry[i])
	}
	fmt.Println(format)
}
