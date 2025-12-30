# Go Cheatsheet

### Additional References

- nice to read

  - [go-cookbook](https://go-cookbook.com/)
  - [go-bascic](https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/)
- optional

  - [interview - 1](https://medium.com/@maskwork77.dev/golang-code-interview-record-72f9fe32f7e3)
  - [interview - 2](https://blog.csdn.net/2401_87300302/article/details/142428512)
  - [interview - 3](https://www.zhihu.com/tardis/bd/art/519979757)

### Some Notes for Live Coding Interview

#### 2. concurrency & 3. race condition

- use thread-safe packages if possble, like sync.Map
- handle Map.Load/Store carefully, use LoadOrStore to prevent race conditions.

```go
sm := sync.Map{}
if v, loaded := sm.LoadOrStore(key, val); loaded {
    println("key exists!")
}
```

#### 7. error & 8. polling

- wrap an error

```go
wrapErr := fmt.Errorf("%v: %w", msg, err) // use %w for wrap error
```

- use select & ticker for error handing and polling

```go
ticker = time.NewTicker(time.Secodn)
defer ticker.Stop()

for range maxAttempts {
    select {
    case <-ticker.C:
        // do something
    case <-ctx.Done():
        return ctx.Err()
    }
}
```
