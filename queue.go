package qtask

type queue struct {
	Count int
	head  *node
	tail  *node
}

func (q *queue) Push(task Task) {
	node := &node{
		task: task,
	}
	tail := q.tail
	if tail == nil {
		q.tail = node
	} else {
		tail.next = node
		q.tail = node
	}
	if q.head == nil {
		q.head = q.tail
	}
	q.Count++
}

func (q *queue) Pop() Task {
	head := q.head
	if head == nil {
		return nil
	}
	if head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = head.next
	}
	q.Count--
	return head.task
}

func (q *queue) Remove(task Task) {
	// fake remove
	q.Pop()
}

type node struct {
	task Task
	next *node
}
