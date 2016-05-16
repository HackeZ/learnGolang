package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    
    // G2
    go func() {
        fmt.Println("The First Goroutine!")
        wg.Done()
    }()
    
    // G3
    go func() {
        fmt.Println("The Second Goroutine!")
        wg.Done()
    }()
    
    // G4
    go func() {
        fmt.Println("The Thrid Goroutine!")
        wg.Done()
    }()
    
    wg.Wait()
    fmt.Println("G2, G3, G4 are ended.")
}