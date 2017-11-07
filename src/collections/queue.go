/*
Package list implements a simple library for list structure.
FIFO
author:heavy
version:1.0
*/
package collections

import (
	. "base"
)

type Queue struct {
    list *List
}


func (queue *Queue) Init() {
    lst := new(List)
    (*queue).list = lst
    
    lst.Init()
}

func (queue *Queue) Enqueue(data Object) bool {
    return (*queue).list.Append(data)
}

func (queue *Queue) Dequeue() Object {
    return (*queue).list.RemoveAt(0)
}

func (queue *Queue) GetSize() uint64 {
    return (*queue).list.GetSize()
}

func (queue *Queue) Peek() Object {
    return (*queue).list.First()
}