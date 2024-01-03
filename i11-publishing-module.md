### pkg.go.dev

- While **there isn’t a single centralized repository of Go modules**, there is a single service that gathers together documentation on Go modules. 

- The **Go team has created a site called pkg.go.dev that automatically indexes open source Go projects**

### Publishing Your Module

- Whether your module is public or private, **you must properly version your module** so that it works correctly with Go’s module system.
  
- **As long as you are adding functionality or patching bugs, the process is simple. Store your changes in your source code repository, then apply a tag that follows the semantic versioning rules**

- If you reach a point where you need to break backward compatibility,**Go supports two ways for creating the different import paths for creating major versions**:

  - **Create a subdirectory within your module named vN, where N is the major version of your module**. 
    
    - For example, *if you are creating version 2 of your module, call this directory v2. Copy your code into this subdirectory, including the README and LICENSE files*.
  
  - **Create a branch in your version control system. You can either put the old code on the branch or the new code**.
  
    - **Name the branch vN if you are putting the new code on the branch**
    
    - **vN-1 if you are putting the old code there**


- **Change the import path in the code in your subdirectory or branch**. 

- **The module path in your go.mod file must end with /vN, and all of the imports within your module must use /vN as well**.

- **place a tag on your repository that looks like vN.0.0**. 
  
  - **If you are using the subdirectory system or keeping the latest code on your main branch, tag the main branch**. 

  - **If you are placing your new code on a different branch, tag that branch instead**.

- More details in [Go Modules: v2 and Beyond]( https://go.dev/blog/v2-go-modules )