- For a program to be useful, it needs to read in and write out data. 

- The heart of *Go’s input/output philosophy can be found in the io package*. 

- In particular, **two interfaces defined in this package are probably the second and third most-used interfaces in Go** after *error*:

  - **io.Reader** 
  - **io.Writer**

- **Both io.Reader and io.Writer define a single method**:


```go
    type Reader interface {
        Read(p []byte) (n int, err error)
    }

    type Writer interface {
        Write(p []byte) (n int, err error)
    }
```

- **The Write method on the io.Writer interface takes in a slice of bytes, which are written to the interface implementation**. 
  
  - **It returns the number of bytes written and an error if something went wrong**. 

- **The Read method on io.Reader is more interesting**.
  
  - **Rather than return data through a return parameter, a slice input parameter is passed into the implementation and modified**. 
  
  - **Up to len(p) bytes will be written into the slice**.


- Good reason why io.Reader is defined the way it is. Let’s write a function that’s representative of how to work with an io.Reader to understand:

```go
func countLetters(r io.Reader) (map[string]int, error) {
    buf := make([]byte, 2048)
    out := map[string]int{}
    for {
        n, err := r.Read(buf)
        for _, b := range buf[:n] {
            if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
                out[string(b)]++
            }
        }
        if err == io.EOF {
            return out, nil
        }
        if err != nil {
            return nil, err
        }
    }
}
```

- There are three things to note. 
  
  - First, **we create our buffer once and reuse it on every call to r.Read**. 
    
    - *This allows us to use a single memory allocation to read from a potentially large data source*. 
  
    - *If the Read method were written to return a []byte, it would require a new allocation on every single call. Each allocation would end up on the heap, which would make quite a lot of work for the garbage collector*.
  
  - Second, **we use the n value returned from r.Read to know how many bytes were written to the buffer and iterate over a subslice of our buf slice, processing the data that was read**.

  - Finally, **we know that we’re done reading from r when the error returned from r.Read is io.EOF**. 
    
    - This error is a bit odd, in that it isn’t really an error. It indicates that there’s nothing left to read from the io.Reader. When **io.EOF is returned, we are finished processing and return our result**.
    
    - **If you get to the end of an io.Reader unexpectedly, a different sentinel error is returned (io.ErrUnexpectedEOF)**.


  #### Other Readers and Writers


- Because we have standard interfaces for reading and writing, there’s a standard function in the io package for copying from an io.Reader to an io.Writer, io.Copy.

- There are other standard functions for adding new functionality to existing io.Reader and io.Writer instances. These include:

  - io.MultiReader
  
    - This returns an io.Reader that reads from multiple io.Reader instances, one after another.
  
  - io.LimitReader

    - This returns an io.Reader that only reads up to a specified number of bytes from the supplied io.Reader.
  
  -  io.MultiWriter

     - This returns an io.Writer that writes to multiple io.Writer instances at the same time.

#### Seeker interface

- In Go, the io.Seeker interface is part of the io package, and it is defined as follows:

    ```go
        type Seeker interface {
            Seek(offset int64, whence int) (int64, error)
        }
    ```

- The **io.Seeker interface is implemented by types that can set the position for the next Read or Write operation**. 

  - It defines a *single method*:

    - **Seek(offset int64, whence int) (int64, error)**:
      
      - **Seek sets the offset for the next Read or Write operation to the given offset, interpreted according to the value of whence**.

      - **The offset is the number of bytes to move the cursor, and whence specifies the reference point for the offset**. 
      
      - The values for whence are **io.SeekStart, io.SeekCurrent, and io.SeekEnd**.

```go
    package main

    import (
        "fmt"
        "io"
        "strings"
    )

    func main() {
        reader := strings.NewReader("Hello, Seeker!")

        // Create a Seeker from the reader
        seeker, ok := reader.(io.Seeker)
        if !ok {
            fmt.Println("The underlying type does not implement Seeker")
            return
        }

        // Seek to a specific offset from the start
        offset, err := seeker.Seek(7, io.SeekStart)
        if err != nil {
            fmt.Println("Error during seek:", err)
            return
        }

        fmt.Printf("Seeked to offset %d\n", offset)

        // Read the rest of the string from the current position
        buffer := make([]byte, 100)
        n, err := seeker.Read(buffer)
        if err != nil {
            fmt.Println("Error during read:", err)
            return
        }

        fmt.Printf("Read %d bytes: %s\n", n, buffer[:n])
    }
```

#### io interface combinations

- The *io package defines interfaces that combine these four interfaces in various ways.*

- They include **io.ReadCloser, io.ReadSeeker, io.ReadWriteCloser, io.ReadWriteSeeker, io.ReadWriter, io.WriteCloser, and io.WriteSeeker**. 
  
- **Use these interfaces to specify what your functions expect to do with the data**. 
  
  - For example, **rather than just using an os.File as a parameter, use the interfaces to specify exactly what your function will do with the parameter**. 
  
    - Not only does **it make your functions more general purpose, it also makes your intent clearer**.
    
    - Furthermore, **make your code compatible with these interfaces if you are writing your own data sources and sinks**. 
    
    - In general, **strive to create interfaces as simple and decoupled as the interfaces defined in io. They demonstrate the power of simple abstractions**.