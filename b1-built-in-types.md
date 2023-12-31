- Booleans
  
  - The **bool** type represents Boolean variables. Variables of bool type can have one of two values: true or false. 
  
  - The **zero value for a bool is false**:

    ```go
        var flag bool // no value assigned, set to false
        var isAwesome = true
    ```
- Integer

    - **int8**      –128 to 127
    - **int16**     –32768 to 32767
    - **int32**     –2147483648 to 2147483647
    - **int64**     –9223372036854775808 to 9223372036854775807
    - **byte or uint8**     0 to 255
    - **uint16**    0 to 65536
    - **uint32**    0 to 4294967295
    - **uint64**    0 to 18446744073709551615
    - **int and uint**       On a 32-bit CPU, int is a 32-bit signed/unsigned integer. On most 64-bit CPUs, int is a 64-bit signed/unsigned integer
    
    - #### which integer type to use ?

      - If you are working with a **binary file format or network protocol that has an integer of a specific size or sign, use the corresponding integer type**.
      
      - If you are **writing a library function that should work with any integer type int64 should be used**
      
      - In all other cases, just use **int**. 

  - Integer operators
  
    - Go integers support the usual arithmetic operators: +, -, *, /, with % for modulus. 
  
    - The result of an integer division is an integer; if you want to get a floating point result, you need to use a type conversion **float64(3)**to make your integers into floating point numbers.
    
    - combine any of the arithmetic operators with = to modify a variable: 
      - $+= ,-= , *=, /=, \%=$

    - compare integers with **==, !=, >, >=, <, and <=**.

    - You can bit shift left and right with << and >>, or do bit masks with & (logical AND), | (logical OR), ^ (logical XOR), and &^ (logical AND NOT).

- Floating Point
  
    - **float32** 
    - **float64** is used more often
    
  - **All the standard mathematical and comparison operators with floats, except %**
  
  - **Dividing a nonzero floating point variable by 0 returns +Inf or -Inf**
  
  -  **Dividing a floating point variable set to 0 by 0 returns NaN** 
  
  - A floating point number cannot represent a decimal value exactly.
  
    - **Do not use them to represent money or any other value that must have an exact decimal representation**
    
    -  While Go lets you use == and != to compare floats, don’t do it.
    
    -  Instead define a maximum allowed variance and see if the difference between two floats is less than that

- Complex Numbers 
  - Not used generally
  
  - If you use **untyped constants or literals for both function parameters**, you’ll create **an untyped complex literal, which has a default type of complex128**.

  - If **both of the values passed into complex are of float32 type, you’ll create a complex64**.
  
  - If **one value is a float32 and the other value is an untyped constant or literal that can fit within a float32**, you’ll create a **complex64**.
  
  - Otherwise, you’ll create a **complex128**.

```go
    func main() {
        x := complex(2.5, 3.1)
        y := complex(10.2, 2)
        fmt.Println(x + y)
        fmt.Println(x - y)
        fmt.Println(x * y)
        fmt.Println(x / y)
        fmt.Println(real(x))
        fmt.Println(imag(x))
        fmt.Println(cmplx.Abs(x))
    }
```

  - If you do want to write **numerical computing applications in Go**, you can use the **third-party Gonum package**. 

  - It takes advantage of complex numbers and provides useful libraries for things like linear algebra, matrices, integration, and statistics.
  
- string

  - strings are compared for equality using **==**
  - difference with **!=**
  
  - ordering with **>, >=, <, or <=**
  
  - concatenated by using the **+** operator
  
  - Strings in Go are **immutable**; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it

- Rune
  
  - The rune type is an **alias for the
int32** type, just like **byte is an alias for uint8**. 

