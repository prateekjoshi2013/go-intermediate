- When the type of the data can’t be determined at compile time, you can use the reflection support in the reflect package to interact with and even construct data. 

- When you need to take advantage of the memory layout of data types in Go, you can use the unsafe package. 

- if there is functionality that can only be provided by libraries written in C, you can call into C code with cgo.

- **```Reflection``` allows us to examine types at runtime. It also provides the ability to examine, modify, and create variables, functions, and structs at runtime**.

### Uses of reflection

- If you look at the Go standard library, you can get an idea. **Its uses fall into one of a few general categories**:

- **Reading and writing from a database. The ```database/sql``` package uses reflection to send records to databases and read data back**.

- **Go’s built-in templating libraries, ```text/template``` and ```html/template```, use reflection to process the values that are passed to the templates**.

- **The ```fmt``` package uses reflection heavily, as all of those calls to ```fmt.Println``` and friends rely on reflection to detect the type of the provided parameters**.

- **The ```errors``` package uses reflection to implement ```errors.Is``` and ```errors.As```**.

- **The ```sort package``` uses reflection to implement functions that sort and evaluate slices of any type**: 

  - ```sort.Slice```, 

  - ```sort.SliceStable```,

  - ```sort.SliceIsSorted```.

- **The last main usage of reflection in the Go standard library is for ```marshaling``` and ```unmarshaling``` data into JSON and XML, along with the other data formats defined in the various encoding packages**. 

- **```Struct tags``` are accessed via reflection, and the fields in structs are read and written using reflection as well**.

- Most of these examples have one thing in **common**: 
  
  - They **involve accessing and formatting data that is being imported into or exported out of a Go program**.
  
  - **Reflection is used at the boundaries between your program and the outside world**.

- **Another use of the reflect package** in the Go standard library: 
  
  - **testing**
  
    - **The ```reflect.DeepEqual``` function checks to see if two values are ```deeply equal``` to each other**. 
    
    - **This is a more thorough comparison than what you get if you use == to compare two things, and it’s used in the standard library as a way to validate test results**. 
    
    - It **can also compare things that can’t be compared using ==, like slices and maps**.

### Types, Kinds, and Values


#### Types and Kinds

##### Type

- **A type defines the properties of a variable, what it can hold, and how you can interact with it**. 

- With reflection, **you are able to query a type to find out about these properties using code**.

- **We get the reflection representation of the type of a variable with the TypeOf function in the reflect package**:

```go
    vType := reflect.TypeOf(v)
```

- The **reflect.TypeOf function returns a value of type reflect.Type, which represents the type of the variable passed into the TypeOf function**. 

- The **reflect.Type type defines methods with information about a variable’s type**.

```go
    var x int
    xt := reflect.TypeOf(x)
    fmt.Println(xt.Name()) // returns int
    f := Foo{}
    ft := reflect.TypeOf(f)
    fmt.Println(ft.Name()) // returns Foo
    xpt := reflect.TypeOf(&x)
    fmt.Println(xpt.Name()) // returns an empty string
```

  - For **primitive types like int, Name() returns the name of the type**, in this case the string ```“int”``` for our int. 

  - **For a struct, the ```name of the struct``` is returned**. 

  - **types, like a slice or a pointer, don’t have names; in those cases, Name returns an ```empty string```**

#### Kind

- **The Kind method on reflect.Type returns a value of type reflect.Kind, which is a constant that says what the type is made of—a slice, a map, a pointer, a struct, an interface, a string, an array, a function, an int, or some other primitive type**. 

- The *```difference``` between **the kind and the type** can be tricky to understand. Remember this rule*: 
  
  - **if you define a struct named Foo, the kind is reflect.Struct and the type is “Foo”**.

- **The kind is very important**. One thing to be **aware of when using reflection is that everything in the reflect package assumes that you know what you are doing**. 

- **Some of the methods defined on reflect.Type and other types in the reflect package only make sense for certain kinds**.

- For example, 

  - *there’s a **method** on **reflect.Type** called **NumIn**.* 

  - *If your reflect.Type instance represents a function, it returns the number of input parameters for the function*. 

  - *If your reflect.Type instance isn’t a function, calling NumIn will panic your program*.

- Another **important method on reflect.Type is Elem**. 
  
  - Some types in Go have references to other types and Elem is how to find out what the contained type is. 
  
  - For example, let’s use *reflect.TypeOf on a pointer to an int*:

    ```go
        var x int
        xpt := reflect.TypeOf(&x)
        fmt.Println(xpt.Name()) // returns an empty string
        fmt.Println(xpt.Kind()) // returns reflect.Ptr
        fmt.Println(xpt.Elem().Name()) // returns "int"
        fmt.Println(xpt.Elem().Kind()) // returns reflect.Int
    ```

#### Reflecting on structs

- **There are methods on reflect.Type for reflecting on structs**. 

- **Use the ```NumField``` method to get ```the number of fields``` in the struct, and get the fields in a struct by index with the Field method**. 

- That **returns each field’s structure described in a reflect.StructField, which has the name, order, type, and struct tags on a field**.

```go
    type Foo struct {
        A int `myTag:"value"`
        B string `myTag:"value2"`
    }

    var f Foo

    ft := reflect.TypeOf(f)
    for i := 0; i < ft.NumField(); i++ {
        curField := ft.Field(i)
        fmt.Println(curField.Name, curField.Type.Name(),
        curField.Tag.Get("myTag"))
    }

```

- output
  
```sh
    A int value
    B string value2
```

### Values

- We can also **use reflection to read a variable’s value, set its value, or create a new value from scratch**.
- We **use the ```reflect.ValueOf``` function to create a ```reflect.Value``` instance that represents the value of a variable**:

  ```go
    vValue := reflect.ValueOf(v)
  ```

- Since every variable in Go has a type, **reflect.Value has a method called Type that returns the reflect.Type of the reflect.Value**. 

- There’s also a **Kind method, just as there is on reflect.Type**.
  
#### To **read our values back out of a reflect.Value**. 

  - The **Interface method returns the value of the variable as an empty interface**. 

  - **However, the type information is lost; when you put the value returned by Interface into a variable, you have to use a type assertion to get back to the right type**:

  ```go
    s := []string{"a", "b", "c"}
    sv := reflect.ValueOf(s) // sv is of type reflect.Value
    s2 := sv.Interface().([]string) // s2 is of type []string
  ```

  - While **Interface can be called for reflect.Value instances that contain values ofany kind there are special case methods that you can use if the kind of the variable is one of the ```built-in, primitive types```: Bool, Complex, Int, Uint, Float, and String**.

  - There’s **also a Bytes method that works if the type of the variable is a slice of bytes**.
   
  - **If you use a method that doesn’t match the type of the reflect.Value, your code will panic**.

#### To Set The Values

- It’s a **three-step process**:

- First, you **pass a pointer to the variable into reflect.ValueOf. This returns a reflect.Value that represents the pointer**:

    ```go
      i := 10
      iv := reflect.ValueOf(&i)
    ```

- Second **you need to get to the actual value to set it**. 
  
  - You **use the Elem method on reflect.Value to get to the value pointed to by the pointer that was passed into reflect.ValueOf**. 
  
  - Just **like Elem on reflect.Type returns the type that’s pointed to by a containing type, Elem on reflect.Value returns the value that’s pointed to by a pointer or the value that’s stored in an interface**:
  
    ```go
      ivv := iv.Elem()
    ```
- Finally, **you get to the actual method that’s used to set the value**. 
  
  - Just like there are special-case methods for reading primitive types, there are special-case methods for setting primitive types: 
    
    - SetBool, SetInt, SetFloat, SetString, and SetUint. 
    - For all other types, you need to use the Set method, which takes a variable of type reflect.Value

    ```go
      ivv.SetInt(20)
      fmt.Println(i) // prints 20
    ``` 
#### Making New Values

- The **reflect.New function is the reflection analog of the new function**.

- **It takes in a reflect.Type and returns a reflect.Value that’s a pointer to a reflect.Value of the specified type**. 

- **Since it’s a pointer, you can modify it and then assign the modified value to a variable using the Interface method**

-Just as reflect.New creates a pointer to a scalar type, you can also use reflection to do the same thing as the make keyword with the following functions:

  ```go
    func MakeChan(typ Type, buffer int) Value
    func MakeMap(typ Type) Value
    func MakeMapWithSize(typ Type, n int) Value
    func MakeSlice(typ Type, len, cap int) Value
  ```
- Each of **these functions takes in a reflect.Type that represents the compound type, not the contained type**.

- You must **always start from a value when constructing a reflect.Type**. 

- However, there’s a **trick that lets you create a variable to represent a reflect.Type if you don’t have a value handy**:

  ```go
    var stringType = reflect.TypeOf((*string)(nil)).Elem()
    var stringSliceType = reflect.TypeOf([]string(nil))
  ```

- **Now that we have these types, we can see how to use reflect.New and reflect.MakeSlice**:

  ```go
    ssv := reflect.MakeSlice(stringSliceType, 0, 10)
    sv := reflect.New(stringType).Elem()
    sv.SetString("hello")
    ssv = reflect.Append(ssv, sv)
    ss := ssv.Interface().([]string)
    fmt.Println(ss) // prints [hello]
  ```

### Use Reflection to Check If an Interface’s Value Is nil

- **If a nil variable of a concrete type is assigned to a variable of an interface type, the variable of the interface type is not nil**. 

- **This is because there is a type associated with the interface variable**. 

- **If you want to check if the value associated with an interface is nil, you can do so with reflection using two methods: ```IsValid``` and ```IsNil```**:

```go
  func hasNoValue(i interface{}) bool {
    iv := reflect.ValueOf(i)
    if !iv.IsValid() {
      return true
    }
    switch iv.Kind() {
      case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
        return iv.IsNil()
      default:
        return false
    }
  }
```

