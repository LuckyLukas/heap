#heap

a simple implementation of a min heap in go. No additional allocations when adding nodes
```
BenchmarkHeap_AddNoCap-2                         5279101               204 ns/op              87 B/op          0 allocs/op
BenchmarkHeap_AddOverCapacity-2                 10372914               250 ns/op              89 B/op          0 allocs/op
BenchmarkHeap_RemoveTopGrowingTree-2             1231784               995 ns/op             105 B/op          1 allocs/op
BenchmarkHeap_RemoveTopStagnantTree-2            4079859               303 ns/op               8 B/op          1 allocs/op
```
