package leetcode

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{
		values: []int{},
		mins:   []int{},
	}
}

func (this *MinStack) Push(value int) {
	this.values = append(this.values, value)
	min := value
	if len(this.mins) > 0 && this.mins[len(this.mins)-1] < min {
		min = this.mins[len(this.mins)-1]
	}
	this.mins = append(this.mins, min)

}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
