/*
Package list implements a simple library for list structure.
FILO
author:heavy
version:1.0
*/

package collections

import (
    . "base"
)

type Stack struct {
    list *List
}

func (stack *Stack) Init() {
    lst := new(List)
    (*stack).list = lst
    lst.Init()
}

func (stack *Stack) Push(data Object) bool {
    lst := (*stack).list

    return lst.InsertAtHead(data)
}

func (stack *Stack) Pop() Object {
    lst := (*stack).list

    return lst.RemoveAt(0)
}

func (stack *Stack) Peek() Object {
    lst := (*stack).list

    return lst.First()
}

func (stack *Stack) GetSize() uint64 {
    lst := (*stack).list

    return lst.GetSize()
}
