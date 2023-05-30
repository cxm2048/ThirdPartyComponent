package main

import (
    "fmt"
    "strconv"
)

type testFun func(int) string

func test1(a int) string {
    return "test1:" + strconv.Itoa(a)
}

func test2(a int) string {
    panic("test2")
    return "test2:" + strconv.Itoa(a)
}

func test(i int, f testFun) string {
    return f(i)
}

func RecoverPanic() (b bool) {
    defer func() {
        if x := recover(); x != nil {
            b = true
        }
    }()
    fmt.Println(test(1002, test2))
    return false
}
func main() {
    fmt.Println(test(1002, test1))
    //fmt.Println(test(1002, test2))
    RecoverPanic()
}
