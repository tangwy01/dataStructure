package main

func main() {
	arry := ne

}

type Array struct {
	Arry    []int
	Len     int
	CruuLen int
}

func (a Array) New(len int) {
	return Array{make([]int, len), len}
}

func (a Array) Add(len int) {
	if a.Len == a.CruuLen {
		return
	}
	a.CruuLen++
	return make([]int, len)
}

func (a Array) RemoteFirst(len int) {
	if a.CruuLen == 0 {
		return
	}
	return make([]int, len)
}

func (a Array) RemoteAt(index int) {
	if index > a.CruuLen {
		return
	}
	return make([]int, len)
}
