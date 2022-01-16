package min_stack

import (
	"fmt"
)


func Example_maxStack() {
	stk := Constructor()

	stk.Push(-2)
	stk.Push(0)
	stk.Push(-3)
	fmt.Println(stk.GetMin())
	stk.Pop()
	fmt.Println(stk.Top())
	fmt.Println(stk.GetMin())

	// Output:
	// -3
	// 0
	// -2
}
