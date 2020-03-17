package queue

type Queue []interface{}

// 参数 v int 即限定了 传入仅能传入 int
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// pop 时，pop 出去的 是一个 interface{} 类型的，不能说保证里面的 值就是 int,所以 类型断言下
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	// head 为 interface{} 类型的变量，将 肚子里面 int 类型的值 拿出来
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	// fmt.Println(len(*q))
	return len(*q) == 0
}
