/*
Package list implements a simple library for list structure.

author:heavy
version:1.0
*/

package collections

import (
   . "base"
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

func (list *List) createNode(data Object) *Node {
    node := new(Node)
    (*node).data = data
    (*node).next = nil

    return node
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

func (list *List) insertAfterNode(node *Node, data Object) bool {
    //TODO:
    return true
}

// remove node at index
func (list *List) RemoveAt(index uint64) Object{
    size := list.GetSize()
    if index >= size { // out of rang
        return nil
    } else if size == 1 { // only one
        node := list.getHead()
        (*list).head = nil
        (*list).tail = nil
        (*list).size = 0
        return (*node).data
    } else if index == 0 { // remove head
        node := list.getHead()
        (*list).head = (*node).next
        (*list).size--

        return (*node).data
    } else if index == (size -1) { // remove tail
        preNode := list.getHead()
        for i := uint64(2); i < size; i++ {
            preNode = (*preNode).next
        }

        tail := list.getTail()
        (*list).tail = preNode
        preNode.next = nil
        (*list).size--

        return (*tail).data
    } else { // middle
        preNode := list.getHead()
        for i := uint64(2); i < index; i++ {
            preNode = (*preNode).next
        }

        node := (*preNode).next
        nxtNode := (*node).next
        (*node).next = nxtNode

        (*list).size--

        return (*node).data
    }
}

// remove data in the list
func (list *List) Remove(data Object) bool {
    if (data == nil || list.IsEmpty()) {
        return false
    }

    head := list.getHead()

    // remove head
    if (list.match(head.getData(), data) == 0) {
        (*list).head = nextNode(head)
    } else {
        // loop match data 
        cur := head
        nxt := nextNode(head)
        for ; nxt != nil; nxt = nextNode(nxt) {
            if list.match(data, nxt.getData()) == 0 {
                (*cur).next = nextNode(nxt)
                break
            }

            cur = nxt
        }

        if (nxt == nil) { // not in set
            return false
        }
    }

    (*list).size--

    return true
}

func (list *List) IsMember(data Object) bool {
    if (list.IsEmpty()) {
        return false
    }
     // get head
    
    head := list.getHead()
    // loop match data 
    for i := head; i != nil; i = nextNode(i) {
        if list.match(data, i.getData()) == 0 {
            return true
        }
    }

    return false
}

/* define interfaces */

func (list *List) Init(yourMatch ...MatchFun) {
    
    (*list).size = 0
    (*list).head = nil
    (*list).tail = nil

    if len(yourMatch) == 0 {
        (*list).myMatch = nil
    } else {
        (*list).myMatch = yourMatch[0]
    }

}

func (list *List) GetSize() uint64 {
    return (*list).size
}

func (list *List) IsEmpty() bool {
    return list.GetSize() == 0
}

func (list *List) Append(data Object) bool {
    newItem := new(Node)
    (*newItem).data = data
    (*newItem).next = nil

    if (*list).size == 0 {
        (*list).head = newItem
        (*list).tail = (*list).head
    } else {
        oldNode := (*list).tail
        (*oldNode).next = newItem
        (*list).tail = newItem
    }

    (*list).size++

    return true
}

func (list *List) InsertAtHead(data Object) bool {
    newNode := list.createNode(data)
    // insert head
    (*newNode).next = list.getHead()
    list.head = newNode
    (*list).size++

    return true
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

// get data at index,index start from 0
func (list *List) GetAt(index uint64) Object {
    size := list.GetSize()
    if index >= size {
        return nil
    } else if index ==0 {
        return list.First()
    } else if index == (size - 1) {
        return list.Last()
    } else {
        item := list.getHead()
        for i := uint64(0); i < size; i++ {
            if i == index {
                break
            }

            item = (*item).next
        }

        return item.getData()
    }
}

func (list *List) InsertAt(index uint64, data Object) bool {
    size := list.GetSize()
    if index > size { // out of index range
        return false
    } else if index == size { // add in list end
        return list.Append(data)
    } else if index == 0 { // insert at head
        return list.InsertAtHead(data)
    } else {
        newNode := list.createNode(data)
        prevIndex := index - 1
        prevItem := list.getHead()
        for i := uint64(0); i < size; i++ {
            if i == prevIndex {
                break
            }
            prevItem = (*prevItem).next
        }

        (*newNode).next = (*prevItem).next
        (*prevItem).next = newNode

        (*list).size++

        return true
    }
}

// remove all nodes
func (list *List) Clear() {
    list.Init()
}

