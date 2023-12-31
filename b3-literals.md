- A literal in Go refers to writing out a 
  
  - **number**
    - **Integer literals** are sequences of numbers
      - **0b for binary (base two)**, 
      - **0o for octal (base eight)**, or 
      - **0x for hexadecimal (base sixteen)**. 
      - You can use either or **upper- or lowercase** letters for the prefix.
      - Go allows you to put **underscores** in the middle of your literal. This allows you to, for example, group by thousands in base ten like (**1_234**)
    -  **Float literals** have decimal points to indicate the fractional portion of the value.      
       -  They can also have an exponent specified with the letter e and a positive or negative number (such as **6.03e23**).
       -  you can use *underscores* to format your floating point literals


  - **character** 
    
    - Rune literals represent characters and are surrounded by single quotes. 
    
    - Unlike many other languages, in Go single quotes and double quotes are not interchangeable. 
    
    - Rune literals can be written as single Unicode characters (**'a'**)
    
    - There are also several backslash escaped rune literals, with the most useful ones being newline (**'\n'**), tab (**'\t'**), single quote (**'\''**), double quote (**'\"'**), and backslash (**'\\'**)
    
    - Octal representations are rare, mostly used to represent POSIX permission flag values (such as **0o777 for rwxrwxrwx**). Hexadecimal and binary are sometimes used for bit filters or networking
  
  - **string**
  
    - Most of the time, you should use double quotes to create an interpreted string literal(eg., type **"Greetings and Salutations"**).
    
    - These contain zero or more rune literals, in any of the forms allowed. The only characters that cannot appear are unescaped backslashes, unescaped newlines, and unescaped double quotes.
  
      - If you use an interpreted string literal and want your greetings on a different line from your salutations and you want “Salutations” to appear in quotes, you need to type ****"Greetings and\n\"Salutations\""****.
    
    - If you need to include backslashes, double quotes, or newlines in your string, use a raw string literal. These are delimited with backquotes (**`**) and can contain any literal character except a **backquote**. When using a raw string literal, we write our *multiline* greeting like so:

    ```go
        `Greetings and
        "Salutations"` 
    ```
  

- Go lets you use an integer literal in floating point expressions or even assign an integer literal to a floating point variable. 

  - This is because literals in Go are untyped; they can interact with any variable that’s compatible with the literal

