package main

import (
    "fmt"
    "sync"
    "time"
)

func main()  {
    var mutex sync.Mutex
    fmt.Println("Lock the lock --> G0 ")
    mutex.Lock()
    fmt.Println("G0 is Locked!")
    for i := 1; i <= 3; i++ {
        go func(i int) {
           fmt.Printf("Lock the lock --> G%d\n", i)
           mutex.Lock()
           fmt.Printf("G%d is Locked!", i) 
        }(i)
    }
    time.Sleep(time.Second)
    fmt.Println("Unlock th lock --> G0")
    mutex.Unlock()
    fmt.Println("G0 is Unlocked!")
    time.Sleep(time.Second)
}