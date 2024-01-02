- Functions declared inside of functions are special; they are closures.

- Closures really become interesting when they are passed to other functions or returned from a function. 

- They allow you to take the variables within your function and use those values outside of your function

```go
    type Person struct {
        FirstName string
        LastName string
        Age int
    }

    people := []Person{
        {"Pat", "Patterson", 37},
        {"Tracy", "Bobbert", 23},
        {"Fred", "Fredson", 18},
    }

    fmt.Println(people)

    // sort by last name

    sort.Slice(people, func(i int, j int) bool {
        return people[i].LastName < people[j].LastName
    })

    fmt.Println(people)
```

- The closure that’s passed to **sort.Slice** has two parameters, i and j, **but within the closure, we can refer to people** so we can sort it by the LastName field.

### Returning Functions from Functions

- Not only can you use a closure to pass some function state to another function, you can also return a closure from a function. 

- Here is our function that returns a closure:

```go
    func makeMult(base int) func(int) int {
        return func(factor int) int {
            return base * factor
        }
    }

    func main() {
        twoBase := makeMult(2)
    }
```