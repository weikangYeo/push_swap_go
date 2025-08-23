package stack

import (
	"fmt"
	"push_swap/model"
)

func SwapStackA(head *model.Node) *model.Node {
	fmt.Println("sa")
	return swap(head)
}

func SwapStackB(head *model.Node) *model.Node {
	fmt.Println("sb")
	return swap(head)
}

func SwapBothStack(a, b *model.Node) (*model.Node, *model.Node) {
	fmt.Println("ss")
	a = swap(a)
	b = swap(b)
	return a, b
}

func PushToA(a, b *model.Node) (*model.Node, *model.Node) {
	fmt.Println("pa")
	b, a = push(b, a)
	return a, b
}

func PushToB(a, b *model.Node) (*model.Node, *model.Node) {
	fmt.Println("pb")
	return push(a, b)
}

func RotateA(head *model.Node) *model.Node {
	fmt.Println("ra")
	return rotate(head)
}

func RotateB(head *model.Node) *model.Node {
	fmt.Println("rb")
	return rotate(head)
}

func RotateBoth(a *model.Node, b *model.Node) (*model.Node, *model.Node) {
	fmt.Println("rr")
	a = rotate(a)
	b = rotate(b)
	return a, b
}

func ReverseRotateA(head *model.Node) *model.Node {
	fmt.Println("rra")
	return reverseRotate(head)
}

func ReverseRotateB(head *model.Node) *model.Node {
	fmt.Println("rrb")
	return reverseRotate(head)
}

func ReverseRotateBoth(a *model.Node, b *model.Node) (*model.Node, *model.Node) {
	fmt.Println("rrr")
	a = reverseRotate(a)
	b = reverseRotate(b)
	return a, b
}

// TODO the rest and cross check if follow Go idiom
