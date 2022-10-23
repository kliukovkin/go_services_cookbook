FROM golang
WORKDIR /app
ADD go.* /app/
ADD *.go /app/
RUN gorilla_hello_world build -o main .
ENTRYPOINT ["/app/main"]
