package main

import "fmt"

type struct1 struct {
    arg1 int
}

func main() {
    var obj1 struct1
    obj2 := struct1{1}
    obj3 := struct{ arg1,arg2 int}{1, 2}
    obj4 := &obj1
    obj5 := new(struct1)
    obj5.arg1 = 5

    obj6 := &struct1{6}

    obj4.arg1 = 4
    fmt.Printf("%d %d %d %d %d\n", obj1.arg1, obj2.arg1, obj3.arg2, obj5.arg1, obj6.arg1);
}
