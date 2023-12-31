### Arrays

- ***Arrays are rarely used directly in Go***

- All of the elements in the array **must be of the type that’s specified (this doesn’t mean they are always of the same type)**

- There are a few different declaration styles.
  
  - Specify the size of the array and the type of the elements in the array:
  
    ```go
        var x [3]int
    ```
  - Since no values were specified, all of the positions are ***initialized to the zero*** value for an int
  
  - If you have initial values for the array, you specify them with an array literal:
  
    ```go
        var x = [3]int{10, 20, 30}
    ```
  
  - If you have a sparse array (an array where most elements are set to their zero value), you can specify only the indices with values in the array literal:
  
    ```go
        var x = [12]int{1, 5: 4, 6, 10: 100, 15}
        // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
    ```
  
  - When using an array literal to initialize an array, you can leave off the number and use … instead:

  ```go
    var x = [...]int{10, 20, 30}
  ```  

  - You can **use == and != to compare arrays**:
    
    ```go
        var x = [...]int{1, 2, 3}
        var y = [3]int{1, 2, 3}
        fmt.Println(x == y) // prints true
    ```

- Go only has one-dimensional arrays, but you can simulate multidimensional arrays:
  
  - This declares x to be an array of length 2 whose type is an array of ints of length 3.
    
    ```go
        var x [2][3]int
    ```
    
- Go are read and written using bracket syntax:

    ```go
        x[0] = 10
        fmt.Println(x[2])
    ```

- built-in function ***len*** takes in an array and returns its length:

    ```go
        len(x)
    ```

- Go considers **the size of the array to be part of the type of
the array**. This makes an array that’s declared to be **[3]int a different type from an array that’s declared to be [4]int**. 

- This also means that **you cannot use a variable to specify the size of an array, because types must be resolved at compile time, not at runtime**

- you **can’t write a function that works with arrays of any size**

- you **can’t assign arrays of different sizes to the same variable**

- ***Don’t use arrays unless you know the exact length you need
ahead of time***. 
