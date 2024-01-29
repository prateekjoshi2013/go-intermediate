### Concurrency Problems with the Race Checker

- It’s easy to accidentally reference a variable from two different goroutines without acquiring a lock
- The computer science term for this is a data race. To help find these sorts of bugs, Go includes a race checker. 
- It isn’t guaranteed to find every single data race in your code, but if it finds one, you should put proper locks around what it finds.

```go
    func getCounter() int {
        var counter int
        var wg sync.WaitGroup
        wg.Add(5)
        for i := 0; i < 5; i++ {
            go func() {
                for i := 0; i < 1000; i++ {
                    counter++
                }
                wg.Done()
            }()
        }
        wg.Wait()
        return counter
    }
```


- This code launches five goroutines, has each of them update a shared counter variable 1000 times, and then returns the result. 

- You’d expect it to be 5000, so let’s verify this with a unit test in test_examples/race/race_test.go:

```go
    func TestGetCounter(t *testing.T) {
        counter := getCounter()
        if counter != 5000 {
            t.Error("unexpected counter:", counter)
        }
    }
```
- If you run go test a few times, you’ll see that sometimes it passes, but most of the time it fails with an error message like:

```go
    unexpected counter: 3673
```

- Let’s see what the race checker does. Use the flag -race with go test to enable it:

```sh
    $ go test -race
    ==================
    WARNING: DATA RACE
    Read at 0x00c000128070 by goroutine 10:
        test_examples/race.getCounter.func1()
            test_examples/race/race.go:12 +0x45
```