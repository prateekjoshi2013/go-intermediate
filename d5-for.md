### for

- for is the only looping keyword in the language.

- Go accomplishes this by using the for keyword in four different formats:
  
  - A complete, C-style for

    ```go
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    ```
    - ***you must use := to initialize the variables; var is not legal here***
  
  
  - A condition-only for
  
    ```go
        := 1
        for i < 100 {
            fmt.Println(i)
            i = i * 2
        }
    ```
    - like the ***while statement***
  
  - An infinite for
    
    ```go
        for {
            fmt.Println("Hello")
        }    
    ```

    - **break and continue statements** work as  in other languages
    
    ```java
        do {
            // things to do in the loop
        } while (CONDITION);
    ```
    - The Go version looks like this:

    ```go
    for {
        // things to do in the loop
        if !CONDITION {
            break
        }
    }
    ```





  - for-range
    
    - **for-range loop and resembles the iterators** found in other languages


    - for-range loop over an array, slice, or string give two loop variables. 
  
      - The first variable is the position in the data structure being iterated
  
      - the second is the value at that position. 
      
      - When looping over an array, slice, or string, an i for index is commonly used. 
  
        ```go
            evenVals := []int{2, 4, 6, 8, 10, 12}
            for i, v := range evenVals {
                fmt.Println(i, v) // 0 2
            }
        ```
    
    - for-range loop over a map  give two loop variables. 
      
      - When iterating through a map, k (for key) is used instead.
      
      - The second variable is frequently called v for value
      
      - **The order of the keys and values varies; some runs may be identical**. 

        ```go
            uniqueNames := map[string]bool{
                "Fred": true, "Raul": true, "Wilma": true
            }
            
            for k := range uniqueNames {
                fmt.Println(k) // fred
            }

            for k,v := range uniqueNames {
                fmt.Println(k,v) // fred true 
            }
        ```
    - Iterating over strings

        ```go
            samples := []string{"hello", "apple_π!"}
            
            for _, sample := range samples {
                for i, r := range sample {
                    fmt.Println(i, r, string(r))
                }
                fmt.Println()
            }
        ```

      - On iterating over a string with a for-range loop. It iterates over the runes, not the bytes.
      
      - Each time the for-range loop iterates over your compound type, it copies the value from the compound type to the value variable. 
      
      - ***Modifying the value variable will not modify the value in the compound type***.
      
        ```go
            evenVals := []int{2, 4, 6, 8, 10, 12}
            for _, v := range evenVals {
                v *= 2
            }
            fmt.Println(evenVals) // [2 4 6 8 10 12]
        ```

      - you’ll see that if you launch goroutines in a for-range loop, you need to be very careful in how you pass the index and value to the goroutines, or you’ll get surprisingly wrong results.


    
