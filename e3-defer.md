### Defer

- Programs often create temporary resources, like files or network connections,that need to be cleaned up. 

- This cleanup has to happen, no matter how many exit points a function has, or whether a function completed successfully or not. 

- In Go, the cleanup code is attached to the function with the defer keyword.

    ```go
        func main() {

            if len(os.Args) < 2 {
                log.Fatal("no file specified")
            }

            f, err := os.Open(os.Args[1])

            if err != nil {
                log.Fatal(err)
            }

            defer f.Close() // called no matter what after the function is exited
        }
    ```

- **Once we know we have a valid file handle, we need to close it after we use it, no matter how we exit the function**. 

- To ensure the cleanup code runs, we use the **defer keyword**, followed by a function or method call. In this case, we use the Close method on the file variable.
  
- Normally, a function call runs immediately, but defer delays the invocation until the surrounding function exits.

- You can **defer multiple closures in a Go function**. 

- They **run in last-in-first-out order; the last defer registered runs first**

- There’s a way for a deferred function to examine or modify the return values of its surrounding function. 

- There is, and it’s the best reason to use named return values. It allows us to take actions based on an error.


```go
        func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
            tx, err := db.BeginTx(ctx, nil)
            
            if err != nil {
                return err
            }
            
            defer func() {

            if err == nil {
                err = tx.Commit()
            }

            if err != nil {
                tx.Rollback()
            }

        }()

        _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)

        if err != nil {
            return err
        }
        
        // use tx to do more database inserts here
        return nil
    }
```

- First we write a helper function that opens a file and returns a closure:

    ```go
        func getFile(name string) (*os.File, func(), error) {
            file, err := os.Open(name)
            if err != nil {
                return nil, nil, err
            }
            return file, func() {
                file.Close()
            }, nil
        }
    ```
- Our helper function returns a file, a function, and an error. That * means that a file
reference in Go is a pointer.

- We’ll talk more about that in the next chapter. Now in main, we use our getFile function:

```go
    f, closer, err := getFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer closer()
```
