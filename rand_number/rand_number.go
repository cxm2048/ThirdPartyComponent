package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // 5 - 30的随机数
    //rand.Seed(time.Now().UnixNano())
    for i := 0; i < 100; i++ {
        r := rand.Intn(30-5) + 5
        fmt.Print(r)
        fmt.Print(" ")
    }
}
