# Go Race Condition in Counter Increment

This repository demonstrates a common race condition in Go programs involving concurrent access to a shared variable.

The `bug.go` file shows the problematic code.  Multiple goroutines increment a counter without proper synchronization, leading to an unpredictable and often incorrect final count.

The solution, in `bugSolution.go`, uses a mutex to protect the shared counter, ensuring that only one goroutine can access and modify it at a time.