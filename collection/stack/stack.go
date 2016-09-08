package stack


type Stack struct  {
	array []interface{}
	len int
}

func New() *Stack {
	//todo 初始化一个有初始容量的slice
	s := new(Stack)
	s.len = 0
	return s
}

//放入元素
func (this *Stack) Push(value ...interface{}) {
	this.array = append(this.array, value...)
	this.len ++
}

//返回下一个元素,并从Stack移除元素
func (this *Stack) Pop() interface{} {
	if this.len == 0 {
		return nil
	}
	this.len --
	value := this.array[this.len]
	this.array = append(this.array[:this.len])
	return value
}

//返回指定索引的元素
func (this *Stack) Get(i int) interface{} {
	return this.array[i]
}

//是否还有下一个元素
func (this *Stack) HasNext() bool {
	return this.len > 0
}

//func (this *Stack) String() string {
//	return strings.Join(this.array, ",")
//}