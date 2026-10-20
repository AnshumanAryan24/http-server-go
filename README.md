# http-server-go
Trying to create a simple HTTP server in Golang.

Commits:
    - Prefix: code/test/doc (join using ",")
    - Message
    - Separator: " | "

This works for now:

```shell
curl --http1.1 http://localhost:8080 --data-ascii "hello world"
```

Response to curl request: `Hello World⏎`
Read data:
```text
HTTP 1.1 server
2026/10/20 22:02:16 listener.listenForHttpTcp mode "ACCEPT ALL" :: connected: 127.0.0.1:49462
POST / HTTP/1.1
Accept:*/*
Content-Length:11
Content-Type:application/x-www-form-urlencoded
Host:localhost:8080
User-Agent:curl/8.22.0
```

Refering RFC 9112 & 9110 - not inncluding caching for now.