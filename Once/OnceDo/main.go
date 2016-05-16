package main

import (
    "fmt"
    "time"
    "sync"
)

func main() {
    // Usage:
    // var once sync.Once
    // once.Do(func()  {
    //     fmt.Println("Once Do!")
    // })
    onceDo()
}

func onceDo() {
    var num int
    sign := make(chan bool)
    var once sync.Once
    f := func(ii int) func() {
        return func() {
            num = (num + ii * 2)
            sign <- true
        }
    }
    
    for i := 0; i < 3; i++ {
        fi := f(i + 1)
        go once.Do(fi)
    }
    
    for j := 0; j < 3; j++ {
        select {
            case <-sign:
                fmt.Println("Received a singal.")
            case <-time.After(100 * time.Millisecond):
                fmt.Println("Time out!")
        }
    }
    fmt.Printf("Num --> %d.\n", num)
}