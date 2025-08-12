package stack

import "push_swap/model"

// swap top 2 node
func swap(head *list.Node) *list.Node {
    if head == nil || head.NextNode == nil {
        return head
    }

    newHead := head.NextNode
    head.NextNode = newHead.NextNode
    newHead.NextNode = head
    return newHead
}

// push moves the top node from src to dest and returns the new heads of both lists.
func push(srcHead *list.Node, destHead *list.Node) (*list.Node, *list.Node) {
    // If the source list is empty, there's nothing to move.
    if srcHead == nil {
        return srcHead, destHead
    }

    newDestHead := srcHead
    newSrcHead := srcHead.NextNode
    newDestHead.NextNode = destHead

    return newSrcHead, newDestHead
}

// rotate the head to tail node
func rotate(head *list.Node) *list.Node {
    if head == nil || head.NextNode == nil {
        return head
    }

    // iterate to end of tail
    tail := head
    for tail.NextNode != nil {
        tail = tail.NextNode
    }

    newHead := head.NextNode
    tail.NextNode = head
    head.NextNode = nil
    return newHead
}

// change tail node to head node
func reverseRotate(head *list.Node) *list.Node {
    if head == nil || head.NextNode == nil {
        return head
    }

    // iterate to end of tail
    tail := head
    secondLast := head
    for tail.NextNode != nil {
        secondLast = tail
        tail = tail.NextNode
    }

    tail.NextNode = head
    secondLast.NextNode = nil

    return tail
}

func CreateStack(intSlice []int) *list.Node {
    var head *list.Node

    for i := len(intSlice) - 1; i >= 0; i-- {
        // first iteration, head would be nil
        // head will always reassigned as new node address at each iteration,
        // and new node.NextNode would be previous head address
        head = &list.Node{intSlice[i], head}
    }

    return head
}
