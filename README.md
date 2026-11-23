# Go In-Memory Cache

[简体中文](README.zh-CN.md)

An in-process key-value cache written in Go. The demo uses a small cache API to set and retrieve Go values, delete entries, count keys, and clear the cache.

## Implementation

- [cache](cache) stores entries in a map and tracks an estimated size for each value. It accepts a configurable size limit and a TTL argument. In the current implementation, entries written with a positive TTL can be treated as expired immediately.
- [cache_server](cache_server) wraps the in-memory implementation as a Go API. Set a size limit on each instance before writing values. The current implementation terminates the process on overflow.
- [test](test) contains exercises for value retrieval, expiry and size limits.

## Run the demo

The [Go module](go.mod) declares Go 1.19. Run the demonstration from the repository root.

```sh
go run .
```
