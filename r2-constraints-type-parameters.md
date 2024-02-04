- **type ```constraints``` are actually ```interface``` types**.
### Enhanced interface syntax

- Some new notations are introduced into Go to make it possible to use interface types as constraints.

- The ~T form, where T is a type literal or type name. T must denote a non-interface type whose underlying type is itself (so T may not be a type parameter, which is explained below). 
  
  - The form denotes a type set, which include all types whose underlying type is T. 
  
  - The ~T form is called a tilde form or type tilde in this book (or underlying term and approximation type elsewhere).

- The T1 | T2 | ... | Tn form, which is called a union of terms (or type/term union in this book). 
  
  - Each Tx term must be a tilde form, a type literal, or a type name, and it may not denote a type parameter. 
  
  - There are some requirements for union terms. 

- Some legal examples of the new notations:

```go

    // tilde forms
    ~int
    ~[]byte
    ~map[int]string
    ~chan struct{}
    ~struct{x int}

    // unions of terms
    uint8 | uint16 | uint32 | uint64
    ~[]byte | ~string
    map[int]int | []int | [16]int | any
    chan struct{} | ~struct{x int}

```

- The following code snippet shows some interface type declarations, in which the interface type literals in the declarations of N and O are only legal since Go 1.18.

```go

    type L interface {
        Run() error
        Stop()
    }

    type M interface {
        L
        Step() error
    }

    type N interface {
        M
        interface{ Resume() }
        ~map[int]bool
        ~[]byte | string
    }

    type O interface {
        Pause()
        N
        string
        int64 | ~chan int | any
    }
```

- **Embedding an interface type in another one is equivalent to (recursively) expanding the elements of the former into the latter**. 

- In the above example, **the declarations of M, N and O are equivalent to the following ones**:


```go

    type M interface {
        Run() error
        Stop()
        Step() error
    }

    type N interface {
        Run() error
        Stop()
        Step() error
        Resume()
        ~map[int]bool
        ~[]byte | string
    }

    type O interface {
        Run() error
        Stop()
        Step() error
        Pause()
        Resume()
        ~map[int]bool
        ~[]byte | string
        string
        int64 | ~chan int | any
    }
```

- **We could view a single type literal, type name or tilde form as a term union with only one term**. 

- So simply speaking, **since Go 1.18, an interface type may specify some methods and embed some term unions**.

### Type sets and type implementations

- **Before Go 1.18, an interface type is defined as a method set. Since Go 1.18, an interface type is defined as a type set**. 

- **A type set consists of only non-interface types**.

- In fact, **every type term defines a method set**. 

- **Calculations of type sets follow the following rules**:

- The type set of a non-interface type literal or type name only contains the type denoted by the type literal or type name.

- As just mentioned above, the type set of a tilde form ~T is the set of types whose underlying types are T. In theory, this is an infinite set.

- The type set of a method specification is the set of non-interface types whose method sets include the method specification. In theory, this is an infinite set.

- The type set of an empty interface (such as the predeclared any) is the set of all non-interface types. In theory, this is an infinite set.

- The type set of a union of terms T1 | T2 | ... | Tn is the union of the type sets of the terms.

- The type set of a non-empty interface is the intersection of the type sets of its interface elements.

- By the current specification, two unnamed constraints are equivalent to each other if their type sets are equal.

Given the types declared in the following code snippet, the type set of each interface type is described in the preceding comment of that interface type.

```go
    type Bytes []byte  // underlying type is []byte
    type Letters Bytes // underlying type is []byte
    type Blank struct{}
    type MyString string // underlying type is string

    func (MyString) M() {}
    func (Bytes) M() {}
    func (Blank) M() {}

    // The type set of P only contains one type:
    // []byte.
    type P interface {[]byte}

    // The type set of Q contains
    // []byte, Bytes, and Letters.
    type Q interface {~[]byte}

    // The type set of R contains only two types:
    // []byte and string.
    type R interface {[]byte | string}

    // The type set of S is empty.
    type S interface {R; M()} // intersection

    // The type set of T contains:
    // []byte, Bytes, Letters, string, and MyString.
    type T interface {~[]byte | ~string}

    // The type set of U contains:
    // MyString, Bytes, and Blank.
    type U interface {M()}

    // V <=> P
    type V interface {[]byte; any} // intersection

    // The type set of W contains:
    // Bytes and MyString.
    type W interface {T; U}

    // Z <=> any. Z is a blank interface. Its
    // type set contains all non-interface types.
    type Z interface {~[]byte | ~string | any}
```

### basic and non-basic type

- **Interface types whose type sets can be defined entirely by a method set (may be empty) are called basic interface types**. 

- Before 1.18, Go only supports basic interface types. 

- **Basic interfaces may be used as either value types or type constraints, but non-basic interfaces may only be used as type constraints (as of Go 1.21)**.

- **Take the types declared above as an example,, L, M, U, Z and any are basic types**.

- In the following code, the declaration lines for x and y both compile okay, but the line declaring z fails to compile.

```go
    var x any
    var y interface {M()}
    // error: interface contains type constraints
    var z interface {~[]byte}
```


### predeclared comparable constraint

- **The comparable interface type could be embedded in other interface types to filter out types which are not strictly comparable from their type sets**. 

- For example, 
  
  - **the type set of the following declared constraint C contains only one type: ```string```**, 
  
  - **because the other types in the union are either incomprarable (the first three) or not strictly comparable (the last one)**.


```go
    // resulting type set: string
    type C interface {
        comparable
        []byte | func() | map[int]bool | string | [2]any
    }
```