- **cgo is most useful at the border between Go programs and the outside world**. 

- Reflection helps integrate with external textual data, unsafe is best used with operating system and network data

- **cgo is best for integrating with C libraries**.

- All of the major operating systems are primarily written in either C or C++, which means that they are bundled with libraries written in C. 

- It also means that nearly every programming language provides a **way to integrate with C libraries**. 

- **Go calls its FFI (foreign function interface) to C cgo**.

- A very simple program that calls C code to do some math:

```go
    package main

    import "fmt"

    /*
        #cgo LDFLAGS: -lm
        #include <stdio.h>
        #include <math.h>
        #include "mylib.h"

        int add(int a, int b) {
            int sum = a + b;
            printf("a: %d, b: %d, sum %d\n", a, b, sum);
            return sum;
        }
    */

    import "C"

    func main() {
        sum := C.add(3, 2)
        fmt.Println(sum)
        fmt.Println(C.sqrt(100))
        fmt.Println(C.multiply(10, 20))
    }
```


- The mylib.h header is in the same directory as our main.go, along with mylib.c:

```mylib.h```

```C
    int multiply(int a, int b);
```

```mylib.c```

```c
    #include "mylib.h"

    int multiply(int a, int b) {
        return a * b;
    }
```

```go

```
- Assuming you have a C compiler installed on your computer, all you need to do is compile your program with go build:


```sh
    $ go build

    $ ./example1
    a: 3, b: 2, sum 5
    5
    10
    200
```

- **Since cgo isn’t fast, and it isn’t easy to use for nontrivial programs, the only reason to use cgo is if there is a C library that you must use and there is no suitable Go replacement**

- **Rather than writing cgo yourself, see if there’s already a third-party module that provides the wrapper**.