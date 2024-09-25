# grpc-client-streaming-test

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```
.
├── README.md
├── scantext
│   └── scantext.proto
└── server
    └── main.go
```

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    scantext/scantext.proto
```

```
$ go mod tidy
```
