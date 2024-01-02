# go-intermediate

- Given the subtle bugs that shadowing can introduce, it’s a good idea to make sure that you don’t have any shadowed variables in your programs

- add shadowing detection to your build process by installing the shadow linter on your machine:

```go
    go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

- If you are building with a Makefile, consider including shadow in the vet task:

    ```Make
        vet:
            go vet ./...
            shadow ./...
        .PHONY:vet
    ```