package max_stack

import (
	"fmt"
)

func Example_maxStack() {
	stk := Constructor()

	stk.Push(74)
	fmt.Println(stk.PopMax())
	stk.Push(89)
	stk.Push(67)
	fmt.Println(stk.PopMax())
	fmt.Println(stk.Pop())
	stk.Push(61)
	stk.Push(-77)
	fmt.Println(stk.PeekMax())
	fmt.Println(stk.PopMax())
	stk.Push(81)
	fmt.Println(stk.Pop())
	stk.Push(-71)
	stk.Push(32)

	// Output:
	// 74
	// 89
	// 67
	// 61
	// 61
	// 81
}

func Example_maxStack1() {
	stk := Constructor()

	stk.Push(5)
	stk.Push(1)
	stk.Push(5)
	fmt.Println(stk.Top())
	fmt.Println(stk.PopMax())
	fmt.Println(stk.Top())
	fmt.Println(stk.PeekMax())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Top())

	// Output:
	// 5
	// 5
	// 1
	// 5
	// 1
	// 5
}

func Example_maxStack2() {
	stk := Constructor()

	stk.Push(-45)
	stk.Push(97)
	stk.Push(-12)
	fmt.Println("Top", stk.Top())
	fmt.Println("PopMax", stk.PopMax())
	fmt.Println("Top", stk.Top())
	fmt.Println("PeekMax", stk.PeekMax())
	fmt.Println("PeekMax", stk.PeekMax())
	fmt.Println("Pop", stk.Pop())
	fmt.Println("PeekMax", stk.PeekMax())
	fmt.Println("Pop", stk.Pop())
	stk.Push(91)

	// Output:
	// Top -12
	// PopMax 97
	// Top -12
	// PeekMax -12
	// PeekMax -12
	// Pop -12
	// PeekMax -45
	// Pop -45
}



