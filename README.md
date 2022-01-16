### 1、 Algorithm Description
```
/ Task: Implement a struct named 'RangeList'
// A pair of integers define a range, for example: [1, 5). This range includes integers: 1, 2, 3, and 4.
// A range list is an aggregate of these ranges: [1, 5), [10, 11), [100, 201)
// NOTE: Feel free to add any extra member variables/functions you like.

````
example:

```
    rl := RangeList{}
   
    rl.Add([2]int{1, 5})
    rl.Print()
    // Should display: [1, 5)
    rl.Add([2]int{10, 20})
    rl.Print()
    // Should display: [1, 5) [10, 20)
    rl.Add([2]int{20, 20})
    rl.Print()
    // Should display: [1, 5) [10, 20)
    rl.Add([2]int{20, 21})
    rl.Print()
    // Should display: [1, 5) [10, 21)
    rl.Add([2]int{2, 4})
    rl.Print()
    // Should display: [1, 5) [10, 21)
    rl.Add([2]int{3, 8})
    rl.Print()
    // Should display: [1, 8) [10, 21)
    rl.Remove([2]int{10, 10})
    rl.Print()
    // Should display: [1, 8) [10, 21)
    rl.Remove([2]int{10, 11})
    rl.Print()
    // Should display: [1, 8) [11, 21)
    rl.Remove([2]int{15, 17})
    rl.Print()
    // Should display: [1, 8) [11, 15) [17, 21)
    rl.Remove([2]int{3, 19})
    rl.Print()
    // Should display: [1, 3) [19, 21)


```
### 2、Implements Result
```
$ go run range_list.go
[1, 5)
[1, 5) [10, 20)
[1, 5) [10, 20)
[1, 5) [10, 21)
[1, 5) [10, 21)
[1, 8) [10, 21)
[1, 8) [10, 21)
[1, 8) [11, 21)
[1, 8) [11, 15) [17, 21)
[1, 3) [19, 21)

```

### 3、Benchmarks

```
$ go test -v -benchmem  -benchtime=1s  -bench=.
goos: windows
goarch: amd64
pkg: github.com/jinri/jinri-algorithm
cpu: Intel(R) Core(TM) i5-1035G7 CPU @ 1.20GHz
BenchmarkAdd
BenchmarkAdd-8          17922484                61.18 ns/op            7 B/op          0 allocs/op
BenchmarkRemove
BenchmarkRemove-8       17300688                69.40 ns/op            8 B/op          0 allocs/op

BenchmarkQuery
BenchmarkQuery-8        27243068                42.99 ns/op            0 B/op          0 allocs/op
BenchmarkPrint
BenchmarkPrint-8        1000000000             0B/op          0 allocs/op
PASS

ok      github.com/jinri/jinri-algorithm        3.691s


$ go test -v -benchmem  -benchtime=2s  -bench=.
goos: windows
goarch: amd64
pkg: github.com/jinri/jinri-algorithm
cpu: Intel(R) Core(TM) i5-1035G7 CPU @ 1.20GHz
BenchmarkAdd
BenchmarkAdd-8          51518065                62.91 ns/op            8 B/op          0 allocs/op
BenchmarkRemove
BenchmarkRemove-8       39625669                67.44 ns/op            7 B/op          0 allocs/op
BenchmarkQuery
BenchmarkQuery-8        41223442                50.85 ns/op            0 B/op          0 allocs/op
BenchmarkPrint
BenchmarkPrint-8        1000000000             0 B/op          0 allocs/op
PASS
ok      github.com/jinri/jinri-algorithm        8.240s


$ go test -v -benchmem  -benchtime=4s  -bench=.
goos: windows
goarch: amd64
pkg: github.com/jinri/jinri-algorithm
cpu: Intel(R) Core(TM) i5-1035G7 CPU @ 1.20GHz
BenchmarkAdd
BenchmarkAdd-8          94476657                67.05 ns/op            7 B/op          0 allocs/op
BenchmarkRemove
BenchmarkRemove-8       97792736                46.72 ns/op            0 B/op          0 allocs/op
BenchmarkQuery
BenchmarkQuery-8        100000000               43.25 ns/op            0 B/op          0 allocs/op
BenchmarkPrint
BenchmarkPrint-8        1000000000             0 B/op          0 allocs/op
PASS
ok      github.com/jinri/jinri-algorithm        15.415s

```
- -8：the GOMAXPROCS value
- -1000000000: number of times
- -64.71 every run take ns
- -7 B/op every run take byte
- -0 allocs/op alloc mem time
