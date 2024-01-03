### Importing Third-Party Code

- Go compiles all code for your application into a single binary, whether it was code you wrote or code from third parties. 
  
- Just as when we imported a package from within our own project, when we import a third-party package, you specify the location in the source code repository where the package is located.

- **Never use floating point numbers when you need an exact representation of a decimal number**.

  - If you do need an exact representation, **one good library is the decimal module from ShopSpring**

- We’ll use both of these modules in a small program that accurately calculates the price of an item with the tax included and prints the output in a neat format

    ```go
    package main

    import (
        "fmt"
        "log"
        "os"
        "github.com/learning-go-book/formatter"
        "github.com/shopspring/decimal"
    )

        func main() {
            if len(os.Args) < 3 {
                fmt.Println("Need two parameters: amount and percent")
                os.Exit(1)
            }
            amount, err := decimal.NewFromString(os.Args[1])
            if err != nil {
                log.Fatal(err)
            }

            percent, err := decimal.NewFromString(os.Args[2])
            if err != nil {
                log.Fatal(err)
            }

            percent = percent.Div(decimal.NewFromInt(100))
            total := amount.Add(amount.Mul(percent)).Round(2)
            fmt.Println(formatter.Space(80, os.Args[1], os.Args[2],total.StringFixed(2)))
        }
    ```

  - The **two imports github.com/learning-go-book/formatter and github.com/shopspring/decimal specify third-party imports**.
  
  - Note that **they include the location of the package in the repository. Once imported, we access the exported items in these packages just like any other imported package**.


  - Do a build and we’ll see what happens:

    ```sh
        $ go build
        go: finding module for package github.com/shopspring/decimal
        go: finding module for package github.com/learning-go-book/formatter
        go: found github.com/learning-go-book/formatter in
        github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc380940ae
        go: found github.com/shopspring/decimal in github.com/shopspring/decimal v1.2.0
    ```

- **Because the location of the package is in the source code, go build is able to get the package’s module and download it**. 

- If you look in the **go.mod file now**, you’ll see:

- The **require section of the go.mod** file lists the modules that you’ve imported into
your module.

  - **After the module name is a version number**. 
  - In the case of the **formatter module, it doesn’t have a version tag, so Go makes up a pseudo-version**.
    
    ```go
        module github.com/learning-go-book/money

        go 1.15

        require (
            github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc380940ae
            github.com/shopspring/decimal v1.2.0
        )
    ```
    - Meanwhile, a go.sum file has been created with the contents:

    ```go
        github.com/google/go-cmp v0.5.2/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22db...
        github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc38094...
        github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc38094...
        github.com/shopspring/decimal v1.2.0 h1:abSATXmQEYyShuxI4/vyW3tV1MrKA...
        github.com/shopspring/decimal v1.2.0/go.mod h1:DKyhrW/HYNuLGql+MJL6WC...
        golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5...
    ```

- Whenever you **run any go command that requires dependencies (such as go run, go build, go test, or even go list), any imports that aren’t already in go.mod are downloaded to a cache**. 

  - **The go.mod file is automatically updated to include the module path that contains the package and the version of the module**. 

  - **The go.sum file is updated with two entries**: 
    
    - **one with the module, its version, and a hash of the module**,
    
    - **other with the hash of the go.mod file for the module**

- **Always include up-to-date go.mod and go.sum files. Doing so specifies exactly what versions of your dependencies are being used**.


