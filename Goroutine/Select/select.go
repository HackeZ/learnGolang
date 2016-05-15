package main

import (
    "time"
    "fmt"
)

func main() {
    ch11 := make(chan int, 1000)
    sign := make(chan bool, 1)
    go func() {
        for i := 0; i < 50; i++ {
            time.Sleep(time.Microsecond)
            ch11 <- i
        }
        close(ch11)
        time.Sleep(3*time.Second)
        sign <- false
    }()
    go func ()  {
        var e int
        ok := true
        for {
            select {
                case e, ok = <-ch11:
                    if !ok {
                        fmt.Println("End.")
                        break
                    } else {
                        fmt.Printf("%d\n", e)
                    }
            }
            if !ok {
                break
            }
        }
    }()
    
    fmt.Printf("sign --> %v\n", <-sign)
}