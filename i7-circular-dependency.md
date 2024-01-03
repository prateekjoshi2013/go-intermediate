### Circular Dependencies

- Go **does not allow you to have a circular dependency between packages**.
  
- This means that **if package A imports package B, directly or indirectly, package B cannot import package A, directly or indirectly**

### Solutions

- If you find yourself with a circular dependency, you have a few options. 
  
  - In some cases, **this is caused by splitting packages up too finely**. 
    
    - **If two packages depend on each other, there’s a good chance they should be merged into a single package**

  - **If you have a good reason to keep your packages separated, it may be possible to move just the items that cause the circular dependency to one of the two packages or to a new package**