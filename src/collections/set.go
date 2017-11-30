/*
set
author:heavy
version:1.0
*/

package collections

import (
    . "base"
    )

type Set struct {
    list *List
}

type SetIterator struct {
    index uint64
    set *Set
}

func (set *Set) getAt(i uint64) Object {
    return (*set).list.GetAt(i)
}

func (set *Set) getSize() uint64 {
    return (*set).list.GetSize()
}

func (set *Set) Init(match ...MatchFun) {
    lst := new(List)
    (*set).list = lst

    if len(match) == 0 {
        lst.Init() 
    } else {
        lst.Init(match[0])
    }
}

func (set *Set) Insert(data Object) bool {
    if (!set.IsMember(data)) {
        return (*set).list.Append(data)
    }
    
    return false
}

func (set *Set) Remove(data Object) bool {
    return (*set).list.Remove(data)
}

func (set *Set) IsEmpty() bool {
    return (*set).list.IsEmpty()
}

func (set *Set) IsMember(data Object) bool {
    return (*set).list.IsMember(data);
}

func (set *Set) Union(set1 *Set) *Set {
    if (set1 == nil) {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).list)).myMatch)

    if (set.IsEmpty() && set1.IsEmpty()) {
        return nSet
    }

    for i := uint64(0); i < set.getSize(); i++ {
        nSet.Insert(set.getAt(i))
    }


    var data Object
    for i := uint64(0); i < set1.getSize(); i++ {
        data = set1.getAt(i)
        if (!nSet.IsMember(data)) {
            nSet.Insert(data)
        }
    }

    return nSet
}

func (set *Set) InterSection(set1 *Set) *Set {
    if (set1 == nil) {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).list)).myMatch)

    if (set.IsEmpty() || set1.IsEmpty()) {
        return nSet
    }

    fSet := set
    sSet := set1
    lenth := set.getSize()
    if (set1.getSize() < lenth) {
        fSet = set1
        sSet = set
    }

    var data Object

    for i := uint64(0) ; i < lenth; i++ {
        data = fSet.getAt(i)

        if (sSet.IsMember(data)) {
            nSet.Insert(data)
        }
    }

    return nSet
}

func (set *Set) Difference(set1 *Set) *Set {
    if (set1 == nil) {
        return nil
    }

    nSet := new(Set)
    nSet.Init((*((*set).list)).myMatch)

    if (set.IsEmpty()) {
        return nSet
    }

    var data Object
    for i := uint64(0); i < set.getSize(); i++ {
        data = set.getAt(i)

        if (!set1.IsMember(data)) {
            nSet.Insert(data)
        }
    }

    return nSet
}

func (set *Set) IsSubSet(subSet *Set) bool {
    if (set == nil) {
        return false
    }

    if (subSet == nil) {
        return true
    }

    for i := uint64(0); i < subSet.getSize(); i++ {
        if (!(set.IsMember(subSet.getAt(i)))) {
            return false
        }
    }

    return true
}

func (set *Set) Equals(set1 *Set) bool {
    if (set == nil || set1 == nil) {
        return false
    }

    if (set.IsEmpty() && set1.IsEmpty()) {
        return true
    }

    nSet := set.InterSection(set1)

    return (set.getSize() == nSet.getSize())
}

func (set *Set) GetIterator() *SetIterator {
    iterator := new(SetIterator)
    (*iterator).index = 0
    (*iterator).set = set

    return iterator
}

func (iterator *SetIterator) HasNext() bool {
    set := (*iterator).set
    index := (*iterator).index

    return (index < set.getSize())
}

func (iterator *SetIterator) Next() Object {
    set := (*iterator).set
    index := (*iterator).index

    if (index < set.getSize()) {
        data := set.getAt(index)
        (*iterator).index++

        return data    
    }

    return nil
}