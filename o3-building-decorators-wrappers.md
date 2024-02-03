- Go lets you do with reflection is create a function
- We can use this technique to wrap existing functions with common functionality without writing repetitive code.
- here’s a factory function that adds timing to any function that’s passed into it:

```go
    func MakeTimedFunction(f interface{}) interface{} {
        ft := reflect.TypeOf(f)
        fv := reflect.ValueOf(f)
        wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
            start := time.Now()
            out := fv.Call(in)
            end := time.Now()
            fmt.Println(end.Sub(start))
            return out
        })
        return wrapperF.Interface()
    }
```

- This function takes in any function, so the parameter is of type interface{}

- It then passes the reflect.Type that represents the function into reflect.MakeFunc, along with a closure that captures the start time, calls the original function using reflection captures the end time, prints out the difference, and returns the value calculated by the original function.

- The value returned from reflect.MakeFunc is a reflect.Value, and we call its Interface method to get the value to return.

### Usage

```go
    func timeMe(a int) int {
        time.Sleep(time.Duration(a) * time.Second)
        result := a * 2
        return result
    }
    
    func main() {
        timed:= MakeTimedFunction(timeMe).(func(int) int)
        fmt.Println(timed(2))
    }
```
### Reflection Can’t Make Methods

- While we can use reflection to create new functions and new struct types, there’s no way to use reflection to add methods to a type. 

- This means you cannot use reflection to create a new type that implements an interface.

