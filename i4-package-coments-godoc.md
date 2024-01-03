- **Go has its own format for writing comments that are automatically converted into documentation**. 

- It’s called **godoc format**
- There are **no special symbols in a godoc comment**. They **just follow a convention**. 

- Here are the **rules**:
  
  - **Place the comment directly before the item being documented with no blank lines between the comment and the declaration of the item**.
  
  - **Start the comment with two forward slashes (//) followed by the name of the item**.
  
  - **Use a blank comment to break your comment into multiple paragraphs**.
  
  - **Insert preformatted comments by indenting the lines**.

- Comments before the package declaration create package-level comments. 

- If you have lengthy comments for the package (such as the extensive formatting documentation
in the fmt package), the convention is to put the comments in a file in your package called doc.go.

```A package-level comment```

```go
// Package money provides various utilities to make it easy to manage money.
package money
```

- Next, we place a comment on an exported struct. Notice that it
starts with the name of the struct.

```A struct comment```

```go
// Money represents the combination of an amount of money
// and the currency the money is in.
type Money struct {
    Value decimal.Decimal
    Currency string
}
```
- Finally, we have a comment on a function 

```A well-commented function```

```go
// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
// USD - US Dollar
// CAD - Canadian Dollar
// EUR - Euro
// INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
    // ...
}
```

- Go includes a ***command-line tool called **go doc** that views godocs***.
  
- ***The command **go doc PACKAGE_NAME** displays the package godocs for the specified package and a list of the identifiers in the package***. 

- ***Use **go doc PACKAGE_NAME.IDENTIFIER_NAME** to display the documentation for a specific identifier in the package***.