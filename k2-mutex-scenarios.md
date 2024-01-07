### When to Use Mutexes Instead of Channels

- **To coordinate access to data across threads in other programming languages, you have probably used a ```mutex```**. 

- This is **short for mutual exclusion**, and **the job of a mutex is to limit the concurrent execution of some code or access to a shared piece of data**. 

- This **protected part is called the ```critical section```**.

- There is a saying in the Go community to describe this philosophy: 
  
  - “**Share memory by communicating; do not communicate by sharing memory**.”

- The most **common case is when your goroutines read or write a shared value, but don’t process the value**. 

  - Let’s use an **inmemory scoreboard for a multiplayer game as an example**.

  - Here’s a function that we can launch as a goroutine to manage the scoreboard:

```go
    func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
        scoreboard := map[string]int{}
        for {
            
            select {
            
            case <-done:
                return
            
            case f := <-in:
                f(scoreboard)
            }
        }
    }
```

- This function declares a map and then listens on one channel for a function that reads or modifies the map and on a second channel to know when to shut down. 

- Let’s create a type with a method to write a value to the map:

```go
    type ChannelScoreboardManager chan func(map[string]int)

    func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
        ch := make(ChannelScoreboardManager)
        done := make(chan struct{})
        go scoreboardManager(ch, done)
        return ch, func() {
            close(done)
        }
    }

    func (csm ChannelScoreboardManager) Update(name string, val int) {
        csm <- func(m map[string]int) {
            m[name] = val
        }
    }
```

- The update method is very straightforward: just pass a function that puts a value into the map. 

- But how about reading from the scoreboard? We need to return a value back. 
- That means using the done pattern to wait for the function passed to the ScoreboardManager to finish running:

```go
    func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
        var out int
        var ok bool
        done := make(chan struct{})
        
        csm <- func(m map[string]int) {
            out, ok = m[name]
            close(done)
        }
        
        <-done
        
        return out, ok
    }
```

#### Using Mutex

- *While this code works, it’s cumbersome and **only allows a single reader** at a time*. 

- A **better approach is to use a mutex**. 

- There are *two mutex implementations in the standard library, both in the **sync** package*. 

- The *first is called **Mutex** and has **two methods**, **Lock** and **Unlock***. 
  
  - **Calling Lock causes the current goroutine to pause as long as another goroutine is currently in the critical section**. 
  
  - **When the critical section is clear, the lock is acquired by the current goroutine and the code in the critical section is executed**. 
  
  - **A call to the Unlock method on the Mutex marks the end of the critical section**.

- The **second mutex implementation is called RWMutex** and it **allows you to have both reader locks and writer locks**. 
  
  - While **only one writer can be in the critical section at a time, reader locks are shared multiple readers can be in the critical section at once**.
  
  - The **writer lock is managed with the Lock and Unlock methods, while the reader lock is managed with RLock and RUnlock methods**.
  
  - **Multiple goroutines can acquire a read lock (RLock()), allowing them to read the shared resource concurrently**. 
  
  - **However, when a write lock (Lock()) is acquired, it excludes all other readers and writers until the write lock is released using Unlock()**.

  - It is **suitable for scenarios where there are many read operations and occasional write operations, allowing for better parallelism**.

- **Any time you acquire a mutex lock, you must make sure that you release the lock. Use a defer statement to call Unlock immediately after calling Lock or RLock**

  

```go
    type MutexScoreboardManager struct {
        l sync.RWMutex
        scoreboard map[string]int
    }

    func NewMutexScoreboardManager() *MutexScoreboardManager {
        return &MutexScoreboardManager{
            scoreboard: map[string]int{},
        }
    }

    func (msm *MutexScoreboardManager) Update(name string, val int) {
        msm.l.Lock()
        defer msm.l.Unlock()
        msm.scoreboard[name] = val
    }

    func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
        msm.l.RLock()
        defer msm.l.RUnlock()
        val, ok := msm.scoreboard[name]
        return val, ok
    }
```

### Caution using Mutex

- You must **correctly pair locks and unlocks or your programs will likely deadlock**

- Another issue is that **mutexes in Go aren’t reentrant**. 
  
  - **If a goroutine tries to acquire the same lock twice, it deadlocks, waiting for itself to release the lock**. 
  
  - This is different from languages like Java, where locks are reentrant.
  
  - Nonreentrant locks make it tricky to acquire a lock in a function that calls itself recursively. **You must release the lock before the recursive function call**

- Like sync.WaitGroup and sync.Once, mutexs must never be copied. **If they are passed to a function or accessed as a field on a struct, it must be via a pointer. If a mutex is copied, its lock won’t be shared**

### to help you decide whether to use channels or mutexes:

- **If you are coordinating goroutines or tracking a value as it is transformed by a series of goroutines**, use **channels**.

- **If you are sharing access to a field in a struct**, use **mutexes**.

- If you discover **a critical performance issue when using channels and you cannot find any other way to fix the issue**, **modify your code to use a mutex**.

### sync.Map—This Is Not the Map You Are Looking For

- When looking through the sync package, you’ll find a type called Map. 

- It provides a concurrency-safe version of Go’s built-in map. Due to trade-offs in its implementation

- **sync.Map is only appropriate in very specific situations**:

  - **When you have a shared map where key/value pairs are inserted once and read many times**

  - **When goroutines share the map, but don’t access each other’s keys and values**

- **Given these limitations, in the rare situations where you need to share a map across multiple goroutines, use a built-in map protected by a sync.RWMutex**.