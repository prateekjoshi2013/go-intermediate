### Pointers

- The **zero value for a pointer is nil**.
- **zero value for slices, maps, functions, channels and interfaces is nil**. All of these types are implemented with pointers.
- The **& is the address operator** .**It precedes a value type and returns the address of the memory location where the value is stored** :
  
    ```go
        x := "hello"
        pointerToX := &x
    ```

- The **\*** is the **indirection operator**. It precedes a variable of pointer type and returns the pointed-to value. 

- This is called **dereferencing**:

    ```go
        x := 10
        pointerToX := &x
        fmt.Println(pointerToX) // prints a memory address
        fmt.Println(*pointerToX) // prints 10
    ```

- Before dereferencing a pointer, you **must make sure that the pointer is non-nil. Your program will panic if you attempt to dereference a nil pointer**:

    ```go
        var x *int
        fmt.Println(x == nil) // prints true
        fmt.Println(*x) // panics
    ```

- The built-in function **new creates a pointer variable. It returns a pointer to a zero value instance of the provided type**:

    ```go
        var x = new(int)
        fmt.Println(x == nil) // prints false
        fmt.Println(*x) // prints 0
    ```
- **The new function is rarely used**

- For **structs, use an & before a struct literal** to create a
pointer instance. 

- You **can’t use an & before a primitive literal (numbers, booleans, and strings) or a constant because they don’t have memory addresses**; they exist only at
compile time.

    ```go
        x := &Foo{}
        var y string
        z := &y
    ```

### Pointers Indicate Mutable Parameters

- If **a pointer is passed to a function, the function gets a copy of the pointer**.
- **This still points to the original data, which means that the original data can be modified by the called function**.

    - There are a couple of related implications of this.
      
      - The first implication is that when you pass a nil pointer to a function, you cannot make the value non-nil. You can only reassign the value if there was a value already assigned to the pointer.
  
      - Since the memory location was passed to the function via call-by-value, we can’t change the memory address

        ```go
            func failedUpdate(g *int) {
                x := 10
                g = &x
            }
            func main() {
                var f *int // f is nil
                failedUpdate(f)
                fmt.Println(f) // prints nil
            }
        ```

        - The second implication of copying a pointer is that if you want the value assigned to a pointer parameter to still be there when you exit the function, you must dereference the pointer and set the value. 
        
        - If you change the pointer, you have changed the copy, not the original. Dereferencing puts the new value in the memory location pointed to by both the original and the copy.

        ```go
            func failedUpdate(px *int) {
                x2 := 20
                px = &x2
            }
            func update(px *int) {
                *px = 20
            }
            func main() {
                x := 10
                failedUpdate(&x)
                fmt.Println(x) // prints 10
                update(&x)
                fmt.Println(x) // prints 20
            }
        ```
### Pointers Are a Last Resort

- Rather than populating a struct by passing a pointer to it into a function, have
the function instantiate and return the struct .

  - Don’t do this

    ```go

        func MakeFoo(f *Foo) error {
            f.Field1 = "val"
            f.Field2 = 20
            return nil
        }
    ```
  - Do this

    ```go
        func MakeFoo() (Foo, error) {
            f := Foo{
                Field1: "val",
                Field2: 20,
            }
            return f, nil
        }
    ```

    - The only time you should use pointer parameters to modify a variable is when the function expects an interface. You see this pattern when working with JSON

    ```go
        f := struct {
            Name string `json:"name"`
            Age int `json:"age"`
        }
        err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
    ```


### Pointer Passing Performance

- If a struct is large enough, there are performance improvements from using a pointer to the struct as either an input parameter or a return value. 

- The time to pass a pointer into a function is constant for all data sizes, roughly one nanosecond. This makes sense, as the size of a pointer is the same for all data types. 

- **Passing a value into a function takes longer as the data gets larger**.

- The behavior for **returning a pointer versus returning a value** is more interesting. 
  
  - For **data structures that are smaller than a mb, it is actually slower to return a pointer type than a value type**. 
  
  - if **data structures are larger than a megabyte, the performance advantage flips**.

### The Zero Value Versus No Value

- **When not working with JSON (or other external protocols), resist the temptation to use a pointer field to indicate no value. While a pointer does provide a handy way to indicate no value, if you are not going to modify the value, you should use a value type instead, paired with a boolean**.