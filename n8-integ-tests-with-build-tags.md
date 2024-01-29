- We should still write integration tests, automated tests that connect to other services.

- These validate that your understanding of the service’s API is correct. 

- The challenge is figuring out how to group your automated tests; you only want to run integration tests when the support environment is present. 

- Also, integration tests tend to be slower than unit tests, so they are usually run less frequently.

- **The Go compiler provides build tags to control when code is compiled. Build tags are specified on the first line of a file with a magic comment that starts with ```// +build```**.

- The **original intent for build tags was to allow different code to be compiled on different platforms but they are also useful for splitting tests into groups**. 

- *Tests in files without build tags run all the time. These are the unit tests that don’t have dependencies on external resources*. 

- **Tests in files with a build tag are ```only run when the supporting resources are available```**.
- The first line of the file, separated from the package declaration by a newline is:

```go
    // +build integration
```
- To run our integration test alongside the other tests we’ve written, use:

```sh
    $ go test -tags integration -v ./...
```