package queue

// Queue 使用别名扩充已有类型
//type Queue []int
type Queue []interface{} // 支持任何类型

// Push 因为值要发生改变，所以要用指针接收者，所以每次Push object对象都会发生改变，不再是原来的那个object了
func (q *Queue) Push(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
