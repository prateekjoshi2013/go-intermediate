- Go generates a panic whenever there is a situation where the Go runtime is unable to figure out what should happen next. 
  
- This could be due to a programming error (like an attempt to read past the end of a slice) or environmental problem (like running out of memory).

- As soon as a panic happens, the current function exits immediately and any defers attached to the current function start running.

- When those defers complete, the defers attached to the calling function run, and so on, until main is reached. The program then exits with a message and a stack trace.

- If there are situations in your programs that are unrecoverable, you can create your
own panics. The built-in function panic takes one parameter, which can be of any
type

    ```go
        func doPanic(msg string) {
            panic(msg)
        }
        func main() {
            doPanic(os.Args[0])
        }
    ```

### Recover

- **Go provides a way to capture a panic to provide a more graceful shutdown or to prevent shutdown at all**. 

- **The built-in recover function is called from within a defer to check if a panic happened**. 

- **If there was a panic, the value assigned to the panic is returned**. 

- **Once a recover happens, execution continues normally**

- **There’s a specific pattern for using recover. We register a function with defer to handle a potential panic**. 

- **We call recover within an if statement and check to see if a non-nil value was found. You must call recover from within a defer because once a panic happens, only deferred functions are run**.

```go
    func div60(i int) {
        defer func() {
            if v := recover(); v != nil {
                fmt.Println(v)
            }
        }()
        fmt.Println(60 / i)
    }

    func main() {
        for _, val := range []int{1, 2, 0, 6} {
            div60(val)
        }
    }
```


### Should panic and recover be used


- If your program panics, be very careful about trying to continue executing after the panic. It’s very rare that you want to keep your program running after a panic occurs. 

- **If the panic was triggered because the computer is out of a resource like memory or disk space, the safest thing to do is use recover to log the situation to monitoring software and shut down with os.Exit(1)**. 

- **If there’s a programming error that caused the panic, you can try to continue, but you’ll likely hit the same problem again. In the preceding sample program, it would be idiomatic to check for division by zero and return an error if one was passed in**.

- The **reason we don’t rely on panic and recover is that recover doesn’t make clear what could fail. It just ensures that if something fails, we can print out a message and continue**. 

- Idiomatic Go favors code that explicitly outlines the possible failure conditions over shorter code that handles anything while saying nothing.

- There is one **situation where recover is recommended**. 
  
  - **If you are creating a library for third parties, do not let panics escape the boundaries of your public API**. 

  - **If a panic is possible, a public function should use a recover to convert the panic into an error, return it, and let the calling code decide what to do with them**.