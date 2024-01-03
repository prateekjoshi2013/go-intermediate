### iota Is for Enumerations—Sometimes

- iota, which lets you assign an increasing value to a set of constants.

- When using iota, the best practice is to first define a type based on int that will represent
all of the valid values:

- Next, use a const block to define a set of values for your type:
  
    ```go
    type MailCategory int
    
    const (
        Uncategorized MailCategory = iota // 0
        Personal // 1
        Spam // 2
        Social // 3
        Advertisements // 4
    )
    ```
- **iota based enumerations only make sense when you care about being able to differentiate between a set of values, and don’t particularly care what the value is behind the scenes. If the actual value matters, specify it explicitly**.

- Don’t use iota for defining constants where its values are explicitly defined (elsewhere).
  
- For example, when implementing parts of a specification and the specification says which values are assigned to which constants, you should explicitly write the constant values.

- Use iota for “internal” purposes only. That is, where the constants are referred
to by name rather than by value.

- That way you can optimally enjoy iota by inserting
new constants at any moment in time / location in the list without the risk of breaking
everything.