package linked

import (
	"errors"
	"fmt"
)

type LinkNode struct {
	Next *LinkNode
	Data interface{}
}

func (l *LinkNode) GetValue() interface{} {
	return l.Data
}

type LinkList struct {
	Head *LinkNode
	Len  uint
}

func NewLinkList() *LinkList {
	return &LinkList{&LinkNode{nil, nil}, 0}
}

func (this *LinkList) InsertAt(index uint, v interface{}) error {
	if err := this.outOfRange(index); err != nil {
		return err
	}
	if this.Len == 0 {
		l := LinkNode{nil, v}
		this.Head.Next = &l
		this.Len++
	} else {
		// important code
		cur := this.Head.Next
		for i := 0; i < int(index); i++ {
			cur = cur.Next
		}

		l := LinkNode{cur.Next.Next, v}
		cur.Next = &l
		this.Len++
	}
	return nil
}

func (this *LinkList) InsertInTail(v interface{}) {
	if this.Len == 0 {
		l := LinkNode{nil, v}
		this.Head.Next = &l
		this.Len++
	} else {
		// important code
		cur := this.Head.Next
		for i := 0; i < int(this.Len); i++ {
			cur = cur.Next
		}

		l := LinkNode{cur.Next.Next, v}
		cur.Next = &l
		this.Len++
	}
}

func (this *LinkList) RemoteInHead() {
	if this.Len == 0 {
		return
	} else {
		curr := this.Head.Next
		curr = curr.Next
	}
}

func (this *LinkList) RemoteAt(index uint) error {
	if err := this.outOfRange(index); err != nil {
		return err
	}
	if this.Len == 0 {
		return nil
	} else {
		// important code
		cur := this.Head.Next
		for i := 0; i < int(index)-1; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		this.Len--
	}
	return nil
}

func (this *LinkList) RemoteInTail() {
	this.RemoteAt(this.Len)
}

func (this *LinkList) outOfRange(index uint) error {
	if index > this.Len-1 || index < 0 {
		return errors.New("link out of range")
	}
	return nil
}

func (this *LinkList) GetValueAt(index uint) (interface{}, error) {
	if err := this.outOfRange(index); err != nil {
		return nil, err
	}
	curr := this.Head.Next
	for i := 0; i < int(index); i++ {
		curr = curr.Next
	}
	return curr.Data, nil
}

func (this *LinkList) Length() uint {
	return this.Len
}

//打印链表
func (this *LinkList) Print() {
	cur := this.Head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
