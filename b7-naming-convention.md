### naming variables and constants

- Identifier names start with a **letter or underscore**, and the **name can contain numbers, underscores, and letters**

- Go uses **camel case** for both **variables and constants* (names like **indexCounter** or **numberTries**)
  
- **Within a function, favor short variable names**. 

- The **smaller the scope for a variable,
the shorter the name** that’s used for it. 

- It is very common in Go to see **single-letter variable names**. 
  
  - For example, the names ***k*** and ***v*** (short for key and value) are used as the variable names in a for-range loop. If you are using a standard ***for loop***, ***i*** and ***j*** are common names for the ***index*** variable.
  
  - Use the first letter of a type as the variable name (for example, ***i for integers f for floats, b for booleans***). When you define your ***own types***, similar patterns apply.
  
  - These short names serve two purposes. 
    
    - The first is that they eliminate repetitive typing, **keeping your code shorter**. 
    
    - Second, they **serve as a check on how complicated your code is**. 

    - If you find it hard to keep track of your short-named variables, it’s likely that your block of code is doing too much.

  - When naming ***variables and constants*** in the ***package block***, use ***more descriptive names***.