package main

import (
    "fmt"
    . "lists"
    . "base"
)

type MyStr struct {
    name string
}

func match(data1 Object, data2 Object) int {
    myStr1 := data1.(*MyStr)
    myStr2 := data2.(*MyStr)

    fmt.Println("myMatch")

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
    }*/

    dLst := new (DList)
    a := 1
    b := 2
    c := 3

    var d int

    dLst.Append(a)
    dLst.Append(b)
    dLst.Append(c)

    node := dLst.GetHead().GetNext()
    fmt.Println(dLst.GetSize())

    d = dLst.Remove(node).(int)

    fmt.Println(dLst.GetSize())
    fmt.Println(d)
}

