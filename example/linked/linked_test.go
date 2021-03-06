package linked

import "testing"

func TestInsertToHead(t *testing.T) {
	l := NewLinkList()
	for i := 0; i < 10; i++ {
		l.InsertAt(uint(i), i+1)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkList()
	for i := 0; i < 10; i++ {
		l.InsertInTail(i + 1)
	}
	l.Print()
}

// func TestFindByIndex(t *testing.T) {
// 	l := NewLinkList()
// 	for i := 0; i < 10; i++ {
// 		l.InsertInTail(i + 1)
// 	}
// 	t.Log(l.FindByIndex(0))
// 	t.Log(l.FindByIndex(9))
// 	t.Log(l.FindByIndex(5))
// 	t.Log(l.FindByIndex(11))
// }

func TestDeleteNode(t *testing.T) {
	l := NewLinkList()
	for i := 0; i < 3; i++ {
		l.InsertInTail(i + 1)
	}
	l.Print()

	// t.Log(l.DeleteNode(l.head.next))
	// l.Print()

	// t.Log(l.DeleteNode(l.head.next.next))
	// l.Print()
}
