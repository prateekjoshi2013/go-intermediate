- In main.go, we have the following third-party imports:

```go
    "github.com/learning-go-book/simpletax"
    "github.com/shopspring/decimal"
```
- on building the module

```sh
    $ go build
    go: finding module for package github.com/learning-go-book/simpletax
    go: finding module for package github.com/shopspring/decimal
    go: downloading github.com/learning-go-book/simpletax v1.1.0
    go: found github.com/learning-go-book/simpletax in
    github.com/learning-go-book/simpletax v1.1.0
    go: found github.com/shopspring/decimal in github.com/shopspring/decimal v1.2.0
```

- go.mod file has been updated to (go.sum is also updated):

```go
    module region_tax
  
    go 1.15
  
    require (
        github.com/learning-go-book/simpletax v1.1.0
        github.com/shopspring/decimal v1.2.0
    )
```

- There might be a bug in this latest version of the library. 
- By default, Go picks the latest version of a dependency when you add it to your project. 
- However, one of the things that makes versioning useful is that you can specify an earlier version of a module. 
- First we can see what versions of the module are available with the go list command:

```sh
    $ go list -m -versions github.com/learning-go-book/simpletax
    github.com/learning-go-book/simpletax v1.0.0 v1.1.0
```

- **By default, the go list command lists the packages that are used in your project**. 
  
  - The **-m flag changes the output to list the specified modules instead**, 
  
  - and the **-versions flag changes go list to report on the available versions for the specified module**. 
  
  - In this case, we see that there are two versions, v1.0.0 and v1.1.0. 
  
- **Let’s downgrade to version v1.0.0 and see if that fixes our problem**. 

```sh
    $ go get github.com/learning-go-book/simpletax@v1.0.0
```

- You might see **dependencies labeled // indirect in your go.mod file. These are dependencies that aren’t declared in your project directly**.


### Minimum Version Selection

- When project will depend on two or more modules that all depend on the same module. 

- As often happens, these modules declare that they depend on different minor or patch versions of that module. 

- Go resolves this by using the lowest version of a dependency that is declared to work in all
of the go.mod files across all of your dependencies.

- for example:

  - Let’s say that your module directly depends on modules A, B, and C. 

  - All three of these modules depend on module D.
    
    - The go.mod file for module A declares that it depends on v1.1.0, 
    
    - module B declares that it depends on v1.2.0,
    
    - module C declares that it depends on v1.2.3.

  - Go will import module D only once, and it will choose version v1.2.3

### Updating to Compatible Versions

- **If we upgraded to version v1.1.0 using go get github.com/learning-go-book/simpletax@v1.1.0** and then ran ```go get -u=patch github.com/learning-go-book/simpletax```, we **would be upgraded to version v1.1.1**.
  
- use the command ```go get -u github.com/learning-go-book/simpletax``` to **get the most recent version of simpletax. That upgrades us to version v1.2.1**.

### Updating to Major Versions

- Go modules follow the **semantic import versioning rule**.

- There are *two parts to this rule*:
  
  - **The major version of the module must be incremented**.
  
  - **For all major versions besides 0 and 1, the path to the module must end in vN, where N is the major version**.
  
- **The path changes because an import path uniquely identifies a package and, by definition, incompatible versions of a package are not the same package**. 

- **Using different paths means that you can import two incompatible versions of a package into different parts of your program, allowing you to upgrade gracefully**.

- Let’s see how this changes our program. First, we are going to change our import of
simpletax to:

    ```go
        "github.com/learning-go-book/simpletax/v2"
    ```

- This changes our import to refer to the ***v2 module***.

- When we call go build, our dependency is automatically updated:

```sh
    $ go build
    go: finding module for package github.com/learning-go-book/simpletax/v2
    go: downloading github.com/learning-go-book/simpletax/v2 v2.0.0
    go: found github.com/learning-go-book/simpletax/v2 in
    github.com/learning-go-book/simpletax/v2 v2.0.0
```

- We can look at the **go.mod file** and we’ll see that the new version of simpletax is included:

```go
    module region_tax

    go 1.15

    require (
        github.com/learning-go-book/simpletax v1.0.0 // indirect
        github.com/learning-go-book/simpletax/v2 v2.0.0
        github.com/shopspring/decimal v1.2.0
    )
```
- And go.sum has been updated as well:

```go
    github.com/learning-go-book/simpletax v1.0.0 h1:iH+7ADkdyrSqrMR2GzuWS...
    github.com/learning-go-book/simpletax v1.0.0/go.mod h1:/YqHwHy95m0M4Q...
    github.com/learning-go-book/simpletax v1.1.0 h1:Z/6s1ydS/vjblI6PFuDEn...
    github.com/learning-go-book/simpletax v1.1.0/go.mod h1:/YqHwHy95m0M4Q...
    github.com/learning-go-book/simpletax/v2 v2.0.0 h1:cZURCo1tEqdw/cJygg...
    github.com/learning-go-book/simpletax/v2 v2.0.0/go.mod h1:DVMa7zPtIFG...
    github.com/shopspring/decimal v1.2.0 h1:abSATXmQEYyShuxI4/vyW3tV1MrKA...
    github.com/shopspring/decimal v1.2.0/go.mod h1:DKyhrW/HYNuLGql+MJL6WC...
```
- The **old versions of simpletax are still referenced , even though they are no longer used. While this doesn’t cause any problems, Go includes a command to remove unused versions**:

```sh
    go mod tidy
```

- **After running this command, you’ll see that the only version referenced in go.mod and go.sum is v2.0.0**.

### Vendoring

- **To ensure that a module always builds with identical dependencies, some organizations like to keep copies of their dependencies inside their module**. 

- **This is known as vendoring**. 

- **It’s enabled by running the command go mod vendor. This creates a directory called *vendor* at the top level of your module that contains all of your module’s dependencies**.

- **If new dependencies are added to go.mod or versions of existing dependencies are upgraded with go get, you need to run go mod vendor again to update the vendor directory**.
  
- **If you forget to do this, go build, go run, and go test will refuse to run and display an error message**.


