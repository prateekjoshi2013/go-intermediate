- Various ways to declare variables

```go
    var x int = 10
    var x = 10
    
    var x int //zero value assigned
    x = 2 //reassigned

    var x, y int = 10, 20

    var x, y int

    var x, y = 10, "hello"

    var (
        x int
        y = 20
        z int = 30
        d, e = 40, "hello"
        f, g string
    )

```

- $:=$ operator to replace a var declaration that uses type inference.

  - The := operator can do one trick that you cannot do with var: 
  
    - **it allows you to assign values to existing variables, too. As long as there is one new variable on the lefthand side of the :=, then any of the other variables can already exist**:

        ```go
            x := 10
            x, y := 30, "hello"
        ```

### What type Declaration to use

- When **initializing a variable to its zero value, use var x int**. This makes it clear
that the zero value is intended.

- ***When assigning an untyped constant*** or ***a literal to a variable and the default type for the constant or literal isn’t the type you want for the variable***, use the long var form with the type specified **var x byte = 20**.

- Because **:=** allows you to assign to both new and existing variables, ***it sometimes creates new variables when you think you are reusing existing ones**. In those situations, **explicitly declare all of your new variables with var to make it clear which variables are new, and then use the assignment operator (=) to assign values to both new and old variables**.

- While **var and := allow you to declare multiple variables on the same line**, only use this style when assigning **multiple values returned from a function** or **the comma ok idiom**
  
- You should **rarely declare variables outside of functions, in what’s called the package block**

- **Package-level variables whose values change are a bad idea**. When you have a variable outside of a function, it can be difficult to track the changes made to it, which makes it hard to understand how data is flowing through your program.