
- ***Library management in Go is based around three concepts: **repositories**, **modules**, and **packages*****. 

- **A ***repository*** is a place in a version control system where the source code for a project is stored**. 

- **A ***module*** is the root of a Go library or application, stored in a repository**. 

- **Modules consist of one or more ***packages***, which give the module organization and structure**.
  
- **While you can store more than one module in a repository, it isn’t encouraged. Everything within a module is versioned together. Maintaining two modules in one repository means tracking separate versions for two different projects in a single repository**

### go.mod

- **Before we can use code from packages outside of the standard library, we need to make sure that we have declared that our project is a module**. 

  - **Every module has a globally unique identifier**.

  - In Go, we usually **use the path to the module repository where the module is found as globally unique identifier**. 
    
    - For example, scotch, a module I wrote to simplify relational database access in Go, can be found at GitHub. It has a module path of github.com/prateek/scotch

  - **A collection of Go source code becomes a module when there’s a valid go.mod file in its root directory**. 

  - Rather than create this file manually, we use the subcommands of the go mod command to manage modules. **The command ```go mod init MODULE_PATH``` creates the go.mod file that makes the current directory the root of a module**. 

  - **The MODULE_PATH is the globally unique name that identifies your module**

    ```go.mod```

    ```go
        
        module github.com/learning-go-book/money
        
        go 1.15

        require (
            github.com/learning-go-book/formatter v0.0.0-20200921021027-5abc380940ae
            github.com/shopspring/decimal v1.2.0
        )
    ```
    - **go.mod file starts with a module declaration that consists of the word module and the module’s unique path**. 
    
    - Next, **the go.mod file specifies the minimum compatible version of Go**. 
    
    - Finally, **the require section lists the modules that your module depends on and the minimum version required for each one**
    
    - ***There are ***two optional sections*** as well***. 
    
      - The ```replace section``` **lets you override the location where a dependent module is located**
      
      - The ```exclude section``` **prevents a specific version of a module from being used**.