- REST APIs have enshrined JSON as the standard way to communicate between services, and Go’s standard library includes support for converting Go data types to and from JSON. 

- **```Marshaling``` means converting from a Go data type to an encoding**

- **```Unmarshaling``` means converting to a Go data type**.


### Use Struct Tags to Add Metadata

- An order management system and have to read and write the following JSON:

```json
    {
        "id":"12345",
        "date_ordered":"2020-05-01T13:01:02Z",
        "customer_id":"3",
        "items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
    }
```
- We define types to map this data:

```go
    type Order struct {
        ID string `json:"id"`
        DateOrdered time.Time `json:"date_ordered"`
        CustomerID string `json:"customer_id"`
        Items []Item `json:"items"`
    }

    type Item struct {
        ID string `json:"id"`
        Name string `json:"name"`
    }
```

- We **specify the rules for processing our JSON with struct tags, strings that are written after the fields in a struct**. 

- Even though struct tags are **strings marked with backticks, they cannot extend past a single line**. 

- Struct **tags are composed of one or more tag/value pairs, written as ```tagName:"tagValue"``` and separated by spaces**. 

- Because they are just strings, the compiler cannot **validate that they are formatted correctly, but ```go vet``` does**. 

- Like any other package, **the code in the encoding/json package cannot access an unexported field on a struct in another package**.
  
- For JSON processing, **we use the tag name json to specify the name of the JSON field that should be associated with the struct field**. 

- **If no json tag is provided, the default behavior is to assume that the name of the JSON object field matches the name of the Go struct field**. 

- Despite this default behavior, **it’s best to use the struct tag to specify the name of the field explicitly, even if \the field names are identical**.

- **When unmarshaling from JSON into a struct field with no json tag, the name match is case-insensitive**. 

- **When marshaling a struct field with no json tag back to JSON , the JSON field will always have an uppercase first letter, because the field is exported**.

- **If a field should be ignored when marshaling or unmarshaling, use a dash (-) for the name**. 

- **If the field should be left out of the output when it is empty, add ,omitempty after the name**.


### Unmarshaling and Marshaling

- **The Unmarshal function in the encoding/json package is used to convert a slice of bytes into a struct**. 

- If we have a string named data, this is the **code to convert data to a struct of type Order**:

```go
    var o Order
    err := json.Unmarshal([]byte(data), &o)
    if err != nil {
        return err
    }
```

- The *json.Unmarshal function populates data into an input parameter, just like the implementations of the io.Reader interface*. 

- Just **like io.Reader implementations, this allows for efficient reuse of the same struct over and over, giving you control over memory usage**.

- use **the Marshal function in the encoding/json package to write an Order instance back as JSON stored in a slice of bytes**:

```go
    out, err := json.Marshal(o)
```

### JSON, Readers, and Writers

- **The json.Decoder and json.Encoder types read from and write to anything that meets the io.Reader and io.Writer interfaces**

- **data in toFile, which implements a simple struct**:

```go
    type Person struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }

    toFile := Person {
        Name: "Fred",
        Age: 40,
    }
```

- The os.File type implements both the io.Reader and io.Writer interfaces, so we can use it to demonstrate json.Decoder and json.Encoder. 


  - First, we write toFile to a temp file by passing the temp file to json.NewEncoder, which returns a json.Encoder for the temp file. We then pass toFile to the Encode method:

    ```go
        tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
        if err != nil {
            panic(err)
        }
        defer os.Remove(tmpFile.Name())
        err = json.NewEncoder(tmpFile).Encode(toFile)
        if err != nil {
            panic(err)
        }
        err = tmpFile.Close()
        if err != nil {
            panic(err)
        }
    ```

  - We can read the JSON back in by passing a reference to the temp file to json.NewDecoder and then calling the Decode method on the returned json.Decoder with a variable of type Person:

    ```go

        tmpFile2, err := os.Open(tmpFile.Name())
        if err != nil {
            panic(err)
        }
        var fromFile Person
        err = json.NewDecoder(tmpFile2).Decode(&fromFile)
        if err != nil {
            panic(err)
        }
        err = tmpFile2.Close()
        if err != nil {
            panic(err)
        }
        fmt.Printf("%+v\n", fromFile)
    ```


### Encoding and Decoding JSON Streams

- What do you do **when you have multiple JSON structs to read or write at once**? Our friends **json.Decoder and json.Encoder can be used for these situations**, too.

```go

    const data = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for dec.More() {
        err := dec.Decode(&t)
		if err != nil {
            panic(err)
		}
        // This lets us read in the data, one JSON object at a time
		fmt.Println(t)  // process one object at a time
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}

    // we are writing to a bytes.Buffer, but any type that meets the io.Writer interface will work

    dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
		err = enc.Encode(t) // writing to buffer
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
```

#### Custom Parsing
- While the default functionality is often sufficient, there are times you need to override it

- While time.Time supports JSON fields in RFC 339 format out of the box, you might have to deal with other time formats. 

- We can handle this by creating a new type that implements two interfaces, json.Marshaler and json.Unmarshaler:

```go

    type RFC822ZTime struct {
        time.Time
    }

    func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
        out := rt.Time.Format(time.RFC822Z)
        return []byte(`"` + out + `"`), nil
    }

    func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
        if string(b) == "null" {
            return nil
        }
        t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
        if err != nil {
            return err
        }
        *rt = RFC822ZTime{t}
        return nil
    }

    type Order struct {
        ID string `json:"id"`
        DateOrdered RFC822ZTime `json:"date_ordered"`
        CustomerID string `json:"customer_id"`
        Items []Item `json:"items"`
    }

```

- **There is a philosophical drawback to this approach: we have allowed the date format of the JSON we are processing to change the types of the fields in our data structure**.

- **To limit the amount of code that cares about what your JSON looks like, define two different structs**. 
  
  - *Use one for converting to and from JSON and the other for data processing*.
  
  - *Read in JSON to your JSON-aware type, and then copy it to the other*.
  
  - *When you want to write out JSON, do the reverse*. 
  
  - *This does create some duplication, but it keeps your business logic from depending on wire protocols*.

- **You can pass a map[string]interface{} to json.Marshal and json.Unmarshal to translate back and forth between JSON and Go, but save that for the exploratory phase of your coding and replace it with a concrete type when you understand what you are processing**.
