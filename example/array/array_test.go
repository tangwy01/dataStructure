package array

import (
	"testing"
)

func TestArray_AppendInTail(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		arr.AppendAt(i, i+1)
	}
	arr.Print()

	arr.AppendAt(6, 999)
	arr.Print()

	arr.AppendInTail(666)
	arr.AppendAt(60, 999)
	arr.Print()
}
