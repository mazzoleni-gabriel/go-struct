package queue

type Queue struct {
	data []int
}

func (q *Queue) Push(data int) {
	q.data = append(q.data, data)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	data := q.data[0]
	q.data = q.data[1:]
	return data
}

func (q *Queue) Top() int {
	if q.IsEmpty() {
		panic("queue is empty")
	}
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
