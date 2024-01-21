
### Passing Value through a context

- There is one more use for the context. It also provides a way to pass per-request metadata through your program.

- By default, you should prefer to pass data through explicit parameters. 
  
  - As has been mentioned before, idiomatic Go favors the explicit over the implicit, and this includes explicit data passing. 
  
  - If a function depends on some data, it should be clear where it came from.

- However, **there are some cases where you cannot pass data explicitly**. 

- **The most common situation is an HTTP request handler and its associated middleware. As we have seen, all HTTP request handlers have two parameters, one for the request and one for the response**.

- **If you want to make a value available to your handler in middleware, you need to store it in the context**. 
  
  - *Some possible situations include **extracting a user from a JWT (JSON Web Token)***
  
  - or *creating a **per-request GUID** that is **passed through multiple layers of middleware** and into your handler and business logic*.


- **Value stored in the context can be of any type, there is an idiomatic pattern that’s used to guarantee the key’s uniqueness**. 

- Like the key for a map, the key for context value must be comparable
  
- Create a new, unexported type for the key, based on an int:

```go
    type userKey int
```

- If you use a string or another public type for the type of the key, different packages could create identical keys, resulting in collisions. 

- This causes problems that are hard to debug, such as one package writing data to the context that masks the data written by another package or reading data from the context that was written by another package.

- After declaring your unexported key type, you then declare an unexported constant of that type:

```go
    const key userKey = 1
```

- With both the type and the constant of the key being unexported, no code from outside of your package can put data into the context that would cause a collision. 

- If your package needs to put multiple values into the context, define a different key of the same type for each value, using the iota pattern 
  
  - Since we only care about the constant’s value as a way to differentiate between multiple keys, this is a perfect use for iota.
  
- Next, build an API to place a value into the context and to read the value from the context. 

- Make these functions public only if code outside your package should be able to read and write your context values. 

- The name of the function that creates a context with the value should start with ContextWith. 

- The function that returns the value from the context should have a name that ends with FromContext. 
  
- Here are the implementations of our functions to get and read the user from the context:

```go
    func ContextWithUser(ctx context.Context, user string) context.Context {
        return context.WithValue(ctx, key, user)
    }

    func UserFromContext(ctx context.Context) (string, bool) {
        user, ok := ctx.Value(key).(string)
        return user, ok
    }
```

- middleware that extracts a user ID from a cookie:


```go
    // a real implementation would be signed to make sure
    // the user didn't spoof their identity
    func extractUser(req *http.Request) (string, error) {
        userCookie, err := req.Cookie("user")
        if err != nil {
            return "", err
        }
        return userCookie.Value, nil
    }

    func Middleware(h http.Handler) http.Handler {
        return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
            user, err := extractUser(req)
            if err != nil {
                rw.WriteHeader(http.StatusUnauthorized)
                return
            }
            ctx := req.Context()
            ctx = ContextWithUser(ctx, user)
            req = req.WithContext(ctx)
            h.ServeHTTP(rw, req)
        })
    }
```