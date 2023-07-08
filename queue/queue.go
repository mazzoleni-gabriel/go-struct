package queue

type Queue struct {
	data []any
}

func (q *Queue) Push(data any) {
	q.data = append(q.data, data)
}

func (q *Queue) Pop() any {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	data := q.data[0]
	q.data = q.data[1:]
	return data
}

func (q *Queue) Top() any {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
