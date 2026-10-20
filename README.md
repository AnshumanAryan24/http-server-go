# http-server-go
Trying to create a simple HTTP server in Golang.

Commits:
    - Prefix: code/test/doc (join using ",")
    - Message
    - Separator: " | "

This works for now:

```shell
curl --http0.9 http://localhost:8080 --data-ascii "hello world"
```

Response to curl request: `sample server⏎`
Read data:
```text
POST / HTTP/1.1
Host: localhost:8080
User-Agent: curl/8.22.0
Accept: */*
Content-Length: 11
Content-Type: application/x-www-form-urlencoded

hello world
```

Read 159 bytes