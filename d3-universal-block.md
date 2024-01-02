## Universal Block

- There’s actually one more block that is a little weird: **the universe block**.

- Go is a small language with only **25 keywords**.

- **Rather than make them keywords, Go considers these predeclared identifiers and defines them in the universe block, which is the block that contains all other blocks**

- **Because these names are declared in the universe block, it means that they can be shadowed in other scopes**.

    ```go
        fmt.Println(true) // true
        true := 10  // shadowing true
        fmt.Println(true) // 10
    ```

- **Must be very careful to never redefine any of the identifiers in the universe block**
  
- **Not even shadow detects shadowing of universe block identifiers**