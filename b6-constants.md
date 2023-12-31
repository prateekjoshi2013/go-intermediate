### Using const

- It is a way to declare a value is **immutable with the ***const*** keyword**.
  
  ```go
    const x int64 = 10
    
    const (
        idKey = "id"
        nameKey = "name"
    )

    const z = 20 * 10
  ```

- ***const*** in Go is very limited.
  
- Constants in Go are a way to give names to
literals. **They can only hold values that the compiler can figure out at compile time**

- they can be assigned:
- Numeric literals
- true and false
- Strings
- Runes
- The **built-in functions** ***complex***, ***real***, ***imag***, ***len***, and ***cap***
- ***Expressions that consist of operators and the preceding values***
- ***iota***

### typed untyped constants

- Constants can be typed or untyped
  
- Leaving a ***constant untyped***
gives you ***more flexibility***. 

```go
    const x = 10
```

- There are situations where you want a constant to enforce a type. We’ll use **typed** constants when we look at creating enumerations with **iota** 

```go
    var y int = x
    var z float64 = x
    var d byte = x
```