package stack

type Stack struct {
	data []any
}

func (q *Stack) Push(data any) {
	q.data = append(q.data, data)
}

func (q *Stack) Pop() any {
	if q.IsEmpty() {
		panic("stack is empty")
	}
	length := len(q.data)
	data := q.data[length-1]
	q.data = q.data[:length-1]
	return data
}

func (q *Stack) Top() any {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.data[len(q.data)-1]
}

func (q *Stack) IsEmpty() bool {
	return len(q.data) == 0
}
