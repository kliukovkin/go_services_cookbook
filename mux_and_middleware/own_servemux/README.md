# Own ServeMux

In production, it is not a good pattern 
to use default serveMux because it is a 
global variable thus any package can access
it and register a new router or something
worse. So let’s create our own serveMux.
To do this we will use `http.NewServeMux`
function.

Because `http.http.ListenAndServe` is waiting for
the Handler as a second argument, we may right
our own Handler with custom logic. Take a look
at handwritten_handler.go file.

