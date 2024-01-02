- Each place where a **declaration occurs** is called a **block**.

- **Variables, constants, types, and functions declared outside of any functions** are placed in the **package block**.

- **Import statements** in our programs **define names for other packages that are valid for the file that contains the import statement**. These names are in the **file block**.

- **All of the variables defined at the top level of a function (including the parameters to a function) are in a block**

- **Within a function, every set of braces ({}) defines another block**

- **Control structures** in Go define blocks of their own

- You can **access an identifier defined in any outer block from within any inner block**.

### Shadowing

- **When we have a declaration with the same name as an identifier in a containing block we shadow the identifier created in the outer block**.

    ```go
        func main() {
            x := 10
            if x > 5 {
                fmt.Println(x) // 10
                x := 5
                fmt.Println(x) // 5
            }
            fmt.Println(x) // 10
        }
    ```

- **Avoid using := because it can make it unclear what variables are being used **

- That’s **because it is very easy to accidentally shadow a variable when using :=**. 
  
  - Remember, **we can use := to create and assign to multiple variables at once. Also, not all of the variables on the lefthand side have to be new for := to be legal**. 

- **You can use := as long as there is at least one new variable on the lefthand side**


        ```go
            func main() {
                x := 10
                if x > 5 {
                    x, y := 5, 20
                    fmt.Println(x, y)
                }
                fmt.Println(x)
            }
        ```
    - There was an existing definition of x in an outer block, x was still shadowed within the if statement. That’s because := only reuses variables that are declared in the current block.
      
    - When using :=, make sure that you don’t have any variables from an outer scope on the lefthand side, unless you intend to shadow them