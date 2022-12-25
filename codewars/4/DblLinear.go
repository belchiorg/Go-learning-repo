type Node struct {
	next *Node
	n50  *Node
	key  int
}

type List struct {
	head  *Node
	count int
}

func DblLinear(n int) int {

	LinkedList := List{count: 0}

	LinkedList.insertNum(1)

	node := LinkedList.head

	for i := 0; i < n; i++ {
		LinkedList.insertNum(node.key*2 + 1)
		LinkedList.insertNum(node.key*3 + 1)
		node = node.next
	}

	node = LinkedList.head
	m := n
	for m > 1000 {
		node = node.n50
		m -= 1000
	}
	for i := 0; i < m; i++ {
		node = node.next
	}
	return node.key
}

func (L *List) insertNum(n int) {
	list := &Node{
		next: nil,
		key:  n,
		n50:  nil,
	}

	if L.head == nil {
		L.head = list
		L.count++
		return
	}

	i := 0
	node50 := L.head
	l := L.head
	for l.n50 != nil && l.n50.key < n {
		l = l.n50
		node50 = l
	}
	for l.next != nil && l.next.key < n {
		l = l.next
		i++
		if i == 1000 {
			i = 0
			node50.n50 = l
			node50 = l
		}
	}
	if l.next != nil && l.next.key == n {
		return
	}
	L.count++
	list.next = l.next
	l.next = list
}