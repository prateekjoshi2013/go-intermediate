### Code Coverage

- **Adding the ```-cover``` flag to the go test command calculates coverage information and includes a summary in the test output**. 

- **If you include a second flag ```-coverprofile```, you can save the coverage information to a file**:
  
    ```go
        go test -v -cover -coverprofile=c.out
    ```

- **The ```cover``` tool included with Go generates an ```HTML``` representation of your source code with that information**:

    ```go
        go tool cover -html=c.out
    ```

- When you run it, your web browser should open and show you a page