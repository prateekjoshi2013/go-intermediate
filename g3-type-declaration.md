
### Type Declarations Aren’t Inheritance

- In addition to declaring types based on built-in Go types and struct literals, you can
also declare a user-defined type based on another user-defined type:

    ```go
        type HighScore Score
        type Employee Person
    ```

  - You **can't assign an instance of type HighScore to a variable of type Score or vice versa without a type conversion** 
  - **Nor can you assign either of them to a variable of type int without a type conversion**. 
  - **Any methods defined on Score aren’t defined on HighScore**

    ```go
        // assigning untyped constants is valid
        var i int = 300
        var s Score = 100
        var hs HighScore = 200
        hs = s // compilation error!
        s = i // compilation error!
        s = Score(i) // ok
        hs = HighScore(s) // ok
    ```

  - For **user-defined types whose underlying types are built-in types, a user-declared type can be used with the operators for those types**.
  
  - As we see in the preceding code, **they can also be assigned literals and constants compatible with the underlying type**.

### When to use Type Declarations

- **Types are documentation**.

  - They **make code clearer by providing a name for a concept and describing the kind of data that is expected**. 

  - ***It’s clearer for someone reading your code when a method has a parameter of type Percentage than of type int, and it’s harder for it to be invoked with an invalid value***.

- For user-defined type based on another user defined
type. 

  - When **we have the same underlying data, but different sets of operations to perform, make two types**. 
    
  - **Declaring one as being based on the other avoids some repetition and makes it clear that the two types are related**