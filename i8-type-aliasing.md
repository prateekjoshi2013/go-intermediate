- After using a module for a while, you might realize that its API is not ideal. You might **want to rename some of the exported identifiers or move them to another package within your module**. 

- **To avoid a backward-breaking change, don’t remove the original identifiers; provide an alternate name instead**

- **With a function or method, this is easy. You declare a function or method that calls the original**. 
  
- **For a constant, simply declare a new constant with the same type and value, but a different name**.

- **When you want to rename or move an exported type, you have to use an alias**. Quite simply, an alias is a new name for a type.

    ```go
        type Foo struct {
            x int
            S string
        }
        func (f Foo) Hello() string {
            return "hello"
        }
        func (f Foo) goodbye() string {
            return "goodbye"
        }

        //If we want to allow users to access Foo by the name Bar, all we need to do is:
        type Bar = Foo
    ```

  - **The alias can even be assigned to a variable of the original type without a type conversion**

  - If you want **to add new methods or change the fields in an aliased struct, you must add them to the original type**.
  
  - You **can alias a type that’s defined in the same package as the original type or in a different package**. 
  
  - You **can even alias a type from another module**. 
  
  - There is **one drawback to an alias in another package**:
   
    - **you cannot use an alias to refer to the unexported methods and fields of the original type**