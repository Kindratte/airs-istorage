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
BenchmarkChanValue-4              	 3000000	       445 ns/op	      48 B/op	       1 allocs/op
BenchmarkChanValueFromPointer-4   	 3000000	       438 ns/op	      48 B/op	       1 allocs/op
BenchmarkChanPointer-4            	 3000000	       480 ns/op	      80 B/op	       2 allocs/op
```

# Slice or Map

Slice is faster for small sizes
https://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
