
#### Cancellation

- Imagine that you have *a request that spawns several goroutines, each one calling a different HTTP service*. 

- *If one service returns an error that prevents you from returning a valid result, there is no point in continuing to process the other goroutines*. 

- In Go, **this is called cancellation and the context provides the mechanism for implementation**.

- **To create a cancellable context, use the ```context.WithCancel``` function.** 
  
  - **It takes in a ```context.Context``` as a parameter and returns a ```context.Context``` and a ```context.CancelFunc```**. 
  
  - **The returned context.Context is not the same context that was passed into the function. Instead, it is a child context that wraps the passed-in parent context.Context**. 
  
  - **A context.CancelFunc is a function that cancels the context, telling all of the code that’s listening for potential cancellation that it’s time to stop processing**.


First we’ll set up two servers in a file called servers.go

```servers.go```

```go

    // These functions launch servers when they are called. 
    // We are using the httptest.Server, which makes it easier 
    // to write unit tests with remote servers. 

    // slow server sleeps for two seconds and then returns the message Slow response. 
    func slowServer() *httptest.Server {
        s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            time.Sleep(2 * time.Second)
            w.Write([]byte("Slow response"))
        }))
        return s
    }

    // The other checks to see if there is a query parameter error set to true. 
    // If there is, it returns the message error. Otherwise, it returns the message ok
    func fastServer() *httptest.Server {
        s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.URL.Query().Get("error") == "true" {
                w.Write([]byte("error"))
                return
            }
            w.Write([]byte("ok"))
        }))
        return s
    }
```
- the client portion of the code in a file called client.go:

  - Our callBoth function creates a cancellable context and a cancellation function from the passed-in context. 

  - **By convention, this function variable is named cancel**. 

  - It is important to **remember that any time you create a cancellable context, you must call the cancel function**. 
    
    - **It is fine to call it more than once; every invocation after the first is ignored**. 
    
    - We **use a defer to make sure that it is eventually called**.

  -  Next, we set up two goroutines and pass the cancellable context, a label, and the URL to callServer, and wait for them both to complete.

  -  If either call to callServer returns an error, we call the cancel function.
  
```client.go```

```go
    var client = http.Client{}

    func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()
        var wg sync.WaitGroup
        wg.Add(2)
        go func() {
            defer wg.Done()
            err := callServer(ctx, "slow", slowURL)
            if err != nil {
                cancel()
            }
        }()
        go func() {
            defer wg.Done()
            err := callServer(ctx, "fast", fastURL+"?error="+errVal)
            if err != nil {
                cancel()
            }
        }()
        wg.Wait()
        fmt.Println("done with both")
    }


    func callServer(ctx context.Context, label string, url string) error {
        // creating request with the cancellable context sent 
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
        if err != nil {
            fmt.Println(label, "request err:", err)
            return err
        }
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println(label, "response err:", err)
            return err
        }
        data, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(label, "read err:", err)
            return err
        }
        result := string(data)
        if result != "" {
            fmt.Println(label, "result:", result)
        }
        if result == "error" {
            fmt.Println("cancelling from", label)
            return errors.New("error happened")
        }
        return nil
    }
```

- the main function, which kicks off the program, in the file main.go:

```main.go```

```go
    func main() {
        ss := slowServer()
        defer ss.Close()
        fs := fastServer()
        defer fs.Close()
        ctx := context.Background()
        callBoth(ctx, os.Args[1], ss.URL, fs.URL)
    }
```
