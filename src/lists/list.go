/*
Package list implements a simple library for list structure.

author:heavy
version:1.0
*/

package lists

import (
   . "base"
   "fmt"
)

/* define structures */

// define a structure for element in linked list 
type Node struct {
    data Object
    next *Node
}

// define a structure for linked list
type List struct {
    size uint64
    head *Node
    tail *Node
    // define your match function
    myMatch MatchFun
}

/* define functions */

// data1 equals to data2 return 0
func defaultMatch(data1 Object, data2 Object) int {
    fmt.Println("defaultMatch")
    if data1 == data2 {
        return 0
    } else {
        return 1
    }
}

// choose suitable match
func (list *List) match(data1 Object, data2 Object) int {
    var matc MatchFun = nil
    if (*list).myMatch == nil {
        matc = defaultMatch
    } else {
        matc = (*list).myMatch
    }

    return matc(data1, data2)
}

func nextNode(node *Node) *Node {
    return (*node).next
}

func (list *List) getHead() *Node {
    return (*list).head
}

func (list *List) getTail() *Node {
    return (*list).tail
}

func (node *Node) getData() Object {
    return (*node).data
}

/* define interfaces */

func (list *List) Init(matchs ...MatchFun) {
    
    (*list).size = 0
    (*list).head = nil
    (*list).tail = nil

    if len(matchs) == 0 {
        (*list).myMatch = nil
    } else {
        (*list).myMatch = matchs[0]
    }
}

func (list *List) GetSize() uint64 {
    return (*list).size
}

func (list *List) Append(data Object) {
    newItem := new(Node)
    (*newItem).data = data
    (*newItem).next = nil

    if (*list).size == 0 {
        (*list).head = newItem
        (*list).tail = (*list).head
    } else {
        (*list).tail.next = newItem
        (*list).tail = newItem
    }

    (*list).size++
}

// TODO:insert

// get the first data
func (list *List) First() Object {
    if (list.GetSize() == 0) {
        return nil
    } else {
        return (*(list.getHead())).data
    }
}

// get the last data
func (list *List) Last() Object {
    if (list.GetSize() == 0) {
        return nil
    } else {
        return (*(list.getTail())).data
    }
}

// get the next element data from cur
func (list *List) Next(curData Object) Object {
    // get head
    head := list.getHead()
    // loop match data 
    for i := head; i != nil; i = nextNode(i) {
        if list.match(curData, i.getData()) == 0 {
            nxt := nextNode(i)
            if nxt == nil {
                return nil
            } else {
                return nxt.getData()
            }
        }
    }

    return nil
}