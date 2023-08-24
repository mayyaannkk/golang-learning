# D ?

When we use http.ListenAndServe() it takes 2 arguments:

1. Patter - as string
2. Handler - of type Handler

---

## What's inside this Handler?

Handler is just an interface that has ServeHTTP method in it.

```Go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

So every type that implements this SeverHTTP method becomes an implementation of Handler type and then can be used in place of Handler.

Which also means that we can pass our own type inside ListenAndServe() if we implement ServeHTTP() method.

```Go
type Router struct{}

func (router Router) ServeHTTP(w ResponseWriter, r *Request) {
    switch r.URL.Path {
        case "/":
            homeHandler(w, r)
        case "/contact":
            contacthandler(w, r)
        default:
            http.Error(w, "Not Found", http.StatusNotFound)
    }
}

func main() {
    var router Router
    fmt.Println("Starting server on 5000")
    http.ListenAndServe(":5000", router)
}
```

We can then just pass this Router in our ListenAndServe() as second argument. And we just made our own router.

---

## HandlerFunc

It is just used for type conversion from a normal function to a Handler. From the golang docs:

> If `f` is a function with appropriate signature, `HandlerFunc(f)` is a Handler that calls `f`.

---

## Handle

http.Handle() also takes a pattern as string and a Handler that gets invoked when this pattern is found.

Internally this http.Handle() calls DefaultServeMux's http.Handle() function and just passes on the arguments that we sent

---

## HandleFunc

It takes a **pattern string** and a **normal function** as argument.

It then internally calls the DefaultServeMux.HandleFunc(pattern, handler), forwarding the arguments that we passed.

And then it uses HandlerFunc to convert our normal function into a Handler.

---
