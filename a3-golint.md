
- Linter tries to ensure your code follows style guidelines. 

- Some of the changes it suggests
  - include properly naming variables, 
  - formatting error messages, and 
  - placing comments on public methods and types. 

- These aren’t errors; they don’t keep your programs from compiling or make your program run incorrectly. 

- Also, you cannot automatically assume that golint is 100% accurate: because the kinds of issues that golint finds are more fuzzy, it sometimes has false positives and false negatives. 
- This means that you don’t have to make the changes that golint suggests. But you should take the suggestions from golint seriously. 

- Go developers expect code to look a certain way and follow certain rules, and if your code does not, it sticks out.


```sh
    go install golang.org/x/lint/golint@latest
```

- Run the linter
  
```sh
    golint
```