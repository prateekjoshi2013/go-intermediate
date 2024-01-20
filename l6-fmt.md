
- the **fmt package provides the fmt.Stringer interface and the fmt package functions that use it to format values as strings**. 

- The **fmt.Stringer interface is designed for types that want to control how they are formatted when printed using functions like ```fmt.Printf```, ```fmt.Println```, and ```others```**.

```go
    type Stringer interface {
        String() string
    }
```

- **The String() method is expected to return a human-readable string representation of the object**. 

- W**hen a value of a type that implements fmt.Stringer is printed using fmt package functions, the String() method is automatically called to obtain the string representation**.


```go
    package main

    import "fmt"

    // Person represents an individual with a name and age.
    type Person struct {
        Name string
        Age  int
    }

    // Implementing fmt.Stringer for the Person type
    func (p Person) String() string {
        return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
    }

    func main() {
        person := Person{Name: "Alice", Age: 25}

        // When printing person, the String() method is automatically called
        fmt.Println(person)
    }
```
