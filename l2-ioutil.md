- The ioutil package provides some simple utilities for things like
  
  - *reading entire io.Reader implementations into byte slices at once*, 
  - *reading and writing files*, and 
  - *working with temporary files*. 

- The **ioutil.ReadAll, ioutil.ReadFile, and ioutil.WriteFile functions are fine for small data sources, but it’s better to use the Reader, Writer, and Scanner in the bufio package to work with larger data sources**.

### Adapter to add a method to an interface 

- One of the more clever functions in ioutil demonstrates **a pattern for adding a method to a Go type**. 

- **If you have a type that implements io.Reader but not io.Closer (such as strings.Reader) and need to pass it to a function that expects an io.ReadCloser, pass your io.Reader into ioutil.NopCloser and get back a type that implements io.ReadCloser**. 

- If you look at the implementation, it’s very simple:

```go
    type nopCloser struct {
        io.Reader
    }

    func (nopCloser) Close() error { return nil }

    func NopCloser(r io.Reader) io.ReadCloser {

        return nopCloser{r}

    }
```

- **Any time you need to add additional methods to a type so that it can meet an interface, use this embedded type pattern**.

- **The ioutil.NopCloser function violates the general rule of not returning an interface from a function, but it’s a simple adapter for an interface that is guaranteed to stay the same because it is part of the standard library**.