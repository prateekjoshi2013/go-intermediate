- The **unsafe package allows us to manipulate memory**. 

- The unsafe package is very small and very odd. 

### Unsafe Functions

- It **defines three functions and one type, none of which act like the types and functions found in other packages**.

- The functions are:
  
  -  ```Sizeof```: **which takes in a variable of any type and returns how many bytes it uses **
  
  -  ```Offsetof```: **which takes in a field of a struct and returns the number of bytes from the start of the struct to the start of the field** 
  
  -  ```Alignof```: **which takes in a field or a variable and returns the byte alignment it requires**. 

-  Unlike every other non-built-in function in Go, **these functions can have any value passed into them, and the values returned are constants, so they can be used in constant expressions**.


### Unsafe Pointer Type

- The **unsafe.Pointer type**: 

- **It is a special type that exists for a single purpose: a pointer of any type can be converted to or from unsafe.Pointer**. 

- In addition to pointers, unsafe.Pointer can also be converted to or from a special integer type, called uintptr. 

- Like any other integer type, you can do math with it. 

- This allows you to walk into an instance of a type, extracting individual bytes.

- You can also perform pointer arithmetic, just like you can with pointers in C and C++.

- This byte manipulation changes the value of the variable.

### Usage

- There are **two common patterns in unsafe code**:

  - The first is a **conversion between two types of variables that are normally not convertible** this is **performed using a series of type conversions with unsafe.Pointer in the middle.**

  - The second is **reading or modifying the bytes in a variable by converting a variable to an unsafe.Pointer, converting the unsafe.Pointer to a uintptr, and then copying or manipulating the underlying bytes**.

- Just as we used reflect to translate text data between the outside world and Go code, we use unsafe to translate binary data. 

- There are **two main reasons for using unsafe**:
  
  - The plurality of the **uses of unsafe are for system interoperability. The Go standard library uses unsafe to read data from and write data to the operating system**.
    
  - The second reason that people **use unsafe is for performance, especially when reading data from a network. If you want to map the data into or out of a Go data structure, unsafe.Pointer gives you a very fast way to do so**.
  
  - Example:
  
  - Let’s imagine that we have a wire protocol with the following structure:
  
    - Value: 4 bytes, representing an unsigned, big-endian 32-bit int
    - Label: 10 bytes, ASCII name for the value
    - Active: 1 byte, boolean flag to indicate if the field is active
    - Padding: 1 byte, because we want everything to fit into 16 bytes
  
  - We can define a data structure that matches this:

```go

    type Data struct {
        Value uint32 // 4 bytes
        Label [10]byte // 10 bytes
        Active bool // 1 byte
        // Go padded this with 1 byte to make it align
    }

    /*
        Let’s say we just read the following bytes off the network:

            [0 132 95 237 80 104 111 110 101 0 0 0 0 0 1 0]
    */
```

  - We’re going to read these bytes into an array of length 16 and then convert that array into the struct described previously.
  
  - With safe Go code, we could map it like this:

    ```go
        func DataFromBytes(b [16]byte) Data {
            d := Data{}
            d.Value = binary.BigEndian.Uint32(b[:4])
            copy(d.Label[:], b[4:14])
            d.Active = b[14] != 0
            return d
        }
    ```

    - We could use unsafe.Pointer instead:

    ```go
        func DataFromBytesUnsafe(b [16]byte) Data {
            data := *(*Data)(unsafe.Pointer(&b))
            if isLE {
                data.Value = bits.ReverseBytes32(data.Value)
            }
            return data
        }
    ```

    - First, 
      
      - we take a pointer to our byte array and convert it to an unsafe.Pointer. 
    
      - Then we convert the unsafe.Pointer to a (\*Data) (we have to put (\*Data) in parentheses because of Go’s order of operations). 
      
      - We want to return the struct, not a pointer to it, so we dereference the pointer
    
    - Next, 
    
      - we check our flag to see if we are on a little-endian platform. If so, we reverse the bytes in the Value field.
    
        -How do we know if we are on a little-endian platform? Here’s the code we’re using:

        ```go
            var isLE bool

            func init() {
                var x uint16 = 0xFF00
                xb := *(*[2]byte)(unsafe.Pointer(&x))
                isLE = (xb[0] == 0x00)
            } 
        ```
        
        - **On a little-endian platform, the bytes that represent x will be stored as [00 FF]**. 
        
        - **On a big-endian platform, x is stored in memory as [FF 00]. We use unsafe.Pointer to convert a number to an array of bytes, and we check what the first byte is to determine the value of isLE**.
    
    - Finally, 
      
      - we return the value.
    
    
    - if we wanted to write our Data back to the network, we could use safe Go:

        ```go
            func BytesFromData(d Data) [16]byte {
                out := [16]byte{}
                binary.BigEndian.PutUint32(out[:4], d.Value)
                copy(out[4:14], d.Label[:])
                if d.Active {
                    out[14] = 1
                }
                return out
            }
        ```

    - Or we could use unsafe:

        ```go
            func BytesFromDataUnsafe(d Data) [16]byte {
                if isLE {
                    d.Value = bits.ReverseBytes32(d.Value)
                }
                b := *(*[16]byte)(unsafe.Pointer(&d))
                return b
            }
        ```

    - **On an Intel i7-8700 computer (which is little-endian), using unsafe.Pointer is roughly twice as fast**
    
    - We can also use unsafe to interact with slices and strings. 
    
    - A string is represented in Go as a pointer to a sequence of bytes and a length. There’s a type in the reflect package called reflect.StringHeader that has this structure, and we use it to access and modify the underlying representation of a string:

    ```go
        s := "hello"
        sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
        fmt.Println(sHdr.Len) // prints 5
    ```

    - We can read the bytes in the string using “pointer arithmetic,” using the Data field in sHdr, which is of type uintptr:

    ```go
        for i := 0; i < sHdr.Len; i++ {
        bp := *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
            fmt.Print(string(bp))
        }
        fmt.Println()
        runtime.KeepAlive(s)
    ```
    - The Data field in the reflect.StringHeader is of type uintptr, and as we’ve discussed, you can’t rely on a uintptr to refer to valid memory for longer than a single statement. 
    
    - How do we keep garbage collection from making this reference invalid? We do so by adding a call to runtime.KeepAlive(s) at the end of the function. 
    
    - This tells the Go runtime to not garbage collect s until after the call to KeepAlive.
    
    - Just as you can get a reflect.StringHeader from a string via unsafe, you can get a reflect.SliceHeader from a slice. It has three fields, Len, Cap, and Data, which represent the length, capacity, and data pointer for the slice, respectively:

    ```go
        s := []int{10, 20, 30}
        sHdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
        fmt.Println(sHdr.Len) // prints 3
        fmt.Println(sHdr.Cap) // prints 3
    ```
    
    - Just like we did for strings, we use a type conversion to convert a pointer to our int slice to an unsafe.Pointer. 
    
    - We then convert the unsafe.Pointer to a pointer to a reflect.SliceHeader. 
    
    - We can then access the length and the capacity of the slice via the Len and Cap fields. 
    
    - Next, we can walk through the slice:


    ```go
        intByteSize := unsafe.Sizeof(s[0])
        fmt.Println(intByteSize)
        for i := 0; i < sHdr.Len; i++ {
        intVal := *(*int)(unsafe.Pointer(sHdr.Data + intByteSize*uintptr(i)))
        fmt.Println(intVal)
        }
        runtime.KeepAlive(s)
    ```

    - **Since the size of an int can be either 32 or 64 bits, we must use unsafe.Sizeof to find out how many bytes each value is in the block of memory that’s pointed to by the Data field**. 
    
    - We then **convert i to uintptr, multiply it by the size of an int, add it to the Data field, convert from uintptr to unsafe.Pointer, and then convert to a pointer to an int, and then finally dereference the int to get the value**
  
### unsafe Tools

- Go is a language that values tooling, and there is a compiler flag to help you **find misuse of uintptr and unsafe.Pointer**. 

- Run your code with the flag to add additional checks at runtime.

```sh
    -gcflags=-d=checkptr 
```

