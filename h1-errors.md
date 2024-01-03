- Go handles errors by **returning a value of type error as the last return value for a function**. 

- This is entirely by convention, but it is such a **strong convention that it should never be breached**
  
- **When a function executes as expected, nil is returned for the error parameter. If something goes wrong, an error value is returned instead**. 
  
- **The calling function then checks the error return value by comparing it to nil, handling the error, or returning an error of its own**.

```go
    func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
        if denominator == 0 {
            return 0, 0, errors.New("denominator is 0")
        }
        return numerator / denominator, numerator % denominator, nil
    }

    func main() {
        numerator := 20
        denominator := 3
        remainder, mod, err := calcRemainderAndMod(numerator, denominator)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Println(remainder, mod)
    }
```


- **A new error is created from a string by calling the New function in the errors package**. 

- Error **messages should not be capitalized nor should they end with punctuation or a newline**. 

- In most cases, you should **set the other return values to their zero values when a non-nil error is returned. except for sentinel errors**.
  
- Go doesn’t have special constructs to detect if an error was returned. **Whenever a function returns, use an if statement to check the error variable to see if it is non-nil**
  
- **The Go compiler requires that all variables must be read. Making errors returned values forces developers to either check and handle error conditions or make it explicit that they are ignoring errors by using an underscore (_) for the returned error value**.

### Create Errors from String

- two ways to create an error from a string. 
  - The first is the errors.New function. It takes in a string and returns an error.
    
    - ```errors.New("only even numbers are processed")```
  
  - The second way is to use the fmt.Errorf function. 
  - This function allows you to use all of the formatting verbs for fmt.Printf to create an error.

    - ```fmt.Errorf("%d isn't an even number", i)```

### Sentinel Errors

- Some **errors are meant to signal that processing cannot continue due to a problem with the current state**.

- **Sentinel errors are one of the few variables that are declared at the package level**. 

- By **convention, their names start with Err (with the notable exception of io.EOF)**.

- They **should be treated as read-only**
  
- **Be sure you need a sentinel error before you define one**. 

- **Once you define one, it is part of your public API and you have committed to it being available in all future backward-compatible releases**. 

- **It’s far better to reuse one of the existing ones in the standard library** or **to define an error type that includes information about the condition that caused the error to be returned**.

- **If you have an ***error condition that indicates a specific state has been reached in your application where no further processing is possible*** and ***no contextual information*** needs to be used to explain the error state, a sentinel error is the correct choice**.

- **use == to test if the error was returned when calling a function whose documentation explicitly says it returns a sentinel error**

-  **Sentinel errors should be rare**, so they can be handled by convention instead of language rules.

- Yes, **they are public package-level variables. This makes them mutable, but it’s highly unlikely someone would accidentally reassign a public variable in a package**.

### Errors Are Values

- Since error is an interface, you can define your own errors that include additional information for logging or error handling. 
  
  - For example, you might want to include a status code as part of the error to indicate the kind of error that should be reported back to the user. 
  
  - This lets you avoid string comparisons (whose text might change) to determine error causes.

```go
type Status int

const (
    InvalidLogin Status = iota + 1
    NotFound
)

type StatusErr struct {
    Status Status
    Message string
}

func (se StatusErr) Error() string {
    return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
    err := login(uid, pwd)
    if err != nil {
        return nil, StatusErr{
            Status: InvalidLogin,
            Message: fmt.Sprintf("invalid credentials for user %s", uid),   
        }
    }
    
    data, err := getData(file)

    if err != nil {

        return nil, StatusErr{
            Status: NotFound,
            Message: fmt.Sprintf("file %s not found", file),
        }
    }
    return data, nil
}
```

- You **shouldn’t declare a variable to be the type of your custom error and then return that variable**.

    ```go
        func GenerateError(flag bool) error {
            var genErr StatusErr // should not be declared with the type StatusErr
            if flag {
                genErr = StatusErr {
                    Status: NotFound,
                }
            }
           return genErr
        }
        
        func main() {
            err := GenerateError(true)
            /*
                The reason why err is non-nil is that error is an interface.
                for an interface to be considered nil, both the underlying 
                type and the underlying value must be nil.
            */
            fmt.Println(err != nil) // true
            err = GenerateError(false)
            fmt.Println(err != nil) // true
        }
    ```
- Two solutions for this is:
  
    - The most common approach is to explicitly return nil for the error value when a function completes successfully:

        ```go
            func GenerateError(flag bool) error {
                if flag {
                    return StatusErr{
                        Status: NotFound,
                    }
                }
                return nil
            } 
        ```

    - Another approach is to make sure that any local variable that holds an error is of type error:

    ```go
        func GenerateError(flag bool) error {
            var genErr error
            if flag {
                genErr = StatusErr{
                    Status: NotFound,
                }
            }
            return genErr
        }
    ```

    ### Wrapping Errors

-***When an error is passed back through your code, you often want to add additional context to it. This context can be the name of the function that received the error or the operation it was trying to perform***. 

- **When you preserve an error while adding additional information, it is called wrapping the error**.

- **When you have a series of wrapped errors, it is called an error chain**

- The **fmt.Errorf function has a special verb, %w**. 

  - Use this to create an error whose formatted string includes the formatted string of another error and which contains the original error as well. 
  
  - **The convention is to write : %w at the end of the error format string and make the error to be wrapped the last parameter passed to fmt.Errorf**.

  - The **standard library also provides a function for unwrapping errors, the ***Unwrap function in the errors package***.** 
  
    - ***You pass it an error and it returns the wrapped error, if there is one. If there isn’t, it returns nil***.

    ```go

    func fileChecker(name string) error {
            f, err := os.Open(name)
            if err != nil {
                return fmt.Errorf("in fileChecker: %w", err) // wraps incoming error
            }
            f.Close()
            return nil
        }
        func main() {
            err := fileChecker("not_here.txt")
            if err != nil {
                fmt.Println(err)
                if wrappedErr := errors.Unwrap(err); wrappedErr != nil { // unwraps wrapped error
                    fmt.Println(wrappedErr)
                }
            }
        }
    ```

    - **If you want to wrap an error with your custom error type, your error type needs to implement the method Unwrap**. 
      
      - **This method takes in no parameters and returns an error**

        ```go
            type StatusErr struct {
                Status Status
                Message string
                err error
            }
            
            func (se StatusErr) Error() string {
                return se.Message
            }
            
            func (se StatusError) Unwrap() error {
                return se.err
            }
        ```

    - If you want to create a new error that contains the message from another error, but don’t want to wrap it, use fmt.Errorf to create an error, but use the %v verb instead of %w:

        ```go
            err := internalFunction()
            if err != nil {
                return fmt.Errorf("internal failure: %v", err)
            }
        ```

### Is and As

#### Is

    - **To check if the returned error or any errors that it wraps match a specific sentinel error instance, use errors.Is**. 
    - **It takes in two parameters, the error that is being checked and the instance you are comparing against**. 
    - **The errors.Is function returns true if there is an error in the error chain that matches the provided sentinel error**.

        ```go
            func fileChecker(name string) error {
                f, err := os.Open(name)
                if err != nil {
                    return fmt.Errorf("in fileChecker: %w", err)
                }
                f.Close()
                return nil
            }

            func main() {
                err := fileChecker("not_here.txt")
                if err != nil {
                    if errors.Is(err, os.ErrNotExist) {
                        fmt.Println("That file doesn't exist") // That file doesn't exist
                    }
                }
            }
        ```
    - **errors.Is uses == to compare each wrapped error with the specified error**. 
    
    - **If this does not work for an error type that you define (for example, if your error is a noncomparable type), implement the Is method on your error**

        ```go
            type MyErr struct {
                Codes []int
            }

            func (me MyErr) Error() string {
                return fmt.Sprintf("codes: %v", me.Codes)
            }

            func (me MyErr) Is(target error) bool {
                if me2, ok := target.(MyErr); ok {
                    return reflect.DeepEqual(me, me2)
                }
                return false
            }
        ```

#### As

  - **The errors.As function allows you to check if a returned error (or any error it wraps matches a specific type**. 
    
  - **It takes in two parameters. The first is the error being examined and the second is a pointer to a variable of the type that you are looking for**.
    
  - If the function returns true, an error in the error chain was found that matched, and that matching error is assigned to the second parameter. If the function returns false, no match was found in the error chain. 
    
  - Let’s try it out with MyErr:

    ```go
        err := AFunctionThatReturnsAnError()
        var myErr MyErr
        if errors.As(err, &myErr) {
            fmt.Println(myErr.Code)
        }
    ```

    - Note that you use var to declare a variable of a specific type set to the zero value. You then pass a pointer to this variable into errors.As.
    - You don’t have to pass a pointer to a variable of an error type as the second parameter to errors.As. 
    - You can pass a pointer to an interface to find an error that meets the interface:

    ```go
        err := AFunctionThatReturnsAnError()
        var coder interface {
            Code() int
        }
        if errors.As(err, &coder) {
            fmt.Println(coder.Code())
        }
    ```
  - We’re using an anonymous interface here, but any interface type is acceptable.
    
  - **If the second parameter to errors.As is anything other than a pointer to an error or a pointer to an interface, the method panics**.

- **Use errors.Is when you are looking for a specific instance or specific values. Use errors.As when you are looking for a specific type**.


### Wrapping Errors with defer

- Sometimes you find yourself wrapping multiple errors with the same message:

```go
    func DoSomeThings(val1 int, val2 string) (string, error) {
        val3, err := doThing1(val1)
        if err != nil {
            return "", fmt.Errorf("in DoSomeThings: %w", err)
        }
        
        val4, err := doThing2(val2)
        if err != nil {
            return "", fmt.Errorf("in DoSomeThings: %w", err)
        }

        result, err := doThing3(val3, val4)
        if err != nil {
            return "", fmt.Errorf("in DoSomeThings: %w", err)
        }
        return result, nil
    }
```
- We can **simplify this code by using defer**:

```go
    func DoSomeThings(val1 int, val2 string) (_ string, err error) {
        defer func() {
            if err != nil {
                err = fmt.Errorf("in DoSomeThings: %w", err)
            }
        }()

        val3, err := doThing1(val1)
        if err != nil {
            return "", err
        }

        val4, err := doThing2(val2)
        if err != nil {
            return "", err
        }

        return doThing3(val3, val4)
    }
```

