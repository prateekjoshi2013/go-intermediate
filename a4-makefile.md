```Makefile

    .DEFAULT_GOAL := build

    fmt:
        go fmt ./...
    .PHONY:fmt
    
    lint: fmt
        golint ./...
    .PHONY:lint
    
    vet: fmt
        go vet ./...
    .PHONY:vet

    build: vet
        go build hello.go
    .PHONY:build

```

- Each possible operation is called a ***target***
- The ***.DEFAULT_GOAL*** defines which target is run when no target is specified. In our case, we are going to run the build target.
- The word before the colon (:) is the name of the target. 
  
  - Any words after the target (like vet in the line build: vet) are the other targets that must be run before the specified target runs. 
  
  - The tasks that are performed by the target are on the indented lines after the target.

  - The .PHONY line keeps make from getting confused if you ever create a directory in your project with the same name as a target.
