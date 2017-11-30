/*
set
author:heavy
version:1.0
*/

package collections

import (
    . "base"
)


type Iterator interface{
    HasNext() bool
    Next() Object
}