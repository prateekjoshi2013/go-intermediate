### Testing in Go

- Testing support has two parts:
  
  -  libraries 
  
  -  tooling. 
  
- The testing package in the standard library provides the types and functions to write tests, while the go test tool that’s bundled with Go runs your tests and generates reports. 

- Unlike many other languages, Go tests are placed in the same directory and the same package as the production code. 
  
  - Since tests are located in the same package, they are able to access and test unexported functions and variables.


- Let’s write a simple function and then a test to make sure the function works. In the

- file ```adder/adder.go```, we have:

    ```go
        func addNumbers(x, y int) int {
            return x + x
        }
    ```

- The corresponding test is in ```adder/adder_test.go```:

    ```go
        func Test_addNumbers(t *testing.T) {
            result := addNumbers(2,3)
            if result != 5 {
                t.Error("incorrect result: expected 5, got", result)
            }
        }
    ```
- Every **test is written in a file whose name ends with ```_test.go```**. 

- **If you are writing tests against ```foo.go```, place your tests in a file named ```foo_test.go```**.

- **Test functions ```start with the word Test``` and take in a single parameter of type ```*testing.T```**. 
  
  - By convention, this **parameter is named t**. 
  
  - Test **functions do not return any values**.
  
  - The **name of the test (apart from starting with the word “Test”) is meant to document what you are testing**, so pick something that explains what you are testing.
  
  - When writing **unit tests for individual functions, the convention is to name the unit test Test followed by the name of the function**. 
  
  - When **testing unexported functions, some people use an underscore between the word Test and the name of the function**.
  
  - call the code being tested and to validate if the responses are as expected. **When there’s an incorrect result, we report the error with the t.Error method, which works like the fmt.Print** function 

- **the command ```go test``` runs the tests in the current directory**:

```sh
    $ go test
    --- FAIL: Test_addNumbers (0.00s)
    adder_test.go:8: incorrect result: expected 5, got 4
    FAIL
    exit status 1
    FAIL test_examples/adder 0.006s
```

- The **go test command allows you to specify which packages to test**. 
  
  - **Using ```./…``` for the package name specifies that you want to run tests in the current directory and all of the subdirectories of the current directory**. 

- **Include a ```-v``` flag to get verbose testing output**.

### Reporting Test Failures

- There are several methods on *testing.T for reporting test failures. 

#### Error and Errof

- We’ve already seen **Error**, which builds a failure description string out of a **comma-separated list of values**.

- If you’d rather use a **Printf**-style formatting string to generate your message, use the
**Errorf** method instead:
    
```go
    t.Errorf("incorrect result: expected %d, got %s", 5, result)
```

- While **Error and Errorf mark a test as failed, the test function continues running**. 

#### Fatal and Fatalf

- If you think a test function should **stop processing as soon as a failure is found, use the ```Fatal``` and ```Fatalf``` methods**. 

- The Fatal method works like Error, and the Fatalf method works like Errorf. 

- **The difference is that the test function exits immediately after the test failure message is generated**. 

- Note that this doesn’t exit all tests; any remaining test functions will execute after the current test function exits.

### Setting Up and Tearing Down

- Sometimes you have some common state that you want **to set up before any tests run and remove when testing is complete**.
  
- Use a **```TestMain``` function to manage this state and run your tests**:

```go
    var testTime time.Time

    func TestMain(m *testing.M) {
        fmt.Println("Set up stuff for tests here")
        testTime = time.Now()
        exitVal := m.Run()
        fmt.Println("Clean up stuff after tests here")
        os.Exit(exitVal)
    }

    func TestFirst(t *testing.T) {
        fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
    }
    
    func TestSecond(t *testing.T) {
        fmt.Println("TestSecond also uses stuff set up in TestMain", testTime)
    }
```

- Both **TestFirst and TestSecond refer to the package-level variable testTime**. 

- We **declare a function called TestMain with a parameter of type ```*testing.M```**. 

- **Running go test on a package with a TestMain function calls the function instead of invoking the tests directly**. 

- **Once the state is configured, call the Run method on *testing.M to run the test functions**.

- **The Run method returns the exit code; 0 indicates that all tests passed**. 

- **Finally, you must call os.Exit with the exit code returned from Run**.

- **TestMain is invoked once, not before and after each individual test**.

- Also be aware that you can have **only one TestMain per package**.

- There are **two common situations where TestMain is useful**: 
  
  - **When you need to set up data in an external repository, such as a database**
  
  - **When the code being tested depends on package-level variables that need to be initialized**

### Cleanup Method

- The **```Cleanup``` method on *testing.T. is used to clean up temporary resources created for a single test**. 

- **This method has a single parameter, a function with no input parameters or return values**. 

- The **function runs when the test completes**. 

- **For simple tests, you can achieve the same result by using a ```defer``` statement, but ```Cleanup``` is useful when tests rely on ```helper``` functions to set up sample data**


```go
    // createFile is a helper function called from multiple tests
    func createFile(t *testing.T) (string, error) {
            f, err := os.Create("tempFile")
            if err != nil {
                return "", err
            }
            // write some data to f
            t.Cleanup(func() {
                os.Remove(f.Name())
            })
            return f.Name(), nil
    }

    func TestFileProcessing(t *testing.T) {
        fName, err := createFile(t)
        if err != nil {
            t.Fatal(err)
        }
        // do testing, don't worry about cleanup
    }
```

### Test Data

- **Go test uses the current package directory as the current working directory**. 

- If you want to use **sample data to test functions in a package, create a subdirectory named testdata to hold your files**. 

- **Go reserves this directory name as a place to hold test files**. 

- **When reading from testdata, always use a ```relative file reference```**. 

- Since **go test changes the current working directory to the current package, each package accesses its own testdata via a relative file path**.

### Caching Test Results

- **Go caches compiled packages if they haven’t changed**, Go **also caches test results when running tests across multiple packages if they have passed and their code hasn’t changed**. 

- **The tests are recompiled and rerun if you change any file in the package or in the testdata directory**. 

- You can also **force tests to always run if you pass the flag ```-count=1``` to go test**.

### Testing Your Public API

- If you want **to test just the public API of your package**, **Go has a convention for specifying this**. 

- Keep **your test source code in the same directory as the production source code, but you use ```packagename_test``` for the package name**

- If we have the following function in the adder package:

```go
    func AddNumbers(x, y int) int {
        return x + y
    }
```

- we can test it as **public API using the following code in a file in the adder package named adder_public_test.go**:
  
```go
    package adder_test

    import (
        "testing"
        "test_examples/adder"
    )

    func TestAddNumbers(t *testing.T) {
        result := adder.AddNumbers(2, 3)
        if result != 5 {
            t.Error("incorrect result: expected 5, got", result)
        }
    }
```

### ```go-cmp``` to Compare Test Results

- It can be verbose to write a **thorough comparison between two instances of a compound type**. 

- While you **can use reflect.DeepEqual to compare structs, maps, and slices**, there’s a **better way**

- Google released a **third-party module called ```go-cmp```** that does **the comparison for you and returns a detailed description of what does not match**

```go

    type Person struct {
        Name string
        Age int
        DateAdded time.Time
    }

    func CreatePerson(name string, age int) Person {
        return Person{
            Name: name,
            Age: age,
            DateAdded: time.Now(),
        }
    }
```
- In our **test file, we need to ```import github.com/google/go-cmp/cmp```**, and our test function looks like this:

```go
    func TestCreatePerson(t *testing.T) {
        expected := Person{
            Name: "Dennis",
            Age: 37,
        }
        
        result := CreatePerson("Dennis", 37)
        
        if diff := cmp.Diff(expected, result); diff != "" {
            t.Error(diff)
        }
    }
```
- The ```cmp.Diff``` function takes in the expected output and the output that was returned by the function that we’re testing.

```sh
    $ go test
    --- FAIL: TestCreatePerson (0.00s)
    ch13_cmp_test.go:16: ch13_cmp.Person {
            Name: "Dennis",
            Age: 37,
    -       DateAdded: s"0001-01-01 00:00:00 +0000 UTC",
    +       DateAdded: s"2020-03-01 22:53:58.087229 -0500 EST m=+0.001242842",
        }
    FAIL
    FAIL ch13_cmp 0.006s
```

- **To ```ignore``` the DateAdded field during comparison** . 
- **Specify comparator function. Declare the function as a local variable in your test**:

```go
    comparer := cmp.Comparer(func(x, y Person) bool {
        return x.Name == y.Name && x.Age == y.Age
    })
```

- P**ass a function to the cmp.Comparer function to create a customer comparator**. 
- The function that’s passed in must have two parameters of the same type and return a bool.
- Then change your call to cmp.Diff to include comparer:

```go
    if diff := cmp.Diff(expected, result, comparer); diff != "" {
        t.Error(diff)
    }
```