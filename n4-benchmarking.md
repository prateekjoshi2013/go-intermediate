- ```benchmarking``` support that’s built into Go’s testing framework.

- benchmarks are functions in your test files that start with the word Benchmark and take in a single parameter of type *testing.B. 

- This type includes all of the functionality of a *testing.T as well as additional support for benchmarking. 

- Let’s start by looking at a benchmark that uses a buffer size of 1 byte:

    ```go
        var blackhole int

        func BenchmarkFileLen1(b *testing.B) {
            for i := 0; i < b.N; i++ {
                result, err := FileLen("testdata/data.txt", 1)
                if err != nil {
                    b.Fatal(err)
                }
                blackhole = result
            }
        }
    ```

- The ```blackhole package-level variable``` is interesting. We **write the results from FileLen to this package-level variable to make sure that the compiler doesn’t get too clever and decide to optimize away the call to FileLen, ruining our benchmark**.

- Every **Go benchmark must have a loop that iterates from 0 to b.N**. 

- The **testing framework calls our benchmark functions over and over with larger and larger values for N until it is sure that the timing results are accurate**. 

- **We run a benchmark by passing the ``-bench flag`` to go test. This flag ```expects a regular expression``` to describe the name of the benchmarks to run**. 

- **Use ```-bench=.``` to run all benchmarks**. 

- **A second flag, ```-benchmem```, includes ```memory allocation``` information in the benchmark output**. 

- All **tests are run before the benchmarks**, so you **can only benchmark code when tests pass**.

- Here’s the **output for the benchmark on my computer**:

```sh
    BenchmarkFileLen1-12 25 47201025 ns/op 65342 B/op 65208 allocs/op
```
    - Output means following:

    - BenchmarkFileLen1-12
      
      - The name of the benchmark, a hyphen, and the value of GOMAXPROCS for the benchmark.

    - 25
      
      - The number of times that the test ran to produce a stable result.

    - 47201025 ns/op
      
      - How long it took to run a single pass of this benchmark, in nanoseconds (there are 1,000,000,000 nanoseconds in a second).

    - 65342 B/op
      
      - The number of bytes allocated during a single pass of the benchmark

    - 65208 allocs/op
      
      - The number of times bytes had to be allocated from the heap during a single pass of the benchmark. This will always be less than or equal to the number of bytes allocated.


- Now that we have results for a buffer of 1 byte, let’s see what the results look like when we **use buffers of different sizes**:
  
```go
    func BenchmarkFileLen(b *testing.B) {
        for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
                b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
                for i := 0; i < b.N; i++ {
                    result, err := FileLen("testdata/data.txt", v)
                    if err != nil {
                        b.Fatal(err)
                    }
                    blackhole = result
                }
            })
        }
    }
```

- We **launched table tests using t.Run**, we’re using **b.Run to launch benchmarks that only vary based on input**.
