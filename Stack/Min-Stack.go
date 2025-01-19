// 155. Min Stack
// Basic stack implementation
// Just using a slice to track values on the stack
    // treating the 'back' of the slice as the 'front'
    // also using a slice to track the minimum value, doing the same thing
type MinStack struct {
    
    entries []int
    minimum []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.entries = append(this.entries, val)

    if len(this.minimum) == 0 || val <= this.minimum[len(this.minimum)-1] {
        this.minimum = append(this.minimum, val)
    }
}


func (this *MinStack) Pop()  {
    if len(this.entries) == 0 {
        return
    }

    if this.entries[len(this.entries)-1] == this.minimum[len(this.minimum)-1] {
        this.minimum = this.minimum[:len(this.minimum)-1]
    }
    this.entries = this.entries[:len(this.entries)-1]
}


func (this *MinStack) Top() int {
    return this.entries[len(this.entries) - 1]
}


func (this *MinStack) GetMin() int {
    return this.minimum[len(this.minimum) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 *
