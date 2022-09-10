package main

func main() {

}

type CQueue struct {
	store []int
	help  []int
}

//func Constructor() CQueue {
//	return CQueue{
//		store: []int{},
//		help:  []int{},
//	}
//}

func (this *CQueue) AppendTail(value int) {
	this.store = append(this.store, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.help) != 0 {
		res := this.help[len(this.help)-1]
		this.help = this.help[:len(this.help)-1]
		return res
	}

	if len(this.store) == 0 {
		return -1
	}

	for i := len(this.store) - 1; i >= 0; i-- {
		this.help = append(this.help, this.store[i])
	}

	this.store = []int{}
	res := this.help[len(this.help)-1]
	this.help = this.help[:len(this.help)-1]
	return res
}

type MinStack struct {
}

/** initialize your data structure here. */
func Constructor() MinStack {
	
}

func (this *MinStack) Push(x int) {

}

func (this *MinStack) Pop() {

}

func (this *MinStack) Top() int {

}

func (this *MinStack) Min() int {

}
