package main

import (
    "fmt"
    "runtime"
    "runtime/debug"
    "sync"
    "sync/atomic"
)

func main() {
    // Stop GC and Star GC befort end this main
    defer debug.SetGCPercent(debug.SetGCPercent(-1))
    var count int32
    newFunc := func() interface{} {
        return atomic.AddInt32(&count, 1)
    }
    pool := sync.Pool{New: newFunc}
    
    v1 := pool.Get()
    fmt.Printf("v1 --> %v\n", v1)
    
    // Pool Get and Put
    pool.Put(newFunc())
    pool.Put(newFunc())
    pool.Put(newFunc())
    v2 := pool.Get()
    fmt.Printf("v2 --> %v\n", v2)
    
    // the influence of GC to Pool
    debug.SetGCPercent(100)
    runtime.GC()
    v3 := pool.Get()
    fmt.Printf("v3 --> %v\n", v3)
    pool.New = nil
    v4 := pool.Get()
    fmt.Printf("v4 --> %v\n", v4)
}