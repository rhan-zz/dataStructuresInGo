package base

// just like Object in java
type Object interface{}

type ObjectPtr interface{}

// an function type for match data
type  MatchFun func (data1 Object, data2 Object) int

func whatType(x interface{}) string {
    switch x.(type){
        case int:
            return "int"
        case string:
            return "字符串"
        default:
            return "未知"
    }
}