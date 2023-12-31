### Slices

- Data structure that holds a sequence of values, a slice is what you should use. 

- Slices so useful is that the length is not part of the type for a slice. This removes the limitations of arrays. 

- We can write a single function that processes slices of any size, and we can grow slices as needed.

- The first thing to notice is that we don’t specify the size of the slice when we declare it:

    ```go
        var x = []int{10, 20, 30}
    ```
- This creates a slice of 3 ints using a slice literal. Just like arrays, we can also specify only the indices with values in the slice literal:

    ```go
        var x = []int{1, 5: 4, 6, 10: 100, 15}
    ```

- You can simulate multidimensional slices and make a slice of slices:

    ```go
        var x [][]int
    ```

- You read and write slices using bracket syntax, and, just like with arrays, you can’t read or write past the end or use a negative index

    ```go
        x[0] = 10
    ```
- ```var x []int```  creates a slice of ints, Since no value is assigned, x is assigned the zero value for a slice, which is something we haven’t seen before: ```nil```

- ```nil``` has no type, so it can be assigned or compared against values of different types

- A slice is the first type we’ve seen that ***isn’t comparable***. It is a **compile-time error to use == to see if two slices are identical or !=** to see if they are different. 

- **The only thing you can compare a slice with is nil**

### functions on slices

- ```len```
  
  - Builtin ```len``` function when looking at arrays. It works for slices, too, and when you pass a ```nil``` slice to len, it returns 0.

- ```cap```
  
  - Every slice has a capacity, which is the number of consecutive memory locations reserved. This can be larger than the length. 
  
  - Each time you append to a slice, one or more values is added to the end of the slice. Each value added increases the length by one. 
  
  - When the length reaches the capacity, there’s no more room to put values. If you try to add additional values when the length equals the capacity, the append function uses the Go runtime to allocate a new slice with a larger capacity. The values in the original slice are copied to the new slice, the new values are added to the end, and the new slice is returned


- ```append```
  
  - The built-in append function is used to grow slices:

    ```go
        var x []int
        x = append(x, 10)
    ```
  - The append function takes at least two parameters, a slice of any type and a value of that type
  
  - You can append more than one value at a time:
    
    ```go
        x = append(x, 5, 6, 7)
    ```
  - One slice is appended onto another by using the … operator
    
    
    ```go
        y := []int{20, 30, 40}
        x = append(x, y...)
    ```

  -  The append function returns a new slice that may refer to a different underlying array if the original array is not large enough to accommodate the additional elements.
  
  -  This behavior is due to the fact that slices in Go are dynamically sized views into arrays, and when the capacity of the underlying array is exceeded, a new larger array is allocated, and the existing elements are copied over.

- ```make```
  
  - ***Using a slice literal or the nil zero value*** 
  - **```make``` allows you to create an empty slice that already has a length or capacity specified**. 

    ```go
        x := make([]int, 5)
        x = append(x, 10)
    ```
  - Since it has a **length of 5**, x[0] through x[4] are valid elements, and **they are all initialized to 0**.
  
  - We can also specify an initial capacity with make:

    ```go
        x := make([]int, 5, 10)
    ```
  - This creates an int slice with a length of 5 and a capacity of 10.
  
  - Never specify a **capacity that’s less than the length! It is a compile time error** to do so with a constant or numeric literal.
  
  - If you use a **variable to specify a capacity that’s smaller than the length, your program will panic at runtime**.

### How to declare slice

- ***Use make with a zero length and a specified capacity. This allows you to use append to add items to the slice. If the number of items turns out to be smaller, you won’t have an extraneous zero value at the end***

```go
    mySlice := make([]int,0)
	mySlice = append(mySlice,1,2,3,4,5,6,7)
```

- ***If you have specified a slice’s length using make, be sure that you mean to append to it before you do so, or you might end up with a bunch of surprise zero values at the beginning of your slice***.


### Slicing slices

- **A slice expression creates a slice from a slice**. 
    
    ```go
        x := []int{1, 2, 3, 4}
    ```

- It’s written inside brackets and consists of a starting offset and an ending offset, separated by a colon (**:**). 

    ```go
        d := x[1:3]
    ```

- If you leave off the starting offset, 0 is assumed. 

    ```go
        y := x[:2]
    ```

- Likewise, if you leave off the ending offset, the end of the
slice is substituted.

    ```go
        z := x[1:]
    ```

- creating a slice copy

    ```go
        e := x[:]
    ```

- **When you take a slice from a slice, you are not making a copy of the data. Instead, you now have two variables that are sharing memory. This means that changes to an element in a slice affect all slices that share that element**

- **To avoid complicated slice situations, you should either never use append with a subslice or make sure that append doesn’t cause an overwrite by using a full slice expression**.

### Converting Arrays to Slices

- If you have an array, you can take a slice from it using a slice expression. This is a useful way to bridge an array to a function that only takes slices. 

- However, be aware that taking a slice from an array has the
same memory-sharing properties as taking a slice from a slice. If you run the following code on

    ```go
        x := [4]int{5, 6, 7, 8}
        y := x[:2]
        z := x[2:]
    ```

### copy

- If you need to create a slice that’s ***independent of the original***, use the built-in ```copy``` function

    ```go
        x := []int{1, 2, 3, 4}
        y := make([]int, 4)
        num := copy(y, x) // y: [1 2 3 4] num: 4
    ```
- You don’t need to copy an entire slice. The following code copies the first two elements of a four-element slice into a two-element slice:

    ```go
        x := []int{1, 2, 3, 4}
        y := make([]int, 2)
        num = copy(y, x) // y: [1 2] num: 2
    ```

- You could also copy from the middle of the source slice:

    ```go
        x := []int{1, 2, 3, 4}
        y := make([]int, 2)
        num = copy(y, x[2:]) // y: [3 4] num: 2
    ```
- copy function allows you to copy between two slices that cover overlapping sections of an underlying slice:

    ```go
        x := []int{1, 2, 3, 4}
        num = copy(x[:3], x[1:]) // y: [2 3 4 4] num: 3
    ```

- use copy with arrays by taking a slice of the array. You can make the array either the source or the destination of the copy.
  
  ```go
    x := []int{1, 2, 3, 4}
    d := [4]int{5, 6, 7, 8}
    y := make([]int, 2)
    copy(y, d[:])
    fmt.Println(y) // [5 6]
    copy(d[:], x)
    fmt.Println(d) // [1 2 3 4]
  ```