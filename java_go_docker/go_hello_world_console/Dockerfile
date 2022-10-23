FROM golang
WORKDIR /app
ADD go.* /app/
ADD *.go /app/
RUN go_hello_world_console build -o main .
ENTRYPOINT ["/app/main"]
