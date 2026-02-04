package solution
type MyQueue struct {
    buffer []int
    head int
    tail int
    over bool   
}


func Constructor() MyQueue {
 return MyQueue{
     buffer: make([]int, 101),
 }
}

func (this *MyQueue) Push(x int)  {
    if this.tail == this.head && this.over {
        if this.head++;this.head == len(this.buffer) {
            this.head = 0
            this.over = false
        }
    }
    this.buffer[this.tail] = x
    this.tail++
    if this.tail == len(this.buffer) {
        this.tail = 0
        this.over = true
    }
}


func (this *MyQueue) Pop() int {
    if this.head != this.tail || this.over {
        val := this.buffer[this.head]
        this.head++
        if this.head == len(this.buffer) {
            this.head = 0
            this.over = false
        }
        return val
    }
    return -1
}


func (this *MyQueue) Peek() int {
    return this.buffer[this.head]
}


func (this *MyQueue) Empty() bool {
    if this.over {
        return 1 * len(this.buffer) + this.tail - this.head == 0
    } else {
        return 0 * len(this.buffer) + this.tail - this.head == 0
    }
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
