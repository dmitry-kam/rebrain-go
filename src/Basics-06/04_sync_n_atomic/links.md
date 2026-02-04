# Sync.Map
The sync package includes a special type — Map — which implements a key-value storage with a built-in 
mutex.
It can be used instead of the map + sync.Mutex combination, but there are important nuances:
- sync.Map uses interfaces (interface{}) for both keys and values, so type assertions will be required.
- sync.Map is only more efficient than map + sync.Mutex starting from a certain number of reads per second
and with at least 4 CPU cores.

# Atomic
The atomic package (sync/atomic), which forms the foundation for much of the sync package, 
provides low-level primitives for memory operations.
This package mainly works with integers and pointers to them.

Except for specific low‑level programming cases, it is recommended to not use the atomic
package directly, but rather the components of the sync package or channels, since they are more optimized and designed for this purpose.

# Links
- https://pkg.go.dev/sync
- https://pkg.go.dev/sync/atomic
- https://habr.com/ru/articles/338718/