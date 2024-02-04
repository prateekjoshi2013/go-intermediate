- In the custom generics world, **a type may be declared as a generic type, and a function may be declared as generic function**. 

- In addition, **generic types are defined types, so they may have methods**.

- **The declaration of a generic type, generic function, or method of a generic type contains a type parameter list part, which is the main difference from the declaration of an ordinary type, function, or method**.

### A generic type example

```go
    package main

    import "sync"

    /*
        - Lockable is a generic type. 
        
        - There is a type parameter list, in the declaration of a generic type. 
    
        - Here, the type parameter list of the Lockable generic type is [T any].
    */

    type Lockable[T any] struct {
        sync.Mutex
        Data T
    }

    func main() {
        var n Lockable[uint32]
        n.Lock()
        n.Data++
        n.Unlock()
        
        var f Lockable[float64]
        f.Lock()
        f.Data += 1.23
        f.Unlock()
        
        var b Lockable[bool]
        b.Lock()
        b.Data = !b.Data
        b.Unlock()
        
        var bs Lockable[[]byte]
        bs.Lock()
        bs.Data = append(bs.Data, "Go"...)
        bs.Unlock()
    }
```

- **A type parameter list may contain one and more type parameter declarations which are enclosed in square brackets and separated by commas**. 

- **Each parameter declaration is composed of a type parameter name and a (type) constraint**. 
  
  - For the above example, **T is the type parameter name and any is the constraint of T**.

- Note that **```any``` is a new predeclared identifier introduced in Go 1.18**. 
  
  - **It is an ```alias``` of the blank interface type ```interface{}```**. 
  
  - **All types implements the blank interface type**.

- We could view **constraints as types of (type parameter) types**. **All type constraints are actually interface types**.

- A generic type is a defined type. It must be instantiated to be used as value types. 
  
  - The notation Lockable[uint32] is called an instantiated type (of the generic type Lockable). 
  
  - In the notation, [uint32] is called a type argument list and uint32 is called a type argument, which is passed to the corresponding T type parameter.

- **A type argument must implement the constraint of its corresponding type parameter**. 

  - **The constraint any is the loosest constraint, any value type could be passed to the T type parameter**. 
  
  - **The other type arguments used in the above example are: float64, bool and []byte**.

- Every instantiated type is a named type and an ordinary type. For example, Lockable[uint32] and Lockable[[]byte] are both named types.

### An example of a method of a generic type

- Comparing with the Lockable implementation in the last section, the new one hides the struct fields from outside package users of the generic type.

```go
    package main

    import "sync"

    type Lockable[T any] struct {
        mu sync.Mutex
        data T
    }

    func (l *Lockable[T]) Do(f func(*T)) {
        l.mu.Lock()
        defer l.mu.Unlock()
        f(&l.data)
    }

    func main() {
        var n Lockable[uint32]
        n.Do(func(v *uint32) {
            *v++
        })
        
        var f Lockable[float64]
        f.Do(func(v *float64) {
            *v += 1.23
        })
        
        var b Lockable[bool]
        b.Do(func(v *bool) {
            *v = !*v
        })
        
        var bs Lockable[[]byte]
        bs.Do(func(v *[]byte) {
            *v = append(*v, "Go"...)
        })
    }
```

- A **method Do is declared for the generic base type Lockable**. 

  - Here, **the receiver type is a pointer type, which base type is the generic type Lockable**.

  - Different from method declarations for ordinary base types, there is a type parameter list part following the receiver generic type name Lockable in the receiver part. 

    - Here, the type parameter list is [T].

- Please note that, **the type parameter names used in a method declaration for a generic base type are not required to be the same as the corresponding ones used in the generic type specification**.

- **Though, it is a bad practice not to keep the corresponding type parameter names the same**. 

### A generic function example

- Now, let's view an example of how to declare and use generic (non-method) functions.


```go
    package main

    // NoDiff checks whether or not a collection
    // of values are all identical.
    func NoDiff[V comparable](vs ...V) bool {
        if len(vs) == 0 {
            return true
        }
        
        v := vs[0]
        for _, x := range vs[1:] {
            if v != x {
                return false
            }
        }
        return true
    }

    func main() {
        var NoDiffString = NoDiff[string]
        println(NoDiff("Go", "Go", "Go")) // true
        println(NoDiffString("Go", "go")) // false
        
        println(NoDiff(123, 123, 123, 123)) // true
        println(NoDiff[int](123, 123, 789)) // false
        
        type A = [2]int
        println(NoDiff(A{}, A{}, A{}))     // true
        println(NoDiff(A{}, A{}, A{1, 2})) // false
        
        println(NoDiff(new(int)))           // true
        println(NoDiff(new(int), new(int))) // false

        println(NoDiff[bool]())   // true

        // _ = NoDiff() // error: cannot infer V
        
        // error: *** does not implement comparable
        // _ = NoDiff([]int{}, []int{})
        // _ = NoDiff(map[string]int{})
        // _ = NoDiff(any(1), any(1))
    }
```

- In the above example, **NoDiff is a generic function**. 

- Different from non-generic functions, and similar to generic types, there is an extra part, a type parameter list, in the declaration of a generic function. 

- Here, **the type parameter list of the NoDiff generic function is [V comparable], in which V is the type parameter name and comparable is the constraint of V**.

- **```comparable``` is new ```predeclared identifier``` introduced in Go 1.18**

- **The whole type argument list may be totally omitted from an instantiated function expression if all the type arguments could be inferred from the value arguments**. 

  - **That is why some calls to the NoDiff generic function have no type argument lists in the above example**.

- Note that **all of these type arguments implement the comparable interface.**

- **Incomparable types, such as []int and map[string]int may not be passed as type arguments of calls to the NoDiff generic function**. 

- And please note that, **although any is a comparable (value) type, it doesn't implement comparable, so it is also not an eligible type argument**. 
