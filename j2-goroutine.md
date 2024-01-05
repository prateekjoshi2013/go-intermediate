- **Goroutines are lightweight processes managed by the Go runtime**.

- **When a Go program starts, the Go runtime creates a number of threads and launches a single goroutine to run your program**.
  
  - **All of the goroutines created by your program, including the initial one, are assigned to these threads automatically by the Go runtime scheduler, just as the operating system schedules threads across CPU cores**. 
  
  - *This might seem like extra work, since the underlying operating system already includes a scheduler that manages threads and processes, but it has several benefits*:

  - Goroutine creation is faster than thread creation, because **you aren’t creating an operating system–level resource**.

  - **Goroutine initial stack sizes are smaller than thread stack sizes and can grow as needed**. This makes goroutines more memory efficient.

- **Switching between goroutines is faster than switching between threads because it happens entirely within the process, avoiding operating system calls that are (relatively) slow**.

- The **scheduler is able to optimize its decisions because it is part of the Go process**.
  
  - The **scheduler works with the network poller, detecting when a goroutine can be unscheduled because it is blocking on I/O**. 
  
  - It also **integrates with the garbage collector, making sure that work is properly balanced across all of the operating system threads assigned to your Go process**.

### Launching GoRoutines

- **Any function can be launched as a goroutine**. 
  
- However, **it is customary in Go to launch goroutines with a closure that wraps business logic. The closure takes care of the concurrent bookkeeping**. 
  
  - For example, **the closure reads values out of channels and passes them to the business logic which is completely unaware that it is running in a goroutine**. 
  
  - **The result of the function is then written back to a different channel**. 
  
  - **This separation of responsibility makes your code modular, testable, and keeps concurrency out of your APIs**:

```go
    func process(val int) int {
        // do something with val
    }

    func runThingConcurrently(in <-chan int, out chan<- int) {
        go func() {
            for val := range in {
                result := process(val)
                out <- result
            }
        }()
    }
```