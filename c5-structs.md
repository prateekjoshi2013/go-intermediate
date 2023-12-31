- A struct type is defined with the keyword type, the name of the struct type, the keyword struct, and a pair of braces ({}). Within the braces, you list the fields in the struct. 

- Just like we put the variable name first and the variable type second in a var declaration, we put the struct field name first and the struct field type second. 

- Unlike map literals, there are no commas separating the fields in a struct
declaration. 

- You can define a struct type inside or outside of a function. A struct type
that’s defined within a function can only be used within that function.

    ```go
        type person struct {
            name string
            age int
            pet string
        }
    ```

- After struct type is declared, we can define variables of that type:
    
    ```go
        var fred person
    ```
  - Here we are using a var declaration. **Since no value is assigned to fred, it gets the zero value for the person struct type**. 
  
  - **A zero value struct has every field set to the field’s zero value**.

- A struct literal can be assigned to a variable as well:

    ```go
        bob := person{}
    ```

  - Unlike maps, **there is no difference between assigning an empty struct literal and not assigning a value at all. Both initialize all of the fields in the struct to their zero values**.

- There are two different styles for a **nonempty struct literal**:
  
  - **A struct literal can be specified as a comma-separated list of values for the fields inside of braces**:

      ```go
          julia := person{
              "Julia",
              40,
              "cat",
          }

      ```

  - The second struct literal style looks like the map literal style:

      ```go
          beth := person{
              age: 30,
              name: "Beth",
          }
      ```
    - You use the names of the fields in the struct to specify the values. When you use this style, **you can leave out keys and specify the fields in any order**. 

    - **Any field not specified is set to its zero value**.

- **A field in a struct is accessed with dotted notation**:

    ```go
        bob.name = "Bob" 
    ```

## Anonymous Structs

- You can also declare that **a variable implements a struct type without first giving the struct type a name**. 

- This is called an **anonymous struct**

        ```go
            pet := struct {
                name string
                kind string
            }{
                name: "Fido",
                kind: "dog",
            }
        ```

  - There are two common situations where anonymous structs are
  handy.

      - The first is when you ***translate external data into a struct or a struct into external data (like JSON or protocol buffers). This is called unmarshaling and marshaling data***.
  
      - Writing tests is another place where ***anonymous structs pop up. We’ll use a slice of anonymous structs when writing table-driven tests***

### Comparing and Converting Structs

- Whether or not a struct is comparable depends on the struct’s fields. 
  
  - **Structs that are entirely composed of comparable types are comparable**; 
  
  - **Those with slice or map fields are not (function and channel fields also prevent a struct from being comparable)**.

- Unlike in Python or Ruby, **in Go there’s no magic method that can be overridden to redefine equality and make == and != work for incomparable structs**. 

- We can, of course, **write your own function that you use to compare structs**.

- **Go doesn’t allow comparisons between variables that represent structs of different types**.

- **Go does allow you to perform a type conversion from one struct type to
another if the fields of both structs have the same names, order, and types**

    ```go

        type firstPerson struct {
            name string
            age  int
        }
        type secondPerson struct {
            name string
            age  int
        }

        f:= firstPerson{}
        s:= secondPerson{}
        fmt.Println(f==firstPerson(s)) // true
    ```

- If two struct variables are being compared and at least one of them has a type that’s an anonymous struct, you can compare them without a type conversion if the fields of both structs have the same names, order, and types. 

- You can also assign between named and anonymous struct types if the fields of both structs have the same names, order, and types:

    ```go
    type firstPerson struct {
        name string
        age int
    }

    f := firstPerson{
        name: "Bob",
        age: 50,
    }

    var g struct {
        name string
        age int
    }

    // compiles -- can use = and == between identical named and anonymous structs
    g = f
    fmt.Println(f == g)
    ```
