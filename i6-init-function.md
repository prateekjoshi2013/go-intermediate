- *There is a way to set up state in a package without explicitly calling anything*: **the init function**. 

- When you declare **a function named init that takes no parameters and returns no values, it runs the first time the package is referenced by another package**. 

- **Since init functions do not have any inputs or outputs, they can only work by side effect, interacting with package-level functions and variables**.

- Go **allows you to declare multiple init functions in a single package, or even in a single file in a package**

- There’s a documented order for **running multiple init functions in a single package**, **but rather than remembering it, it’s better to simply avoid them**.

- *Some packages, like database drivers, use init functions to register the database driver*. 

- However, **you don’t use any of the identifiers in the package. As mentioned earlier, Go doesn’t allow you to have unused imports**. 

- To work around this, **Go allows blank imports, where the name assigned to an import is the underscore (_)**. 
  
- *Just as an underscore allows you to skip an unused return value from a function, a **blank import triggers the init function** in a package but **doesn’t give you access to any of the exported identifiers** in the package*

    ```go
        import (
            "database/sql"
            _ "github.com/lib/pq"
        )
    ```
### Primary Use of Init functions

- The **primary use of init functions today is to initialize package-level variables that can’t be configured in a single assignment**. 

- It’s a **bad idea to have mutable state at the top level of a package, since it makes it harder to understand how data flows through your application**. 

- That means that **any package-level variables configured via init should be effectively immutable**. While Go doesn’t provide a way to enforce that their value does not change, you should make sure that your code does not change them. 

- If you have **package-level variables that need to be modified while your program is running**, see if you can refactor your code to **put that state into a struct that’s initialized and returned by a function in the package**.

- There are a couple of **additional caveats on the use of init**. 
  
  - *You should only **declare a single init function per package**, even though Go allows you to define multiple*. 

  - *If your **init function loads files or accesses the network, document this behavior,** so that security-conscious users of your code aren’t surprised by unexpected I/O*.
