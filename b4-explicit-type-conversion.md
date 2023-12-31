- Go doesn’t allow automatic type promotion between variables.

- You must use a type conversion when variable types do not match. 

- Even different-sized integers and floats must be converted to the same type to interact

```go
    var x int = 10
    var y float64 = 30.2
    var z float64 = float64(x) + y
    var d int = x + int(y)
```