### Methods

- The methods for a type are defined at the package block level:

```go
    type Person struct {
        FirstName string
        LastName string
        Age int
    }
    
    func (p Person) String() string {
        return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
    }
```

- Method declarations look just like function declarations, with one addition: **the receiver specification**. 

- The **receiver appears between the keyword func and the name of the method**. 

- Just like all other variable declarations, the **receiver name appears before the type**. 

- By convention, **the receiver name is a short abbreviation of the type’s name, usually its first letter**. 

- **It is nonidiomatic to use this or self**.

- **Just like functions, method names cannot be overloaded**. 

- You **can use the same method names for different types, but you can’t use the same method name for two different methods on the same type**.

- **Methods must be declared in the same package as their associated type; Go doesn’t allow you to add methods to types you don’t control**. 
  
- **While you can define a method in a different file within the same package as the type declaration, it is best to keep your type definition and its associated methods together so that it’s easy to follow the implementation**

- 