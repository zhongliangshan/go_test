package queue

type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
