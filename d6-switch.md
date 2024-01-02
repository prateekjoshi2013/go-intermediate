```go
    words := []string{"a", "cow", "smile", "gopher",
    "octopus", "anthropologist"}

    for _, word := range words {
        switch size := len(word); size {
            case 1, 2, 3, 4:
                fmt.Println(word, "is a short word!")
            case 5:
                wordLen := len(word)
                fmt.Println(word, "is exactly the right length:", wordLen)
            case 6, 7, 8, 9:
            default:
                fmt.Println(word, "is a long word!")
        }
    }
```

- By default, cases in switch statements in Go don’t fall through
  
- An empty case means nothing happens

- Go does include a fallthrough keyword, which lets one case continue on to the next one. 
  
- Please think twice before implementing an algorithm that uses it 

- You ***can switch on any type that can be compared with ==, which includes all of the built-in types except slices, maps, channels, functions, and structs that contain fields of these types***


### Case of missing label
    - The below **break doesnt break out of the loop, it breaks out of the case**.

    ```go
        func main() {
            for i := 0; i < 10; i++ {
                switch {
                    case i%2 == 0:
                        fmt.Println(i, "is even")
                    case i%3 == 0:
                        fmt.Println(i, "is divisible by 3 but not 2")
                    case i%7 == 0:
                        fmt.Println("exit the loop!")
                        break
                    default:
                        fmt.Println(i, "is boring")
                }
            }
        }
    ```

    ```go
        func main() {
        loop:
            for i := 0; i < 10; i++ {
                switch {
                    case i%2 == 0:
                        fmt.Println(i, "is even")
                    case i%3 == 0:
                        fmt.Println(i, "is divisible by 3 but not 2")
                    case i%7 == 0:
                        fmt.Println("exit the loop!")
                        break loop
                    default:
                        fmt.Println(i, "is boring")
                }
            }
        }
    ```

  - **To break out of the loop we have to use a label as above**
  
  - **A label needs to be placed at function level always**

### Blank Switches

- **A regular switch only allows you to check a value for equality. A blank switch allows you to use any boolean comparison for each case**

  ```go
    words := []string{"hi", "salutations", "hello"}
    for _, word := range words {
        switch wordLen := len(word); {
            case wordLen < 5:
                fmt.Println(word, "is a short word!")
            case wordLen > 10:
                fmt.Println(word, "is a long word!")
            default:
                fmt.Println(word, "is exactly the right length.")
        }
    }
  ```


### When to use switch cases

- Favor blank switch statements over if/else chains when you have
multiple **related cases**.

