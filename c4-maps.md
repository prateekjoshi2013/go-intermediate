- The map type is written as **map[keyType]valueType**.

- A var declaration to create a map variable that’s set to its zero value:

    ```go
        var nilMap map[string]int
    ```

  - In this case, nilMap is declared to be a map with string keys and int values. 

  - The zero value for a map is nil. 

  - A nil map has a length of 0. 

  - **Attempting to read a nil map always returns the zero value for the map’s value type. However, attempting to write to a nil map variable causes a panic**.

- We can use a **:=** declaration to create a map variable by assigning it a map literal:

    ```go
        totalWins := map[string]int{}
    ```

  - In this case, we are using an **empty map literal**
  
  - This is **not the same as a nil map**. 
  
  - It **has a length of 0, but you can read and write to a map assigned an empty map literal**.

- a nonempty map literal looks like:

    ```go
        teams := map[string][]string {
            "Orcas": []string{"Fred", "Ralph", "Bijou"},
            "Lions": []string{"Sarah", "Peter", "Billie"},
            "Kittens": []string{"Waldo", "Raul", "Ze"},
        }
    ```

- If we know how many key-value pairs you intend to put in the map, but don’t know the exact values, you can **use make to create a map with a default size**:

    ```go
        ages := make(map[int][]string, 10)
    ```

- ***Maps created with make still have a length of 0, and they can grow past the initially specified size***

- **Maps automatically grow as you add key-value pairs to them**.

- ***Order of keys is random***

- If you know how many key-value pairs you plan to insert into a map, you can **use make to create a map with a specific initial size**.

- Passing a map to the **len function tells you the number of key-value pairs** in a map.

- **The zero value for a map is nil**.

- **Maps are not comparable**. You can check if they are equal to nil, but you cannot check if two maps have identical keys and values using == or differ using !=.

- The **key for a map can be any comparable type**. This means you **cannot use a slice or a map as the key for a map**.


### Reading and Writing a Map

    ```go
        totalWins := make(map[string]int,0)
        totalWins["Orcas"] = 1 // writing
        totalWins["Lions"] = 2
        fmt.Println(totalWins["Orcas"]) // 1  reading 
        fmt.Println(totalWins["Kittens"]) // returns zero value of value type (int)
    ```

### comma ok idiom

- Go provides the **comma ok idiom to tell the difference between a key that’s associated with a zero value and a key that’s not in the map**:

    ```go
        m := map[string]int{
            "hello": 5,
            "world": 0,
        }
        v, ok := m["hello"] 
        fmt.Println(v, ok) // 5, true
        v, ok = m["world"]
        fmt.Println(v, ok) // 0, true
        v, ok = m["goodbye"]
        fmt.Println(v, ok) // 0, false
    ```

### Deleting from Maps

- **Key-value pairs are removed** from a map via the built-in **delete** function:

    ```go
        m := map[string]int{
            "hello": 5,
            "world": 10,
        }
        delete(m, "hello")
    ```
- **If m is nil or there is no such element, delete is a no-op**

### Using Maps as Sets

- Go doesn’t include a set, but you can use a map to simulate some of its features.
  
- Use struct{} for the value when a map is
being used to implement a set. 

- The advantage is that an empty struct uses zero bytes, while a
boolean uses one byte.

```go
intSet := map[int]struct{}{}
vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
for _, v := range vals {
    intSet[v] = struct{}{}
}
if _, ok := intSet[5]; ok {
    fmt.Println("5 is in the set") // 5 is in the set
}
```