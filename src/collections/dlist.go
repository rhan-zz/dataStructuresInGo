/*
Package list implements a simple library for list structure.
double list
author:heavy
version:1.0
*/

package collections

import (
	. "base"
)

/* define structures  */

// define a node for double linked list
type DNode struct {
    data Object
    prev *DNode
    next *DNode
}

// define the double linked list
type DList struct {
    size uint64
    head *DNode
    tail *DNode
}

/* define functions */



/* define methods */



func (dList *DList) Init() {
    _dList := *(dList)
    _dList.size = 0
    _dList.head = nil
    _dList.tail = nil
}

func (dList *DList) GetSize() uint64 {
    return (*dList).size
}

func (dList *DList) GetHead() *DNode {
    return (*dList).head
}

func (dList *DList) GetTail() *DNode {
    return (*dList).tail
}

func (dList *DList) isHead(elmt *DNode) bool {
    return dList.GetHead() == elmt
}

func (dList *DList) isTail(elmt *DNode) bool {
    return dList.GetTail() == elmt
}

func (dList *DList) Append(data Object) {
    newNode := new(DNode)
    (*newNode).data =  data

    if (*dList).GetSize() == 0 { // empty
        (*dList).head = newNode
        (*dList).tail = newNode
        (*newNode).prev = nil
        (*newNode).next = nil
    } else { // tail
        (*newNode).prev = (*dList).tail
        (*newNode).next = nil
        (*((*dList).tail)).next = newNode
        (*dList).tail = newNode
    }

    (*dList).size++;
}

func (dList *DList) InsertNext(elmt *DNode, data Object) bool {
    if elmt == nil { // apend
        return false   
    }

    if dList.isTail(elmt) {
        dList.Append(data)
    } else {
        newNode := new(DNode)
        (*newNode).data =  data
        (*newNode).prev = elmt
        (*newNode).next = (*elmt).next

        (*elmt).next = newNode
        (*((*newNode).next)).prev = newNode
        (*dList).size++;
    }

    return true
}

func (dList *DList) InsertPrev(elmt *DNode, data Object) bool {
    if elmt == nil {
        return false
    }

    if dList.isHead(elmt) { 
        newNode := new(DNode)
        (*newNode).data = data
        (*newNode).next = dList.GetHead()
        (*newNode).prev = nil

        (*(dList.head)).prev = newNode
        dList.head = newNode
        dList.size++
        return true
    } else {
        prev := (*elmt).prev
        return dList.InsertNext(prev, data)
    }
}


// remove the element in double list
func (dList *DList) Remove(elmt *DNode) Object {
    if elmt == nil {
        return false
    }

    prev := (*elmt).prev
    next := (*elmt).next

    if dList.isHead(elmt) {
        dList.head = next
    } else {
        (*prev).next = next
    }

    if dList.isTail(elmt) {
        dList.tail = prev
    } else {
        (*next).prev = prev
    }

    dList.size--

    return (*elmt).GetData() 
}

// find the node in double list that the data match
func (dList *DList) Search(data Object, yourMatch ...MatchFun) *DNode {
    if dList.GetSize() == 0 {
        return nil
    }

    match := defaultMatch
    if len(yourMatch) > 0 {
        match = yourMatch[0]
    }

    node := dList.GetHead()
    for ; node != nil; node = node.GetNext() {
        if match(node.GetData(), data) == 0 {
            break
        }
    }

    return node
}



func (dNode *DNode) GetData() Object {
    return (*dNode).data
}

func (dNode *DNode) GetNext() *DNode {
    return (*dNode).next
}

func (dNode *DNode) GetPrev() *DNode {
    return (*dNode).prev
}

