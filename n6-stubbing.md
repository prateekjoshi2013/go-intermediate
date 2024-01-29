- Most code is filled with dependencies. 

- **There are two ways that Go allows us to abstract function calls**: 

  - **defining a function type** 

  - **defining an interface**. 

- These **abstractions not only help us write modular production code; they also help us write unit tests**.

- **When code depends on abstractions, it’s easier to write unit tests**!

- **Example of stubbing**

```go
    type Processor struct {
        Solver MathSolver
    }

    type MathSolver interface {
        Resolve(ctx context.Context, expression string) (float64, error)
    }

    func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
        curExpression, err := readToNewLine(r)
        if err != nil {
            return 0, err
        }
        if len(curExpression) == 0 {
            return 0, errors.New("no new expressions to read")
        }

        answer, err := p.Solver.Resolve(ctx, curExpression)
        return answer, err
    }

    type MathSolverStub struct{}

    func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
        switch expr {
        case "2 + 2 * 10":
            return 22, nil
        case "( 2 + 2 ) * 10":
            return 40, nil
        case "( 2 + 2 * 10":
            return 0, errors.New("invalid expression: ( 2 + 2 * 10")
        }
        return 0, nil
    }

    func TestProcessorProcessExpression(t *testing.T) {
        p := Processor{MathSolverStub{}}
        in := strings.NewReader(
            `2 + 2 * 10
        ( 2 + 2 ) * 10
        ( 2 + 2 * 10`)
        data := []float64{22, 40, 0}
        hasErr := []bool{false, false, true}
        for i, d := range data {
            result, err := p.ProcessExpression(context.Background(), in)
            if err != nil && !hasErr[i] {
                t.Error(err)
            }
            if result != d {
                t.Errorf("Expected result %f, got %f", d, result)
            }
        }
    }
```

### Stubbing big go interface

- While most Go interfaces only specify one or two methods, this isn’t always the case. 
- You sometimes find yourself with an interface that has many methods. 
  
  - Assume you have an interface that looks like this:

    ```go
        type EntitiesStub struct {
            getUser func(id string) (User, error)
            getPets func(userID string) ([]Pet, error)
            getChildren func(userID string) ([]Person, error)
            getFriends func(userID string) ([]Person, error)
            saveUser func(user User) error
        }
    ```

  - There are two different approaches:
  
  - The first is to **embed the interface in a struct**. 
  
  - **Embedding an interface in a struct automatically defines all of the interface’s methods on the struct**. 
  - **It doesn’t provide any implementations of those methods, so you need to implement the methods that you care about for the current test**. 
  
  - Let’s assume that Logic is a struct that has a field of type Entities:
    
    
    ```go
        type Logic struct {
            Entities Entities
        }
    ```
  
  - Assume you want to test this method:

```go
    func (l Logic) GetPetNames(userId string) ([]string, error) {
        pets, err := l.Entities.GetPets(userId)
        if err != nil {
            return nil, err
        }
        out := make([]string, len(pets))
        for _, p := range pets {
            out = append(out, p.Name)
        }
        return out, nil
    }
```
- This method uses only one of the methods declared on Entities, GetPets.  

- Rather than creating a stub that implements every single method on Entities just to test GetPets, you can write a stub struct that only implements the method you need to test this method:

```go
    type GetPetNamesStub struct {
        Entities
    }


    func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
        switch userID {
            case "1":
                return []Pet{{Name: "Bubbles"}}, nil
            case "2":
                return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
            default:
                return nil, fmt.Errorf("invalid id: %s", userID)
        }
    }
```
- We then write our unit test, with our stub injected into Logic:

```go
    func TestLogicGetPetNames(t *testing.T) {
        data := []struct {
            name string
            userID string
            petNames []string
        }{
            {"case1", "1", []string{"Bubbles"}},
            {"case2", "2", []string{"Stampy", "Snowball II"}},
            {"case3", "3", nil},
        }
        l := Logic{GetPetNamesStub{}}
        for _, d := range data {
            t.Run(d.name, func(t *testing.T) {
            petNames, err := l.GetPetNames(d.userID)
                if err != nil {
                    t.Error(err)
                }
                if diff := cmp.Diff(d.petNames, petNames); diff != "" {
                    t.Error(diff)
                }
            })
        }
    }
```

- **If you need to implement only one or two methods in an interface for a single test, this technique works well**. 
- **The drawback comes when you need to call the same method in different tests with different inputs and outputs**
- **Solution is to create a stub struct that proxies method calls to function fields**. 
- **For each method defined on Entities, we define a function field with a matching signature on our stub struct**
  
```go
    type EntitiesStub struct {
        getUser func(id string) (User, error)
        getPets func(userID string) ([]Pet, error)
        getChildren func(userID string) ([]Person, error)
        getFriends func(userID string) ([]Person, error)
        saveUser func(user User) error
    }
```

- We then make EntitiesStub meet the Entities interface by defining the methods.
- In each method, we invoke the associated function field. For example:

```go
    func (es EntitiesStub) GetUser(id string) (User, error) {
        return es.getUser(id)
    }

    func (es EntitiesStub) GetPets(userID string) ([]Pet, error) {
        return es.getPets(userID)
    }
```

- Once you create this stub, you can supply different implementations of different methods in different test cases via the fields in the data struct for a table test:

```go
    func TestLogicGetPetNames(t *testing.T) {
        data := []struct {
            name string
            getPets func(userID string) ([]Pet, error)
            userID string
            petNames []string
            errMsg string
        }{
            {"case1", func(userID string) ([]Pet, error) {
                return []Pet{{Name: "Bubbles"}}, nil
            }, "1", []string{"Bubbles"}, ""},
            {"case2", func(userID string) ([]Pet, error) {
                return nil, errors.New("invalid id: 3")
            }, "3", nil, "invalid id: 3"},
        }
        l := Logic{}
        for _, d := range data {
            t.Run(d.name, func(t *testing.T) {
                l.Entities = EntitiesStub{getPets: d.getPets}
                petNames, err := l.GetPetNames(d.userID)
                if diff := cmp.Diff(petNames, d.petNames); diff != "" {
                    t.Error(diff)
                }
                var errMsg string
                if err != nil {
                    errMsg = err.Error()
                }
                if errMsg != d.errMsg {
                    t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
                }
            })
        }
    }
```

### Mock

- **A stub returns a canned value for a given input, whereas a mock validates that a set of calls happen in the expected order with the expected inputs**.

- The two most popular options are the ```gomock library``` from Google and the ```testify```
library from Stretchr, Inc.