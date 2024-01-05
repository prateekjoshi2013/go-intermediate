- **Goroutines communicate using channels**. Like slices and maps, **channels are a built-in type created using the make function**:
    
    ```go
        ch := make(chan int)
    ```

- **Like maps, channels are reference types**. **When you pass a channel to a function, you are really passing a pointer to the channel**. 

- Also **like maps and slices, the zero value for a channel is nil**.


### Reading, Writing, and Buffering

- the **<- operator to interact with a channel**. 

- You **read** from a channel by placing the **<- operator to the left of the channel variable**

- You **write** to a channel by **placing it to the right**:

    ```go
        a := <-ch // reads a value from ch and assigns it to a
        ch <- b // write the value in b to ch
    ```

- **Each value written to a channel can only be read once**. 
  
- If **multiple goroutines are reading from the same channel, a value written to the channel will only be read by one of them**.
  
- It is **rare for a goroutine to read and write to the same channel**.
   
  - **When assigning a channel to a variable or field, or passing it to a function, use an arrow before the chan keyword (```ch <-chan int```) to indicate that the goroutine only reads from the channel**.
  
  - **Use an arrow after the chan keyword (```ch chan<- int```) to indicate that the goroutine only writes to the channel**. 

  - Doing so allows the Go **compiler to ensure that a channel is only read from or written by a function**.

### Unbuffered Channel

- **By default channels are unbuffered. Every write to an open, unbuffered channel causes the writing goroutine to pause until another goroutine reads from the same channel**.
  
- **A read from an open, unbuffered channel causes the reading goroutine to pause until another goroutine writes to the same channel**.
  
- **This means you cannot write to or read from an unbuffered channel without at least two concurrently running goroutines**.

### Buffered Channel

- Go also has **buffered channels**. **These channels buffer a limited number of writes without blocking**. 

- **If the buffer fills before there are any reads from the channel, a subsequent write to the channel pauses the writing goroutine until the channel is read**.

- **Just as writing to a channel with a full buffer blocks, reading from a channel with an empty buffer also blocks**.

- A **buffered channel is created by specifying the capacity of the buffer when creating the channel**:
  
```go
    ch := make(chan int, 10)
```

- The built-in functions len and cap return information about a buffered channel. 
- Use len to find out how many values are currently in the buffer and use cap to find out the maximum buffer size. 
- The capacity of the buffer cannot be changed

### for-range and Channels

You can also read from a channel using a for-range loop:

```go
    for v := range ch {
        fmt.Println(v)
    }
```

- **Unlike other for-range loops, there is only a single variable declared for the channel, which is the value**.
  
- **The loop continues until the channel is closed, or until a break or return statement is reached**.


### Closing a Channel

- **When you are done writing to a channel, you close it using the built-in close function**:

```go
    close(ch)
```

- **Once a channel is closed, any attempts to write to the channel or close the channel again will panic**. 

- Interestingly, **attempting to read from a closed channel always succeeds**.

  - **If the channel is buffered and there are values that haven’t been read yet, they will be returned in order**. 

  - **If the channel is unbuffered or the buffered channel has no more values, the zero value for the channel’s type is returned**.

  - When we read from a channel, how do we **tell the difference between a zero value that was written and a zero value that was returned because the channel is closed**?
    
    - we **use the comma ok idiom to detect whether a channel has been closed or not**:

      ```go
          v, ok := <-ch
      ```

    - **If ok is set to true, then the channel is open**. 
    - **If it is set to false, the channel is closed**.

  - Any time you are reading from a channel that might be closed, use the comma ok idiom to ensure that the channel is still open.

  - The **responsibility for closing a channel lies with the goroutine that writes to the channel**. 

  - Be aware that **closing a channel is only required if there is a goroutine waiting for the channel to close** (**such as one using a for-range loop to read from the channel**). 

- Since a channel is just another variable, **Go’s runtime can detect channels that are no longer used and garbage collect them**.


### How channels Behave

| Action | Unbuffered,open |  Unbuffered,closed |  Buffered,open | Buffered,closed | Nil |
| --------- | --------- | --------- | --------- | --------- | --------- |
| Read | Pause until something is written | Return zero value (use comma ok to see if closed) | Pause if buffer is empty | Return a remaining value in the buffer. If the buffer is empty, return zero value (use comma ok to see if closed) | Hang forever |
| Write | Pause until something is read | PANIC | Pause if buffer is full | PANIC | Hang forever |
| Close | Works | PANIC | Works, remaining values still there | PANIC | PANIC |

### Situations to avoid woth channels

- Must **avoid situations that cause Go programs to panic**. 

- The **standard pattern is to make the writing goroutine responsible for closing the channel when there’s nothing left to write**. 

- **When multiple goroutines are writing to the same channel, this becomes more complicated, as calling close twice on the same channel causes a panic**. 

- Furthermore, **if you close a channel in one goroutine, a write to the channel in another goroutine triggers a panic as well**. 

- **The way to address this is to use a sync.WaitGroup**. 

- **A nil channel can be dangerous as well, but there are cases where it is useful**.
