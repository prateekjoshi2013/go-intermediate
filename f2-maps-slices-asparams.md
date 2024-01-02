### Map as a parameter for a function

- A map is implemented as a pointer to a struct. Passing a map to a function means that you are copying
a pointer. 

- Because of this, you should avoid using maps for input parameters or return values, especially on public APIs. 

- On an API-design level, maps are a bad choice because they say nothing about what values are contained within; there’s nothing that explicitly defines what keys are in the map, so the only way to know what they are is to trace through the code.

- Go is a strongly typed language; rather than passing a map around, use a struct


### Slice as a parameter for a function

- passing a slice to a function has more complicated behavior: 
  
  - Any modification to the contents of the slice is reflected in the original variable, but using append to change the length isn’t reflected in the original variable, even if the slice has a capacity greater than its length. 

  - That’s because a slice is implemented as a struct with three fields: an int field for length, an int field for capacity, and a pointer to a block of memory
  
  - If we append to a slice from the parameter and if the slice is expanded beyond capacity a new underlying array is created with a new memory blocak and gets disassociated with the original passed slice

- By default, **you should assume that a slice length and capacity is not modified by a function. Your function’s documentation should specify if it modifies the slice’s contents**.

- **We can’t change the length or capacity of a slice when we pass it to a function, but we can change the contents up to the current length**.

- Since we cannot change the length and capacity of the passes slice we use a fixed slice as a buffer and overwrite contents in it below.

    ```go
        file, err := os.Open(fileName)
        if err != nil {
            return err
        }
        defer file.Close()
        data := make([]byte, 100)
        for {
            count, err := file.Read(data)
            if err != nil {
                return err
            }
            if count == 0 {
                return nil
            }
            process(data[:count])
        }
    ```
### Garbage Collector

- Go encourages you to use pointers sparingly. We reduce the workload of the garbage collector by making sure that as much as possible is stored on the stack. 

- Slices of structs or primitive types have their data lined up sequentially in memory for rapid access. 

- When the garbage collector does do work, it is optimized to return quickly rather than gather the most garbage. The key to making this approach work is to simply create less garbage in the first place