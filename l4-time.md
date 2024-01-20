### Time

#### Time Duration
- Go’s standard library includes time support, found in the **time package**.
  
- There are **two main types used to represent time, time.Duration and time.Time**.

- A **period of time is represented with a time.Duration, a type based on an int64**. 

- **Time package defines constants of type time.Duration to represent a nanosecond, microsecond, millisecond, second, minute, and hour**. 

##### Creation of Durations:

- You can create a time.Duration using the time.ParseDuration function or by using literalswith a time unit suffix (e.g., time.Second, time.Minute, etc.).

  - A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as **“300ms”, “-1.5h” or “2h45m”**. **Valid time units are “ns”, “us” (or “μs”), “ms”, “s”, “m”, “h”**.

```go
    import "time"

    duration1, _ := time.ParseDuration("2h30m")
    duration2 := 5 * time.Minute
```

##### Methods:

- Time.Duration has various methods that allow you to perform arithmetic operations, such as addition, subtraction, and scaling.

```go
    duration := 2 * time.Second
    newDuration := duration + 500*time.Millisecond
```
##### Conversion:

- You can convert a time.Duration to its equivalent in other time units (e.g., milliseconds, seconds) using methods like Seconds(), Milliseconds(), etc.

```go
    duration := 1500 * time.Millisecond
    seconds := duration.Seconds()
```
##### String Representation:

A time.Duration can be converted to a string using the String() method. This is useful for logging or displaying durations.

```go
    duration := 3 * time.Hour + 30 * time.Minute
    fmt.Println(duration.String()) // Output: 3h30m
```










- In Go, the time.Time type is a built-in type representing a point in time. 

- It's part of the time package and is used to represent dates, times, and timestamps. 

- Instances of time.Time capture both the date and the time of day, along with information about the time zone.



- It has various fields and methods for working with dates, times, and durations. Some of the important aspects and features of time.Time are as follows:

##### Creation of Time Instances:

- You can create a time.Time instance representing the current time using time.Now(). 
  
- You can also create instances representing specific dates and times using the time.Date function.

```go
    import "time"

    now := time.Now()
    specificTime := time.Date(2022, time.May, 1, 12, 0, 0, 0, time.UTC)
```


##### Formatting and Parsing:

- You can **format a time.Time instance as a string using the Format method or the time.Format function**. 

- Conversely, **you can parse a string into a time.Time instance using the time.Parse or time.ParseInLocation functions**.

- it uses its own date and **time formatting language**.
   
  - It **relies on the idea of formatting the date and time ```January 2, 2006 at 3:04:05PM MST (Mountain Standard Time)``` to specify your format**.
  
  - **Why that date? Because each part of it represents one of the numbers from 1 to 7 in sequence, that is, 01/02 03:04:05PM ’06 -0700 (MST is 7 hours before UTC)**

```go
    import "time"

    now := time.Now()
    formatted := now.Format("2006-01-02 15:04:05")
    parsed, _ := time.Parse("2006-01-02 15:04:05", "2022-01-09 12:30:00")
```

##### Arithmetic Operations:

You can **perform arithmetic operations with time.Time instances, such as adding or subtracting durations to get a new time.Time**.

```go
    import "time"

    now := time.Now()
    futureTime := now.Add(2 * time.Hour)
    pastTime := now.Add(-1 * time.Day)
```

###### Comparison:

- The fact that a time.Time instance contains a time zone means that **you should not use == to check if two time.Time instances refer to the same moment in time. Instead, use the Equal Before method, which corrects for time zone**.

```go
import "time"

time1 := time.Now()
time2 := time.Now().Add(2 * time.Hour)

if time1.Before(time2) {
    // Do something...
}
```

###### Time Zones:

- time.Time instances carry information about the time zone in which they were created. You **can convert times between time zones using the time.Time.In method or time.Time.UTC** method.

```go
    import "time"

    now := time.Now()
    utcTime := now.UTC()
```


###### Duration Since/Until:

- You can **calculate the duration between two time.Time instances using the time.Since or time.Until functions**.

```go
    import "time"

    startTime := time.Now()
    // Some operation...
    elapsedTime := time.Since(startTime)
```

- time.Time is an essential type for working with time-related functionality in Go programs. It is versatile and provides a comprehensive set of methods for various time-related operations.

##### Timer:

- A time.Timer represents a single **event that will happen after a specified duration**. 

- You can use **a timer to execute a function or signal an event after a certain period**.


```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        // Create a timer that will expire after 2 seconds
        timer := time.NewTimer(2 * time.Second)

        fmt.Println("Waiting for the timer to expire...")
        <-timer.C // Blocking until the timer expires

        fmt.Println("Timer expired!")
    }
```

###### Timeout:

- A **timeout is a duration within which an operation is expected to complete**. 
- **If the operation doesn't complete within the specified time, it is considered as having timed out**.

```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        // Create a channel to signal the end of an operation
        done := make(chan bool)

        // Perform an operation that may take some time
        go func() {
            time.Sleep(3 * time.Second)
            done <- true
        }()

        // Set a timeout for the operation
        select {
        case <-done:
            fmt.Println("Operation completed successfully.")
        case <-time.After(2 * time.Second):
            fmt.Println("Operation timed out.")
        }
    }
```

###### time.AfterFunc:

- **```time.AfterFunc``` creates a new timer that will call a function after a specified duration**.

- **It's useful when you want to execute a function asynchronously after a certain period**.

```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        timer := time.AfterFunc(2*time.Second, func() {
            fmt.Println("Function executed after 2 seconds.")
        })

        // Block the main goroutine to allow the timer to execute
        // In a real program, you may use channels or other synchronization mechanisms.
        select {}
    }
```

###### time.NewTicker:

- **time.NewTicker returns a new Ticker that will send the current time on its channel at regular intervals**.
  
- **You have more control over the Ticker, allowing you to stop it explicitly when needed**.


```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        ticker := time.NewTicker(1 * time.Second)
        defer ticker.Stop() // Ensure the Ticker is stopped to release resources

        for {
            <-ticker.C
            fmt.Println("Tick at", time.Now())
        }
    }
```