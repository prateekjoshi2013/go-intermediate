
- The **reflect.StructOf function takes in a slice of reflect.StructField and returns a reflect.Type that represents a new struct type**. 
  
  - **These structs can only be assigned to variables of type interface{}, and their fields can only be read and written using reflection**.

- While **we can use reflection to create new functions and new struct types**, **there’s no way to use reflection to add methods to a type**.

- **Using reflection is roughly 30 times ```slower``` than a custom function for string filtering and nearly 70 times slower for ints**. 
  
- **It uses significantly ```more memory and performs thousands of allocations```, which creates additional work for the garbage collector**.