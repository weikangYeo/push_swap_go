package stack

import (
	"fmt"
	"push_swap/model"
)

func SwapStackA(head *list.Node) *list.Node {
	fmt.Println("sa")
	return swap(head)
}

func SwapStackB(head *list.Node) *list.Node {
	fmt.Println("sb")
	return swap(head)
}

func SwapBothStack(a, b *list.Node) (*list.Node, *list.Node) {
	fmt.Println("ss")
	a = swap(a)
	b = swap(b)
	return a, b
}

func PushToA (a, b *list.Node) (*list.Node, *list.Node) {
	
	fmt.Println("pa")
	return push(b, a)
}

func PushToB (a, b *list.Node) (*list.Node, *list.Node) {
	fmt.Println("pb")
	return push(a, b)
}

