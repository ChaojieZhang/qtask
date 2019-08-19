package qtask

type queue struct {
	Count int
	head  *node
	tail  *node
}

func (queue *queue) Push(data interface{}) {
	node := &node{
		data: data,
	}
	tail := queue.tail
	if tail == nil {
		queue.tail = node
	} else {
		tail.next = node
		queue.tail = node
	}
	if queue.head == nil {
		queue.head = queue.tail
	}
	queue.Count++
}

func (queue *queue) Pop() interface{} {
	head := queue.head
	if head == nil {
		return nil
	}
	if head.next == nil {
		queue.head = nil
		queue.tail = nil
	} else {
		queue.head = head.next
	}
	queue.Count--
	return head.data
}

type node struct {
	data interface{}
	next *node
}
