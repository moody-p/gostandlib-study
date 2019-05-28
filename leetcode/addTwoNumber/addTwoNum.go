package main
import "fmt"
type ListNode struct {
	     Val int
	     Next *ListNode
}
func (l *ListNode)Len()int{
	var i int= 0
	for c := l; c != nil; c = c.Next{
		i += 1
	}
	return i
}
func (l *ListNode)Print(){
	for c := l; c != nil; c = c.Next{
		fmt.Println(c.Val)
	}
	fmt.Println()
}
func Init(nums []int) *ListNode{
	l := new(ListNode)
	c := l
	for _, v := range nums[:len(nums)-1]{
		c.Val = v
		c.Next = new(ListNode)
		c = c.Next
	}
	c.Val = nums[len(nums)-1]
	return l
}
func addTwoNum(l1 *ListNode, l2 *ListNode) *ListNode{
	newListNode := new(ListNode)
	c := newListNode
	c1,  c2 := l1, l2 
	for ; c1 != nil && c2 != nil; c1, c2 = c1.Next, c2.Next {
		c.Val += c1.Val + c2.Val
		c.Next = new(ListNode)
		if c.Val >= 10{
			c.Val = c.Val%10
			c.Next.Val = 1
		}else{
			c.Next.Val = 0
		}
		c = c.Next
	}
	if c1 != nil{
		c.Val += c1.Val
		c.Next = c1.Next
	}
	if c2 != nil{
		c.Val += c2.Val
		c.Next = c2.Next
	}
	if c.Val == 0 && c.Next == nil{
		fmt.Println(c.Next, c.Val)
		c = nil
		fmt.Println(c)
	}
	return newListNode
}
func main(){
	num1 := []int{2,4,3}
	num2 := []int{5, 6, 4}
	l1 := Init(num1)
	l2 := Init(num2)
	l1.Print()
	l2.Print()
	l3 := addTwoNum(l1, l2)
	l3.Print()
}