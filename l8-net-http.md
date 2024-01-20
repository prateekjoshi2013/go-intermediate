### Client

```net/http``` library includes a **production quality HTTP/2 client and server**.

#### Instantiating Client

- The **net/http package defines a Client type to make HTTP requests and receive HTTP responses**.

- **A default client instance (cleverly named ```DefaultClient```) is found in the net/http package, but you should avoid using it in production applications, because it defaults to having no timeout**. 

- **Instead, instantiate your own. You only need to create a single http.Client for your entire program, as it properly handles multiple simultaneous requests across goroutines**:

    ```go
        client := &http.Client{
            Timeout: 30 * time.Second,
        }
    ```

#### Making Request

- When you want **to make a request, you create a new ```*http.Request``` instance with the ```http.NewRequestWithContext``` function, passing it a context, the method, and URL that you are connecting to**. 

- **If you are making a PUT, POST, or PATCH request, specify the body of the request with the last parameter as an io.Reader**. 

- **If there is no body, use nil**:

    ```go
        req, err := http.NewRequestWithContext(context.Background(), 
        http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
        if err != nil {
            panic(err)
        }
    ```
#### Preparing Request

- Once you have an *http.Request instance, you can set any headers via the Headers field of the instance. 

```go
    req.Header.Add("X-My-Client", "Learning Go")
```

#### Send Request

- Call the Do method on the http.Client with your http.Request and the result is returned in an http.Response:

```go
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
```

#### Processing Response

- The **response has several fields with information on the request**. 

  - The **numeric code of the response status is in the ```StatusCode``` field**, 
  - **the text of the response code is in the ```Status``` field**, 
  - **the response headers are in the ```Header``` field**, 
  - and any **returned content is in a ```Body``` field of type io.ReadCloser. This allows us to use it with ```json.Decoder``` to process REST API responses**:

  ```go
      defer res.Body.Close()
      if res.StatusCode != http.StatusOK {
          panic(fmt.Sprintf("unexpected status: got %v", res.Status))
      }
      fmt.Println(res.Header.Get("Content-Type"))
      var data struct {
          UserID int `json:"userId"`
          ID int `json:"id"`
          Title string `json:"title"`
          Completed bool `json:"completed"`
      }
      err = json.NewDecoder(res.Body).Decode(&data)
      if err != nil {
          panic(err)
      }
      fmt.Printf("%+v\n", data)
  ```

- **There are functions in the net/http package to make GET, HEAD, and POST calls. Avoid using these functions because they use the default client, which means they don’t set a request timeout**.

### Server

- The **HTTP Server is built around the concept of an ```http.Server``` and the ```http.Handler interface```**

- **```http.Server``` is responsible for ```listening for HTTP requests```**. 

- It is a **performant HTTP/2 server that supports TLS**.

- **A request to a server is handled by an ```implementation of the http.Handler interface``` that’s assigned to the Handler field**. 
  
  - This interface defines a single method:

    ```go
        type Handler interface {
            ServeHTTP(http.ResponseWriter, *http.Request)
        }
    ```

- The **\*http.Request should look familiar, as it’s the exact same type that’s used to send a request to an HTTP server**. 

  - The **http.ResponseWriter is an interface with three methods**:

    ```go
        type ResponseWriter interface {
            Header() http.Header
            Write([]byte) (int, error)
            WriteHeader(statusCode int)
        }
    ```
  - These **methods must be called in a specific order**: 
  
    - First, **call Header to get an instance of http.Header and set any response headers you need. If you don’t need to set any headers, you don’t need to call it**. 
    
    - Next, **call WriteHeader with the HTTP status code for your response. (All the status codes are defined as constants in the net/http package; all status code constants are untyped integers.)** 
    
      - If you are sending a response that has a 200 status code, you can skip WriteHeader
    
    - Finally, **call the Write method to set the body for the response**.

    ```go
        type HelloHandler struct{}

        func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello!\n"))
        }
    ```

#### Instantiate Server

- You instantiate a new http.Server just like any other struct:

    ```go
        s := http.Server{
            Addr: ":8080",
            ReadTimeout: 30 * time.Second,
            WriteTimeout: 90 * time.Second,
            IdleTimeout: 120 * time.Second,
            Handler: HelloHandler{},
        }
        err := s.ListenAndServe()
        if err != nil {
            if err != http.ErrServerClosed {
                panic(err)
            }
        }
    ```

- The **Addr field specifies the host and port the server listens on.** If you don’t specify them, your **server defaults to listening on all hosts on the standard HTTP port, 80**.

- You **specify timeouts for the server’s reads, writes, and idles using time.Duration values. Be sure to set these to properly handle malicious or broken HTTP clients, as the default behavior is to not time out at all**. 

- Finally, **you specify the http.Handler for your server with the Handler field**.

#### http.ServerMux

- **A server that only handles a single request isn’t terribly useful, so the Go standard library includes a request router, ```*http.ServeMux```**. 

- You **create an instance with the ```http.NewServeMux``` function**. 

- **It meets the http.Handler interface, so it can be assigned to the Handler field in http.Server**. 

- **It also includes two methods that allow it to dispatch requests**. 
  
  - **The first method is simply called Handle and takes in two parameters, a path and an http.Handler. If the path matches, the http.Handler is invoked**.

- While you could create implementations of http.Handler, the **more common pattern is to use the HandleFunc method on ```*http.ServeMux```**:

```go
    mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!\n"))
    })
```  

- **This method takes in a function or closure and converts it to a http.HandlerFunc**.

- Example:

- Because an *http.ServeMux dispatches requests to http.Handler instances, and since the *http.ServeMux implements http.Handler, you can create an *http.ServeMux instance with multiple related requests and register it with a parent *http.ServeMux:

- In this example, a request for /person/greet is handled by handlers attached to person, while /dog/greet is handled by handlers attached to dog. When we register person and dog with mux, **we use the http.StripPrefix helper function to remove the part of the path that’s already been processed by mux**.
  
```go
    person := http.NewServeMux()
    person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("greetings!\n"))
    })
    dog := http.NewServeMux()
    dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("good puppy!\n"))
    })
    mux := http.NewServeMux()
    mux.Handle("/person/", http.StripPrefix("/person", person))
    mux.Handle("/dog/", http.StripPrefix("/dog", dog))
```

### Middleware

- One of the most common requirements of an HTTP server is **to perform a set of actions across multiple handlers, such as checking if a user is logged in, timing a request, or checking a request header**. 

- **Go handles these cross-cutting concerns with the middleware pattern**. 

- **Rather than using a special type, the middleware pattern uses a function that takes in an http.Handler instance and returns an http.Handler**.

- Usually, **the returned http.Handler is a closure that is converted to an http.HandlerFunc**.

- Here are **two middleware generators, one that provides timing of requests and another that uses perhaps the worst access controls imaginable**:

  - We do setup operations or checks. If the checks don’t pass, we write the output in the middleware (usually with an error code) and return. If all is well, we call the handler’s ServeHTTP method. When that returns, we run cleanup operations.

  ```go
      func RequestTimer(h http.Handler) http.Handler {
          return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
              start := time.Now()
              h.ServeHTTP(w, r)
              end := time.Now()
              log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
          })
      }
  ```
  - The TerribleSecurityProvider shows how to create configurable middleware. You pass in the configuration information (in this case, the password), and the function returns middleware that uses that configuration information. It is a bit of a mind bender, as it returns a closure that returns a closure.
  
  - The TerribleSecurityProvider shows how to create configurable middleware. You pass in the configuration information (in this case, the password), and the function returns middleware that uses that configuration information.
  
    ```go
        var securityMsg = []byte("You didn't give the secret password\n")

        func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
            return func(h http.Handler) http.Handler {
                return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                    if r.Header.Get("X-Secret-Password") != password {
                        w.WriteHeader(http.StatusUnauthorized)
                        w.Write(securityMsg)
                        return
                    }
                    h.ServeHTTP(w, r)
                })
            }
        }
    ```

#### Passing values through layers

- **to pass values through the layers of middleware. This is done via the context**

#### Configuring Middleware

- We add middleware to our request handlers by chaining them:

```go
    terribleSecurity := TerribleSecurityProvider("GOPHER")

    mux.Handle("/hello", terribleSecurity(RequestTimer(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello!\n"))
        }))
    ))
```

- We get back our middleware from the TerribleSecurityProvider and then wrap our handler in a series of function calls. 
- This calls the terribleSecurity closure first, then calls the RequestTimer, which then calls our actual request handler.
- Because the *http.ServeMux implements the http.Handler interface, you can apply a set of middleware to all of the handlers registered with a single request router:
  
```go
    terribleSecurity := TerribleSecurityProvider("GOPHER")

    wrappedMux := terribleSecurity(RequestTimer(mux))
    s := http.Server{
        Addr: ":8080",
        Handler: wrappedMux,
    }
```


#### Disadvantages of http.ServerMux

- **The biggest weakness in the HTTP support in the standard library is the built-in ```*http.ServeMux``` request router**. 

  - **It doesn’t allow you to specify handlers based on an HTTP verb or header** 

  - **It doesn’t provide support for variables in the URL path**.

  - **Nesting ```*http.ServeMux``` instances is also a bit clunky**. 

- **There are many, many projects to replace it, but two of the most popular ones are ```gorilla mux``` and ```chi```**. 

- *Both are considered idiomatic because they work with http.Handler and http.HandlerFunc instances*