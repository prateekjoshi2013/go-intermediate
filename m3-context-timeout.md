
### Timers

- A server is a shared resource. Like all shared resources, each user wants to get as much as they can out of it and isn’t terribly concerned with the needs of other users. 

- It’s the responsibility of the shared resource to manage itself so that it provides a fair amount of time to all of its users.

- There are generally four things that a server can do to manage its load:

  - Limit simultaneous requests 
    
    - By limiting the number of goroutines, a server manages simultaneous load

  - Limit how many requests are queued waiting to run  
    
    - The size of the waiting queue is handled via buffered channels

  - **Limit how long a request can run**
    
    - **The context provides a way to control how long a request runs**.

  - Limit the resources a request can use (such as memory or disk space)

- **two different functions to create a ```time-limited context```**. 

- The first is ```context.WithTimeout```. 
  
  - It **takes two parameters, an existing context and time.Duration that specifies the duration until the context automatically cancels**. 
  
  - It **returns a context that automatically triggers a cancellation after the specified duration as well as a cancellation function that is invoked to cancel the context immediately**.

- The second function is ```context.WithDeadline```. 
  
  - **This function takes in an existing context and a time.Time that specifies the time when the context is automatically canceled**. 
  
  - **Like context.WithTimeout, it returns a context that automatically triggers a cancellation after the specified time has elapsed as well as a cancellation function**.
  
  - **If you pass a time in the past to context.WithDeadline, the context is created already canceled**.


#### Parent vs Child context deadline

- You control how long an individual call takes by creating a child context that wraps a parent context using context.WithTimeout or context.WithDeadline.
  
- **Any timeout that you set on the child context is bounded by the timeout set on the parent context;** 
  
  - **if a parent context times out in two seconds, you can declare that a child context times out in three seconds, but when the parent context times out after two seconds, so will the child**.

```go
    ctx := context.Background()
    parent, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()
    child, cancel2 := context.WithTimeout(parent, 3*time.Second)
    defer cancel2()
    start := time.Now()
    <-child.Done()
    end := time.Now()
    fmt.Println(end.Sub(start)) // 2s
```

- we specify a two-second timeout on the parent context and a threesecond timeout on the child context. 
- We then wait for the child context to complete by waiting on the channel returned from the Done method on the child context.Context
- This program ends bounded by timeout of the parent context.


### Handling Context Cancellation in Your Own Code

- *Whenever you call another HTTP service or the database, you ***should pass along the context; those libraries properly handle cancellation via the context****.

- If you do write code that should be interrupted by a context cancellation, you implement the cancellation checks using the concurrency features. 

- **The context.Context type has two methods that are used when managing cancellation**.
  
  - **The ```Done``` method returns a ```channel of struct{}```. (The reason this is the chosen return type is that an empty struct uses no memory.)** 
  
    - **The channel is closed when the context is canceled due to a timer or the cancel function being invoked**. 
    
    - Remember,
     
      - **A closed channel always immediately returns its zero value when you attempt to read it.**
      
      - **A read from a nil channel never returns. If this is not done inside a case in a select statement, your program will hang**

  - **The ```Err``` method returns nil if the context is still active, or it returns one of two sentinel errors if the context has been canceled**: 

    - **```context.Canceled``` - after explicit cancellation**

    - **```context.DeadlineExceeded``` - returned when a timeout triggered cancellation**

- Here’s the pattern for supporting context cancellation in your code:

  - we take the output from the long-running function and put it in the buffered channel. 

  - We then have a select with two cases. 
    
    - In our first select case, we read the data from the long-running function and return it. This is the case that’s triggered if the context isn’t canceled due to timeout or invocation of the cancel function.
    
    - The second select case is triggered if the context is canceled. We return the zero value for the data and the error from the context to tell us why it was canceled.

```go
    func longRunningThingManager(ctx context.Context, data string) (string, error) {
        type wrapper struct {
            result string
            err error
        }

        ch := make(chan wrapper, 1)
        
        go func() {
            // do the long running thing
            result, err := longRunningThing(ctx, data)
            ch <- wrapper{result, err}
        }()
        
        select {
            case data := <-ch:
                return data.result, data.err
            case <-ctx.Done():
                return "", ctx.Err()
        }
    }
```