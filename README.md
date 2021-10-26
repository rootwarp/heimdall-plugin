# Heimdall Plugin for Node Manager

Build plugin.

```
~$ make build
```

List APIs by using reflection.

```
~$ grpcurl -plaintext localhost:9090 list nodemanager.v1.NodeManager
```

Install Go plugins for protobuf.

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

And open path

https://grpc.io/docs/languages/go/quickstart/
