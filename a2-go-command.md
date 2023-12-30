### go mod 

- #### go mod init

  - To create a go module

```sh
    go mod init github.com/m
    go mod init github.com/m/v2
```

- #### go mod tidy

  - Ensures that the go.mod file matches the dependencies in your project's source code. It adds any missing dependencies and removes unnecessary ones.
  
```sh
      go mod tidy
```

- #### go mod vendor

  - Copies dependencies into the vendor directory within your project, allowing for a more reproducible build.
  
```sh
      go mod vendor
```

### go run

- The go run command does in fact compile your code into a binary. However, the binary is built in a temporary directory. 
  
- The go run command builds the binary, executes the binary from that temporary directory, and then deletes the binary after your program finishes.

### go build

- Most of the time you want to build a binary for later use. That’s where you use the **go
build** command.
  
```sh
    go build hello.go
```

- The name of the binary matches the name of the file or package that you passed in. If you want a different name for your application, or if you want to store it in a different location, use the **-o flag**.

```sh
    go build -o hello_world hello.go
```

- If you are cross-compiling for a different operating system along with a different architecture, you can set both **GOARCH** and **GOOS**:

```sh
    GOARCH=target_architecture GOOS=target_os go build
```

- To find all the GOARCH and GOOS combinations"

```sh
    go tool dist list
```


### go fmt

- To fix file format
    
```sh
    go fmt ./...
```

### go get 

- get or install a dependency

```sh
    go get example.com/package/path
```


- get or install a dependency and updated all dependencies

```sh
    go get -u example.com/package/path
```

### go vet 

- **go vet** is to identify issues that might not be caught by the Go compiler but can still lead to runtime errors or unexpected behavior. 

- It performs static analysis on the code, looking for common pitfalls, misuse of language features, and other potential problems

  -  Run go vet on a specific package

```sh
        go vet package/path
```  

  -  Run go vet on specific files
    
```sh
        go vet file1.go file2.go
```  





