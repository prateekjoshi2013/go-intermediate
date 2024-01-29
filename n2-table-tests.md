### Table Tests

- Rather than writing usecases over and over, a pattern called ```table tests``` can be used.

- Assume we have the following function in the table package:


```go
    func DoMath(num1, num2 int, op string) (int, error) {
        switch op {
            case "+":
                return num1 + num2, nil
            case "-":
                return num1 - num2, nil
            case "*":
                return num1 + num2, nil
            case "/":
                if num2 == 0 {
                    return 0, errors.New("division by zero")
                }
                return num1 / num2, nil
            default:
                return 0, fmt.Errorf("unknown operator %s", op)
        }
    }
```

```go
    data := []struct {
        name string
        num1 int
        num2 int
        op string
        expected int
        errMsg string
    }{
        {"addition", 2, 2, "+", 4, ""},
        {"subtraction", 2, 2, "-", 0, ""},
        {"multiplication", 2, 2, "*", 4, ""},
        {"division", 2, 2, "/", 1, ""},
        {"bad_division", 2, 0, "/", 0, `division by zero`},
    }
```
```go
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            result, err := DoMath(d.num1, d.num2, d.op)
            if result != d.expected {
                t.Errorf("Expected %d, got %d", d.expected, result)
            }
            var errMsg string
            if err != nil {
                errMsg = err.Error()
            }
            if errMsg != d.errMsg {
                t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
            }
        })
    }
```


- **Comparing error messages can be fragile, because there may not be any compatibility guarantees on the message text**. 

- **The function that we are testing uses ```errors.New``` and ```fmt.Errorf``` to make errors, so the only option is ```to compare the messages```**. 

- **If an error has a custom type, use ```errors.Is``` or ```errors.As``` to check that the correct error is returned**
