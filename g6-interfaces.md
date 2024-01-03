- Go’s concurrency model gets all of the publicity, the real star of Go’s design is its implicit **interfaces, the only abstract type in Go**.

- Definition of the **Stringer interface in the fmt package**:

    ```go
        type Stringer interface {
            String() string
        }
    ```

- **In an interface declaration, an interface literal appears after the name of the interface type**.
  
- **The methods defined by an interface are called the method set of the interface**.

- **Interfaces can be declared in any block**.

- ***Interfaces are usually named with **“er”** endings***.

  - examples: **io.Reader**, **io.Closer**, **io.ReadCloser**, **json.Marshaler**, and **http.Handler**

### Interfaces Are Type-Safe Duck Typing

- Go’s interfaces are special in that they are **implemented implicitly**.
  
- **A concrete type does not declare that it implements an interface**. 

- **If the method set for a concrete type contains all of the methods in the method set for an interface, the concrete type implements the interface**. 

- This means that the **concrete type can be assigned to a variable or field declared to be of the type of the interface**.

- This implicit behavior makes interfaces the most interesting thing about types in Go, because they **enable both type-safety and decoupling**.

- **The client code defines the interface to specify what functionality it requires**.

- ***If there’s an interface in the standard library that describes what your code needs, use it!***

- It’s perfectly fine for a type that meets an interface to specify additional methods that aren’t part of the interface. One set of client code may not care about those methods, but others do. 

  - For example, the **io.File type also meets the io.Writer interface. If your code only cares about reading from a file, use the io.Reader interface to refer to the file instance and ignore the other methods**

### Embedding and Interfaces

- Like you can embed a type in a struct, you can also embed an interface in an interface.
  
- For example, the io.ReadCloser interface is built out of an io.Reader and an io.Closer:

    ```go
        type Reader interface {
            Read(p []byte) (n int, err error)
        }
        type Closer interface {
            Close() error
        }
        type ReadCloser interface {
            Reader
            Closer
        }
    ```

- **Just like we can embed a concrete type in a struct, you can also embed an interface in a struct**.

## Accept Interfaces, Return Structs

- **The business logic invoked by your functions should be invoked via interfaces, but the output of your functions should be a concrete type**.

- **If you create an API that returns interfaces, you are losing one of the main advantages of implicit interfaces: decoupling**. 

  - **You want to limit the third-party interfaces that your client code depends on because your code is now permanently dependent on the module that contains those interfaces, as well as any dependencies of that module, and so on. This limits future flexibility**.

- **Another reason to avoid returning interfaces is versioning**. 
  
  - **If a concrete type is returned, new methods and fields can be added without breaking existing code. The same is not true for an interface**.
  
  - **Adding a new method to an interface means that you need to update all existing implementations of the interface, or your code breaks**.
  
  - **If you make a backward-breaking change to an API, you should increment your major version number**.

- **Rather than writing a single factory function that returns different instances behind an interface based on input parameters, try to write separate factory functions for each concrete type**. 
  
  - **In some situations (such as a parser that can return one or more different kinds of tokens)it’s unavoidable and you have no choice but to return an interface**.
  
  - **Errors are an exception to this rule**

### Interfaces and nil

- We use nil to represent the zero value for an interface instance, but it’s not as simple as it is for concrete types.
- In order for an interface to be considered nil both the type and the value must be nil. 

- The following code prints out true on the first two lines and false on the last:

    ```go
        var s *string
        fmt.Println(s == nil) // prints true
        var i interface{}
        fmt.Println(i == nil) // prints true
        i = s
        fmt.Println(i == nil) // prints false
    ```

- Interfaces are implemented as a pair of pointers, one to the underlying type and one to the underlying value. 

  - **As long as the type is non-nil, the interface is non-nil**. (Since you cannot have a variable without a type, if the value pointer is non-nil, the type pointer is always non-nil.)
    
  - **What nil indicates for an interface is whether or not you can invoke methods on it**.
  
  - **If an interface is nil, invoking any methods on it triggers a panic**


### The Empty Interface Says Nothing

- Sometimes in a statically typed language, you need a way to say that a variable could
store a value of any type. Go uses interface{} to represent this:

    ```go
    var i interface{}
    i = 20
    i = "hello"
    i = struct {
        FirstName string
        LastName string
    } {"Fred", "Fredson"}
    ```

- You should note that interface{} isn’t special case syntax. 
  
  - **An empty interface type simply states that the variable can store any value whose type implements zero or more methods**. 
  
  - **This just happens to match every type in Go. Because an empty interface doesn’t tell you anything about the value it represents, there isn’t a lot you can do with it**. 
  
  - One common **use of the empty interface is as a placeholder for data of uncertain schema that’s read from an external source, like a JSON file**
  
  - If you see a function that takes in an empty interface, it’s likely that it is using reflection to either populate or read the value. 
  
  - These situations should be relatively rare. Avoid using interface{}. Go is designed as a strongly typed language and attempts to work around this are unidiomatic.

### Type Assertion and Type Switches

- If you find yourself in a situation where you had to store a value into an empty interface, you might be wondering how to read the value back again. To do that, we need to look at type assertions and type switches

    #### Type Assertion

    - **A type assertion names the concrete type that implemented the interface, or names another interface that is also implemented by the concrete type underlying the interface**.

    ```go
    type MyInt int

    func main() {
        var i interface{}
        var mine MyInt = 20
        i = mine
        i2 := i.(MyInt)
        fmt.Println(i2 + 1)
    }
    ```  

    - **if a type assertion is wrong. In that case, your code panics**
    
    - **Even if two types share an underlying type, a type assertion must match the type of the underlying value**

    - if panicking is not desired behavior. We avoid this by using the **comma ok idiom**

        ```go
            i2, ok := i.(int)
            if !ok {
            return fmt.Errorf("unexpected type for %v",i)
            }
            fmt.Println(i2 + 1)
        ```

    - **Type assertions can only be applied to interface types and are checked at runtime**.


    #### Type Switches
    
    - When an interface could be one of multiple possible types, use a type switch instead:

    ```go
        func doThings(i interface{}) {
            switch j := i.(type) {
            case nil:
            // i is nil, type of j is interface{}
            case int:
            // j is of type int
            case MyInt:
            // j is of type MyInt
            case io.Reader:
            // j is of type io.Reader
            case string:
            // j is a string
            case bool, rune:
            // i is either a bool or rune, so j is of type interface{}
            default:
            // no idea what i is, so j is of type interface{}
            }
        }
    ```

    - You can **use nil for one case to see if the interface has no associated type**.
  
    - **If you list more than one type on a case, the new variable is of type interface{}**.
    
    - you can have **a default case that matches when no specified type does**.

    #### Use Type Assertions and Type Switches Sparingly

    - **While it might seem handy to be able to extract the concrete implementation from an interface variable, you should use these techniques infrequently**. 
    
    - For the most part, **treat a parameter or return value as the type that was supplied and not what else it could be**
    
    - There are use cases where type assertions and type switches are useful. 
      
      - One common use of a type assertion is to see if the concrete type behind the interface also implements another interface.

        ```go
        // copyBuffer is the actual implementation of Copy and CopyBuffer.
        // if buf is nil, one is allocated.
        
        func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
            // If the reader has a WriteTo method, use it to do the copy.
            // Avoids an allocation and a copy.
            if wt, ok := src.(WriterTo); ok {
            return wt.WriteTo(dst)
            }
            // Similarly, if the writer has a ReadFrom method, use it to do the copy.
            if rt, ok := dst.(ReaderFrom); ok {
            return rt.ReadFrom(src)
            }
            // function continues...
        }
        ```

    ### Drawback of Optional Interface Technique

    - Drawback to the optional interface technique. We saw earlier that it is common for implementations of interfaces to use the decorator pattern to wrap other implementations of the same interface to layer behavior. 
  
    - The problem is that if there is an optional interface implemented by one of the wrapped implementations, you cannot detect it with a type assertion or type switch


    ```go

        package main

        import "fmt"

        // Interface with required method
        type MyInterface interface {
            RequiredMethod()
        }

        // Optional interface with an optional method
        type OptionalInterface interface {
            OptionalMethod()
        }

        // Implementation of MyInterface
        type MyType struct{}

        func (mt MyType) RequiredMethod() {
            fmt.Println("Required method called.")
        }

        // Implementation of OptionalInterface
        type OptionalType struct{}

        func (ot OptionalType) OptionalMethod() {
            fmt.Println("Optional method called.")
        }

        // Decorator struct that embeds MyInterface
        type MyDecorator struct {
            MyInterface
        }

        // OptionalMethodDecorator adds the optional method to the decorator
        func (md MyDecorator) OptionalMethodDecorator() {
            fmt.Println("Optional method from decorator called.")
        }

        func main() {
            // Wrap the MyType implementation with the decorator
            decoratedInstance := MyDecorator{MyType{}}

            // Call the required method
            decoratedInstance.RequiredMethod()

            // Check if the decorated instance implements OptionalInterface
            if optionalInstance, ok := interface{}(decoratedInstance).(OptionalInterface); ok {
                optionalInstance.OptionalMethod()
            } else {
                fmt.Println("Optional method not implemented.")  
            }

            // Call the optional method added by the decorator
            decoratedInstance.OptionalMethodDecorator()
        }
    ```

    ###  Protect from unexpected interface implementations

    - You can further protect yourself from unexpected interface implementations by making the interface unexported and at least one method unexported.
    
    - If the interface is exported, then it can be embedded in a struct in another package, making the struct implement the interface


### Interface implementation on function type

- Go allows methods on any user-defined type, including user-defined function types. 
- They allow functions to implement interfaces. 
- The most common usage is for HTTP handlers.
  
  - An HTTP handler processes an HTTP server request. It’s defined by an interface:

    ```go
        type Handler interface {
            ServeHTTP(http.ResponseWriter, *http.Request)
        }
    ```

  - By using a type conversion to http.HandlerFunc, any function that has the signature ```func(http.ResponseWriter,*http.Request)``` can be used as an http.Handler:

    ```go
        type HandlerFunc func(http.ResponseWriter, *http.Request)
        
        func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
            f(w, r)
        }
    ```
    - This **lets you implement HTTP handlers using functions, methods, or closures using the exact same code path as the one used for other types that meet the http.Handler interface**.

    ### when should your function or method specify an input parameter of a function type and when should you use an interface ?

    - Go encourages small interfaces, and an interface of only one method could easily replace a parameter of function type.
    - **If your single function is likely to depend on many other functions or other state that’s not specified in its input parameters, use an interface parameter and define a function type to bridge a function to the interface**. 
      
      - ***That’s what’s done in the http package; it’s likely that a Handler is just the entry point for a chain of calls that needs to be configured***. 
    
    - **However, if it’s a simple function (like the one used in sort.Slice), then a parameter of function type is a good choice**


### Dependency Injection

- One of the surprising benefits of Go’s implicit interfaces is that they make dependency injection an excellent way to decouple your code. While developers in other languages often use large, complicated frameworks to inject their dependencies, the truth is that it is easy to implement dependency injection in Go without any additional libraries.



```go

    func LogOutput(message string) {
        fmt.Println(message)
    }

    type SimpleDataStore struct {
        userData map[string]string
    }

    func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
        name, ok := sds.userData[userID]
        return name, ok
    }


    func NewSimpleDataStore() SimpleDataStore {
        return SimpleDataStore{
            userData: map[string]string{
                "1": "Fred",
                "2": "Mary",
                "3": "Pat",
            },
        }
    }

    type DataStore interface {
        UserNameForID(userID string) (string, bool)
    }

    type Logger interface {
        Log(message string)
    }

    type LoggerAdapter func(message string)

    func (lg LoggerAdapter) Log(message string) {
        lg(message)
    }

    type SimpleLogic struct {
        l Logger
        ds DataStore
    }

    func (sl SimpleLogic) SayHello(userID string) (string, error) {
        sl.l.Log("in SayHello for " + userID)
        name, ok := sl.ds.UserNameForID(userID)
        
        if !ok {
            return "", errors.New("unknown user")
        }
        return "Hello, " + name, nil
    }

    func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
        sl.l.Log("in SayGoodbye for " + userID)
        name, ok := sl.ds.UserNameForID(userID)
        if !ok {
            return "", errors.New("unknown user")
        }
        return "Goodbye, " + name, nil
    }

    type Logic interface {
        SayHello(userID string) (string, error)
    }

    type Controller struct {
        l Logger
        logic Logic
    }

    func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
        c.l.Log("In SayHello")
        userID := r.URL.Query().Get("user_id")
        message, err := c.logic.SayHello(userID)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }
        w.Write([]byte(message))
    }

    func NewController(l Logger, logic Logic) Controller {
        return Controller{
            l: l,
            logic: logic,
        }
    }

    func main() {
        // only main function knows concrete types being used 
        // while the function parameters are all interfaces
        l := LoggerAdapter(LogOutput)
        ds := NewSimpleDataStore()
        logic := NewSimpleLogic(l, ds)
        c := NewController(l, logic)
        http.HandleFunc("/hello", c.SayHello)
        http.ListenAndServe(":8080", nil)
    }

```

  - We have a struct with two fields, ***one a Logger, the other a DataStore***. 
  - **There’s nothing in SimpleLogic that mentions the concrete types, so there’s no dependency on them**. 
  - **There’s no problem if we later swap in new implementations from an entirely different provider, because the provider has nothing to do with our interface**.
  
  - **The line http.HandleFunc("/hello", c.SayHello) **
    
    - First, **we are treating the SayHello method as a function**.
    
    - Second, **the http.HandleFunc function takes in a function and converts it to an http HandlerFunc function type, which declares a method to meet the http.Handler interface, which is the type used to represent a request handler in Go**. 
  
      - We took a method from one type and converted it into another type with its own method. That’s pretty neat.