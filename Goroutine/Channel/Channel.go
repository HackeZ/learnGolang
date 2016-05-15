// package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int, 5)
    sign := make(chan byte, 2)
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
            time.Sleep(1*time.Second)
        }
        close(ch)
        fmt.Println("The Channel is close")
        sign <- 0
    }()
    
    go func() {
        for {
           e, ok := <-ch
           fmt.Printf("Get %d --> (%v)\n", e, ok)
           if !ok {
               break
           } 
           time.Sleep(2*time.Second)
        }
        fmt.Println("Get Done")
        sign <- 1
    }()
    
    // 阻塞在这里，直到两个Gorotine结束
    <-sign
    <-sign    
}