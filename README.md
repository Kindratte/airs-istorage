# istorage

- Simple key-value persistency interface
- Idea borrowed from Project Kaiser storage


# Channel: `pointer` or `value`

Seems `chan RecordParts` is more effective than `chan *RecordParts`

https://groups.google.com/forum/#!topic/golang-nuts/eM_a09l8yU0
You may be surprised how large a struct can get before passing it has a noticeable performance impact compared to passing a pointer to that
struct (not to mention accessing data behind a pointer involves an indirection, and especially when the data is shared across processors,
	there can be additional cost.

```
BenchmarkChanInterface-4          	 3000000	       428 ns/op	      64 B/op	       1 allocs/op
BenchmarkChanValue-4              	 5000000	       297 ns/op	       0 B/op	       0 allocs/op
BenchmarkChanValueFromPointer-4   	 5000000	       301 ns/op	       0 B/op	       0 allocs/op
BenchmarkChanPointer-4            	 5000000	       395 ns/op	      64 B/op	       1 allocs/op
```

# Slice or Map

Slice is faster for small sizes
https://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
