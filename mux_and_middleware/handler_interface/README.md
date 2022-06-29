# Handler interface

We need to implement interface:
```go
type Handler interface{
    ServeHTTP(ResponseWriter, *Request)
}
```

To achieve that we need to provide our handler struct with single method:

`ServeHTTP(w http.ResponseWriter, r *http.Request)`
