type MyQueue struct {
    StackIn []int
    StackOut []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.StackIn = append(this.StackIn,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.StackOut) > 0{
        x:=this.StackOut[len(this.StackOut)-1]
        this.StackOut = this.StackOut[:len(this.StackOut)-1]
        return x
    }
    for i:=len(this.StackIn)-1;i>=0;i--{
        this.StackOut = append(this.StackOut,this.StackIn[i])  
    }
    this.StackIn =this.StackIn[0:0]
    x:=this.StackOut[len(this.StackOut)-1]
    this.StackOut = this.StackOut[:len(this.StackOut)-1]
    return x 
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.StackOut) > 0{
        x:=this.StackOut[len(this.StackOut)-1]
        return x
    }
    for i:=len(this.StackIn)-1;i>=0;i--{
        this.StackOut = append(this.StackOut,this.StackIn[i])  
    }
    this.StackIn =this.StackIn[0:0]
    return this.StackOut[len(this.StackOut)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.StackOut) == 0 && len(this.StackIn) == 0 {
        return true
    }
    return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */