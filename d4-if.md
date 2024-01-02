### if

- The most visible difference between if statements in Go and other languages is that **you don’t put parenthesis around the condition**.

- **Variables declared in if condition are scoped to the condition and to both the if and else blocks**.

    ```go
        if n := rand.Intn(10); n == 0 {
            fmt.Println("That's too low")
        } else if n > 5 {
            fmt.Println("That's too big:", n)
        } else {
            fmt.Println("That's a good number:", n) // That's a good number: 6
        }
        fmt.Println(n) // undefined: n
    ```