### Pointer Receivers and Value Receivers

- ***In general the following rules apply***: 

  - **If your method modifies the receiver, you must use a pointer receiver**.
  - **If your method needs to handle nil instances, then it must use a pointer receiver**.
  - **If your method doesn’t modify the receiver, you can use a value receiver**.

- ***Whether or not you use a value receiver for a method that doesn’t modify the receiver depends on the other methods declared on the type***. 
  
- **When a type has any pointer receiver methods, a common practice is to be consistent and use pointer receivers for all methods, even the ones that don’t modify the receiver**

    ```go
        type Counter struct {
            total int
            lastUpdated time.Time
        }

        func (c *Counter) Increment() {
            c.total++
            c.lastUpdated = time.Now()
        }

        func (c Counter) String() string {
            return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
        }

        func main(){
            var c Counter
            fmt.Println(c.String())
            
            c.Increment()  // c.Increment() is converted to (&c).Increment() automatically by Go
            // this increments only on a copy of Counter c but not on the original counter c
            fmt.Println(c.String())
        }
    ```

- **Do not write getter and setter methods for Go structs, unless you need them to meet an interface**
  
- **Go encourages you to directly access a field**. Reserve methods for business logic. 
  
- The **exceptions are when you need to update multiple fields as a single operation or when the update isn’t a straightforward assignment of a new value**. 
  
- The Increment method defined earlier demonstrates both of these properties.


### Code Your Methods for nil Instances

- When we call a method on a nil instance.

  - **If it’s a method with a value receiver, you’ll get a panic**, as there is no value being pointed to by the pointer
  - If it’s a **method with a pointer receiver, it can work if the method is written to handle the possibility of a nil instance**

### Methods Are Functions Too

- Methods in Go are so much like functions that you **can use a method as a replacement for a function any time there’s a variable or parameter of a function type**.

    ```go
        type Adder struct {
            start int
        }
        
        func (a Adder) AddTo(val int) int {
            return a.start + val
        }

        myAdder := Adder{start: 10}
        fmt.Println(myAdder.AddTo(5)) // prints 15

        f1 := myAdder.AddTo
        fmt.Println(f1(10)) // prints 20
    ```

### Method Expression

- You can also **create a function from the type itself**. This is called a **method expression**:

```go
    f2 := Adder.AddTo
    fmt.Println(f2(myAdder, 15)) // prints 25
```
- In the case of a method expression, **the first parameter is the receiver for the method**; our **function signature is func(Adder, int) int**
  
- This **feature is used in Dependency Injection**

### Functions Versus Methods

- When you should declare a function and when you should use a method? 

  - **package-level state should be effectively immutable**. 

  - **Any time your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct and that logic should be implemented as a method**. 

  - **If your logic only depends on the input parameters, then it should be a function**.

