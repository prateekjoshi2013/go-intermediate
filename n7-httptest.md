- **The Go standard library includes the ```net/http/httptest``` package to make it easier to stub HTTP services**.
- An implementation of MathSolver that calls an HTTP service to evaluate expressions

```go
    type RemoteSolver struct {
        MathServerURL string
        Client *http.Client
    }

    func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, rs.MathServerURL+"?expression="+url QueryEscape(expression), nil)

        if err != nil {
            return 0, err
        }
        resp, err := rs.Client.Do(req)
        if err != nil {
            return 0, err
        }
        defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return 0, err
        }
        if resp.StatusCode != http.StatusOK {
            return 0, errors.New(string(contents))
        }
        result, err := strconv.ParseFloat(string(contents), 64)
        if err != nil {
            return 0, err
        }
        return result, nil
    }
```

- Using the httptest library to test this code without standing up a server.
- We want to make sure that the data that’s passed into the function arrives on the server. So in our test function, we define a type called info to hold our input and output and a variable called io that is assigned the current input and output:

```go
    type info struct {
        expression string
        code int
        body string
    }
    var io info
```
- Next, we set up our fake remote server and use it to configure an instance of RemoteSolver:

```go
    server := httptest.NewServer(
        http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
                expression := req.URL.Query().Get("expression")
                if expression != io.expression {
                    rw.WriteHeader(http.StatusBadRequest)
                    fmt.Fprintf(rw, "expected expression '%s', got '%s'",
                    io.expression, expression)
                    return
                }
                rw.WriteHeader(io.code)
                rw.Write([]byte(io.body))
            }))
    defer server.Close()
    rs := RemoteSolver{
        MathServerURL: server.URL,
        Client: server.Client(),
    }
```
- The rest of the function works like every other table test that we’ve seen:

```go
    data := []struct {
        name string
        io info
        result float64
    }{
        {"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22},
        // remaining cases
    }
    
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            io = d.io
            result, err := rs.Resolve(context.Background(), d.io.expression)
            if result != d.result {
                t.Errorf("io `%f`, got `%f`", d.result, result)
            }
            var errMsg string
            if err != nil {
                errMsg = err.Error()
            }
            if errMsg != d.errMsg {
                t.Errorf("io error `%s`, got `%s`", d.errMsg, errMsg)
            }
        })
    }
```


- The interesting thing to note is that the variable io has been captured by two different closures: 
  
  - The one for the stub server and the one for running each test. 
  
  - We write to it in one closure and read it in the other. 
  
  - This is a bad idea in production code, but it works well in test code within a single function.