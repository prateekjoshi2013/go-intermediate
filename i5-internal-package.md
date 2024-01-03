- Sometimes you want to share a function, type, or constant between packages in your module, but you don’t want to make it part of your API. 

- Go supports this via the special internal package name.

- When you create a package called internal, the exported identifiers in that package and its subpackages are only accessible to 
  
  - the direct parent package of internal 
  - the sibling packages of internal

- We’ve declared a simple function in the internal.go file in the internal package:

```internal/internal.go```

```go
func Doubler(a int) int {
    return a * 2
}
```
- Root structure

```
bar
- bar.go
example.go
foo
- foo.go
- internal
  - internal.go
- sibling
  - sibling.go
go.mod
```

the internal package is accessible only in foo.go, sibling.go