- **Go doesn’t have inheritance, it encourages code reuse via built-in support for composition and promotion**

```go
    type Employee struct {
        Name string
        ID string
    }
    
    func (e Employee) Description() string {
        return fmt.Sprintf("%s (%s)", e.Name, e.ID)
    }
    
    type Manager struct {
        Employee
        Reports []Employee
    }

    func (m Manager) FindNewEmployees() []Employee {
        // do business logic
    }

    m := Manager{
        Employee: Employee{
            Name: "Bob Bobson",
            ID: "12345",
        },
        Reports: []Employee{},
    }

    fmt.Println(m.ID) // prints 12345
    
    fmt.Println(m.Description()) // prints Bob Bobson (12345)
```

- **Any type can be embedded within a struct, not just another struct**.

- **This promotes the methods on the embedded type to the containing struct**.

- If the containing struct has fields or methods with the same name as an embedded field, you need to use the embedded field’s type to refer to the obscured fields or methods. 

  - If you have types defined like this:

  ```go
    type Inner struct {
      X int
    }
    
    type Outer struct {
        Inner
        X int
    }
    
    // You can only access the X on Inner by specifying Inner explicitly:

    o := Outer{
            Inner: Inner{
                X: 10,
            },
            X: 20,
    }
    fmt.Println(o.X) // prints 20
    fmt.Println(o.Inner.X) // prints 10
  ```

### Embedding Is Not Inheritance

- You cannot assign a variable of type Manager to a variable of type Employee
  
- If you want to access the Employee field in Manager, you must do so explicitly.

```go
    var eFail Employee = m // compilation error!
    var eOK Employee = m.Employee // ok!
    // You’ll get the error:
    // cannot use m (type Manager) as type Employee in assignment
```

- **There is no dynamic dispatch for concrete types in Go**. 

  - **The methods on the embedded field have no idea they are embedded**. 

  - **If you have a method on an embedded field that calls another method on the embedded field, and the containing struct has a method of the same name, the method on the embedded field will not invoke the method on the containing struct**
