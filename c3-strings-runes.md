- Go uses a **sequence of bytes to represent a string**. 
- These bytes don’t have to be in any particular character encoding, but several Go library functions (and the for range loop that we discuss in the next chapter) assume that a **string is composed of a sequence of UTF-8-encoded code points**.
- you can extract a single value from a string by using an index expression:

    ```go
        var s string = "Hello, 😼"
        var bs []byte = []byte(s)
        var rs []rune = []rune(s)
        fmt.Printf("%c\n",bs[7]) // ð
        fmt.Printf("%c\n",rs[7]) // 😼
        fmt.Println(len(bs)) // 11
        fmt.Println(len(rs)) // 8
    ```
- If using bytes if the  😼 is more than a byte in utf-8 code points it cannot be extracted fully so should be converted into rune slice
- we ***should extract substrings and code points from strings using the functions in the strings and unicode/utf8 packages in the standard library***.
