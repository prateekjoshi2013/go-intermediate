### Imports and Exports

- Go’s import statement allows you to access exported constants, variables, functions, and types in another package. 

- A package’s exported identifiers (an identifier is the name of a variable, constant, type, function, method, or a field in a struct) cannot be accessed from another current package without an import statement.

- **how do you export an identifier in Go**? 
  
  - Rather than use a special keyword, **Go uses capitalization to determine if a package-level identifier is visible outside of the package where it is declared**. 
  
  - **An identifier whose name starts with an ***uppercase letter is exported***. Conversely, an identifier whose ***name starts with a lowercase letter or underscore can only be accessed from within the package*** where it is declared**.


- ***Anything you export is part of your package’s API. Before you export an identifier, besure that you intend to expose it to clients***. 

- ***Document all exported identifiers and keep them backward-compatible unless you are intentionally making a major version change***

### Creating and Accessing Packages

- Inside package_example, you’ll see two additional directories, math and formatter. 
- In math, there’s a file called math.go with the following contents:

```package_example/math/math.go```

```go
  package math

  func Double(a int) int {
    return a * 2
  }
```

- The **first line of the file is called the package clause**. 
  
  - It consists of the **keyword package and the name for the package**. 
  
  - **The package clause is always the first nonblank, noncomment line in a Go source file**.

- **In formatter, there’s a file called formatter.go with the following contents**:

```github.com/learning-go-book/package_example/formatter```

```go
  package print

  import "fmt"

  func Format(num int) string {
    return fmt.Sprintf("The number is %d", num)
  }
```

- Note that we said the **package name is print in the package clause, but it’s in the formatter directory**. 
  
- Finally, the following **contents are in the file main.go in the root directory**:

```main.go```

```go
  package main

  import (
    "fmt"
    "github.com/learning-go-book/package_example/formatter"
    "github.com/learning-go-book/package_example/math"
  )

  func main() {
    num := math.Double(2)
    output := print.Format(num)
    fmt.Println(output)
  }
```
- The first line of this file is familiar. 

- All of our programs before this chapter have put package main as the first line in our code. 

- Next, we have our import section. We’re imported three packages. 
  
  - The first is fmt, which is in the standard library. We’ve done this in previous chapters. The next two
  
  - **Imports refer to the packages within our program**. 
  
  - **You must specify an import path when importing from anywhere besides the standard library**. 
  
  - **The import path is built by appending the path to the package within the module to the module path**
  
  - ***We imported the print package with the import path **github.com/learning-go-book/package_ example/formatter**. That’s **because the name of a package is determined by its package clause, not its import path*****. 

- **As a general rule, you should make the name of the package match the name of the directory that contains the package. It is hard to discover a package’s name if it does not match the containing directory**. 
  
  - However, there are a few situations where you use a different name for the package than for the directory.
  
    - **The first is something we have been doing all along without realizing it. We declare a package to be a staring point for a Go application by using the special package name main. Since you cannot import the main package, this doesn’t produce confusing import statements**.
    
    - **If your directory name contains a character that’s not valid in a Go identifier, then you must choose a package name that’s different from your directory name. It’s better to avoid this by never creating a directory with a name that’s not a valid identifier**.
    
    - **The final reason for creating a directory whose name doesn’t match the package name is to support versioning using directories**.

### Naming Packages

- Having the package name as part of the name used to refer to items in the package
has some implications. 
  
  - The first is that package names should be descriptive. Rather than have a package called util, create a package name that describes the functionality provided by the package. 
  
  - For example, say you have two helper functions: 
  
    - one to extract all names from a string and another to format names properly. 
  
    - Don’t create two functions in a util package called ExtractNames and FormatNames. 
  
    - If you do, every time you use these functions, they will be referred to as util.ExtractNames and util.FormatNames, and that util package tells you nothing about what the functions do.
  
    - It’s better to create one function called Names in a package called extract and a second function called Names in a package called format. 
  
    - It’s OK for these two functions to have the same name, because they will always be disambiguated by their package names. 
  
    - The first will be referred to as extract.Names when imported, and the second will be referred to as format.Names.
  
  - You should also avoid repeating the name of the package in the names of functions and types within the package. 
  
  - Don’t name your function ExtractNames when it is in the extract package. 
  
  - The exception to this rule is when the name of the identifier is the same as the name of the package. 
  
    - For example, the package sort in the standard library has a function called Sort, and the context package defines the Context interface.

### Overriding package name

```go
  import (
    crand "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/rand"
  )
```

- **For conflicting package names for crypto/rand and math/rand we alias crypto/rand as crand**
- We use **_ "package_name"** if we just need to execute the init function for a package