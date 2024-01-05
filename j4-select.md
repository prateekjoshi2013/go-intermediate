- **If you can perform two concurrent operations, which one do you do first? You can’t favor one operation over others, or you’ll never process some cases. This is called starvation.**

- The **select keyword allows a goroutine to read from or write to one of a set of multiple channels**. 

- **It looks a great deal like a blank switch statement**:

```go
    select {
        case v := <-ch:
            fmt.Println(v)
        case v := <-ch2:
            fmt.Println(v)
        case ch3 <- x:
            fmt.Println("wrote", x)
        case <-ch4:
            fmt.Println("got value on ch4, but ignored it")
    }
```

- **Each case in a select is a read or a write to a channel**. 

- **If a read or write is possible for a case, it is executed along with the body of the case. Like a switch, each case in a select creates its own block**.

- What happens **if multiple cases have channels that can be read or written?**
  
  - **The select algorithm is simple: it picks randomly from any of its cases that can go forward; order is unimportant**. 
  
  - **It also cleanly resolves the starvation problem, as no case is favored over another and all are checked at the same time**.

- Another advantage of select choosing at random is that **it prevents one of the most common causes of deadlocks: acquiring locks in an inconsistent order**.
  
  - **If you have two goroutines that both access the same two channels, they must be accessed in the same order in both goroutines, or they will deadlock**. 
  
  - **This means that neither one can proceed because they are waiting on each other. If every goroutine in your Go application is deadlocked, the Go runtime kills your program**



```go
    func main() {
        ch1 := make(chan int)
        ch2 := make(chan int)
        go func() {
            v := 1
            ch1 <- v
            v2 := <-ch2
            fmt.Println(v, v2)
        }()
        v := 2
        ch2 <- v
        v2 := <-ch1
        fmt.Println(v, v2)
    }
```
```sh
    fatal error: all goroutines are asleep - deadlock!
```

        Above program is deadlocked because:

        - Deadlocked because goroutine is blocked until ch1 is read from in main
        - Main is blocked until ch2 can be read from in go routine


#### Using select to avoid deadlocks

```go
    func main() {
        ch1 := make(chan int)
        ch2 := make(chan int)
        
        go func() {
            v := 1
            ch1 <- v
            v2 := <-ch2
            fmt.Println(v, v2)
        }()

        v := 2
        var v2 int
        select {
            case ch2 <- v:
            case v2 = <-ch1:
        }
        fmt.Println(v, v2)
    }
```

- Because a **select checks if any of its cases can proceed, the deadlock is avoided**.
  
- **The goroutine that we launched wrote the value 1 into ch1, so the read from ch1 into v2 in the main goroutine is able to succeed**.

### for-select loop

- Since select is responsible for communicating over a number of channels, it is often embedded within a for loop:

```go
    for {
        select {
            case <-done:
                return
            case v := <-ch:
                fmt.Println(v)
        }
    }
```
- This is so **common that the combination is often referred to as a for-select loop**.

- When using a **for-select loop, you must include a way to exit the loop**.

- **like switch statements, a select statement can have a default clause**. **Also just like switch, default is selected when there are no cases with channels that can be read or written**. 
  
- **If you want to implement a nonblocking read or write on a channel, use a select with a default**. 

- The following **code does not wait if there’s no value to read in ch; it immediately executes the body of the default**:

```go
    select {
        case v := <-ch:
            fmt.Println("read from ch:", v)
        default:
            fmt.Println("no value written to ch")
    }
```
### Think before using default case in for select loop

- **Having a default case inside a for-select loop is almost always the wrong thing to do**. 

- *It will be **triggered every time through the loop when there’s nothing to read or write** for any of the cases*. 
  
- **This makes your for loop run constantly, which uses a great deal of CPU.**