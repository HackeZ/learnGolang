package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    
    fmt.Println("")    
    fmt.Println(" ==== int32 and int64 ==== ")    
    var i32 int32
    i32 = 10
    
    // Add of Atomic on int32 and int64
    addi32 := atomic.AddInt32(&i32, 3)
    fmt.Println("addi32 --> ", addi32, " || i32 --> ", i32)
    
    // Reduce of Atomic on int32 and int64
    reducei32 := atomic.AddInt32(&i32, -3)
    fmt.Println("reducei32 --> ", reducei32, " || i32 --> ", i32)
    
    fmt.Println("")
    fmt.Println(" ==== uint32 and uint64 ==== ")
    
    var ui32 uint32
    ui32 = 10
    
    // Add of Atomic on uint32 and uint64
    addui32 := atomic.AddUint32(&ui32, 3)
    fmt.Println("addui32 --> ", addui32, " || ui32 --> ", ui32)
    
    // Reduce of Atomic on uint32 and uint64
    number := -3
    reduceui32 := atomic.AddUint32(&ui32, ^uint32(-number-1))
    fmt.Println("reduceui32 --> ", reduceui32, " || ui32 --> ", ui32)
    
}