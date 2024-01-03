- There’s no one official way to structure the Go packages in your module, but several
patterns have emerged over the years.

- When module is small, keep all of your code in a single package. As long as there are no other modules that depend on your module, there is no harm in delaying organization

- **If your module consists of one or more applications**
   
  - ***Create a directory called **cmd at the root** of your module.*** 
  - ***Within cmd, **create one directory for each binary built** from your module. ***

  - For example
    
    - ***you might have a module that contains both a web application and a command-line tool that analyzes data in the web application’s database***. 

    - ***Use **main as the package name** within each of these directories***.

- If your module’s root directory contains many files for managing the testing and deployment of your project (such as shell scripts, continuous integration configuration files, or Dockerfiles),
  
  - Place all of your Go code (besides the main packages under cmd) into packages under a directory called pkg.

- Within the pkg directory, organize your code to limit the dependencies between packages.
  
  - One common pattern is to organize your code by slices of functionality. 
    
    - For example, 
      
      - if you wrote a shopping site in Go
      
      - you might place all of the code to support customer management in one package and all of the code to manage inventory in another. This style limits the dependencies between packages, which makes it easier to later refactor a single web application into multiple microservices

    - conference talk on: [how to structure your apps] (https://www.youtube.com/watch?v=oL6JBUk6tj0&ab_channel=GopherAcademy)

