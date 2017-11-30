package main

import (
    "fmt"
    . "collections"
    . "base"
)

type MyStr struct {
    name string
}

func match(data1 Object, data2 Object) int {
    myStr1 := data1.(*MyStr)
    myStr2 := data2.(*MyStr)

    if (*myStr1).name == (*myStr2).name {
        return 0
    } else {
        return 1
    }

} 


func main(){
    /*
    var lst List
    var plst *List = &lst
    plst.Init(match)
    
    myStr1 := new(MyStr)
    myStr2 := new(MyStr)
    myStr3 := new(MyStr)
    myStr4 := new(MyStr)

    myStrTag := new(MyStr)

    (*myStr1).name = "1"
    (*myStr2).name = "2"
    (*myStr3).name = "3"
    (*myStr4).name = "4"

    (*myStrTag).name = "2"    

    plst.Append(myStr1)
    plst.Append(myStr2)
    plst.Append(myStr3)
    plst.Append(myStr4)

    fmt.Println(plst.GetSize())

    txt1 := lst.Next(myStrTag)
    txt2,r := txt1.(*MyStr)
    if (!r) {
        fmt.Println(r)
    } else {
        fmt.Println((*txt2).name)
    }

    lst := new (Queue)
    lst.Init()*/
    /*
    a := 1
    b := 2
    c := 3
    */

    //var d int
    /*lst.Enqueue("a")
    lst.Enqueue("c")
    lst.Enqueue("d")
    

    for ; lst.GetSize() > 0;  {
        fmt.Println(lst.Dequeue().(string))
    }
*/
   
/*
    node := lst.GetHead().GetNext()
    fmt.Println(lst.GetSize())

    d = lst.Remove(node).(int)

    fmt.Println(lst.GetSize())
    fmt.Println(d)
    */

    set := new(Set)
    set.Init(match)

    set1 := new(Set)
    set1.Init(match)

    myStr1 := new(MyStr)
    myStr2 := new(MyStr)
    myStr3 := new(MyStr)
    myStr4 := new(MyStr)

    (*myStr1).name = "1"
    (*myStr2).name = "2"
    (*myStr3).name = "3"
    (*myStr4).name = "4"

    set.Insert(myStr1)
    set.Insert(myStr2)
    set.Insert(myStr3)
    set.Insert(myStr4)

    myStr11 := new(MyStr)
    myStr21 := new(MyStr)
    myStr31 := new(MyStr)
    myStr41 := new(MyStr)

    (*myStr11).name = "1"
    (*myStr21).name = "2"
    (*myStr31).name = "3"
    (*myStr41).name = "4"

    set1.Insert(myStr11)
    set1.Insert(myStr21)
    set1.Insert(myStr31)
    set1.Insert(myStr41)

    set.Remove(myStr1)

    nset := set.InterSection(set1)

    var iterator Iterator
    for iterator = nset.GetIterator(); iterator.HasNext(); {
        txt1 := iterator.Next()
        txt2,r := txt1.(*MyStr)
        if (!r) {
            fmt.Println(r)
        } else {
            fmt.Println((*txt2).name)
        }
    }
}

