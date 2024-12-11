```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Added mutex
        
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Acquire lock before accessing
                        count++
                        mu.Unlock() // Release lock after accessing
                }()
        }

        wg.Wait()
        fmt.Println("count:", count)
}
```