### functions

    ```go
        func div(numerator int, denominator int) int, error {
            if denominator == 0 {
                return 0, errors.New("cannot divide by zero")
            }
            return numerator / denominator, nil
        }
    ```

 
- this function is taking two int params and returning an int and error

- When you have multiple input parameters of the same type, you can write your input parameters like this:

    ```go
        func div(numerator, denominator int) int {
    ```

### Simulating Named and Optional Parameters

- **Go doesn’t have: named and optional input parameters**
- If you want **to emulate named and optional parameters, define a struct that has fields that match the desired parameters, and pass the struct to your function**.
  
  ```go
    type MyFuncOpts struct {
        FirstName string
        LastName string
        Age int
    }

    func MyFunc(opts MyFuncOpts) error {
        // do something here
    }

    func main() {

        MyFunc(MyFuncOpts {
            LastName: "Patel",
            Age: 50,
        })

        My Func(MyFuncOpts {
            FirstName: "Joe",
            LastName: "Smith",
        })
    }
  ```

### Variadic Input Parameters and Slices

- The **variadic parameter must be the last (or only) parameter in the input parameter list**. 
  
- You indicate it with **three dots (…) before the type**. 

- The variable that’s created within the function is **a slice of the specified type**. 


```go

    func addTo(base int, vals ...int) []int {
        out := make([]int, 0, len(vals))
        for _, v := range vals {
            out = append(out, base+v)
        }
        return out
    }

    func main() {
        fmt.Println(addTo(3))
        fmt.Println(addTo(3, 2))
        fmt.Println(addTo(3, 2, 4, 6, 8))
        a := []int{4, 3}
        fmt.Println(addTo(3, a...))
        fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
    }

```

### Functions Are Values

- Functions in Go are values.
  
- The type of a function is built out of the keyword func and the types of the parameters and return values.
  
- This combination is called the signature of the function. Any function that has the exact same number and types of parameters and return values meets the type signature.
  
- function signature is ***div(numerator int, denominator int) int, error***

- This means we can do something like 
  
    ```go
        func add(i int, j int) int { return i + j }

        func sub(i int, j int) int { return i - j }

        func mul(i int, j int) int { return i * j }

        var opMap = map[string]func(int, int) int{
            "+": add,
            "-": sub,
            "*": mul,
        }
    ```

### function type declaration 

- **we can declare a type for a function like **
  
    ```go
        type opFuncType func(int,int) int

        var opMap = map[string]opFuncType {
        // same as before
        }
    ```


### Anonymous Functions

- We can also define new functions within a function and assign them to variables.

- **These inner functions are anonymous functions**; 

- **They don’t have a name**. 

- You don’t have to assign them to a variable, either. You can write them inline and call them immediately.

```go

func main() {
    
    for i := 0; i < 5; i++ {
        func(j int) {
            fmt.Println("printing", j, "from inside of an anonymous function")
        }(i)
    }
}
```

- There are two situations where declaring anonymous functions without assigning them to variables is useful: 

  - **defer statements**
  
  - **launching goroutines**