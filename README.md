# Go In-Memory Cache

[简体中文](README.zh-CN.md)

An in-process Go key-value cache with a small API for writing, reading, deleting and counting entries or clearing the store. The implementation is an exercise in tracking entry size and expiry.

## Implementation

- [cache](cache) stores entries in a map and tracks an estimated size for each value. It accepts a configurable size limit and a TTL argument. In the current implementation, entries written with a positive TTL can be treated as expired immediately.
- [cache_server](cache_server) wraps the in-memory implementation as a Go API. Set a size limit on each instance before writing values. The current implementation terminates the process on overflow.
- [test](test) contains exercises for value retrieval, expiry and size limits.

## Run the demo

The [Go module](go.mod) declares Go 1.19. Run the demonstration from the repository root.

```sh
go run .
```
