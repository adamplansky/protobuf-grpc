# TODO
simple app for demonstrate protobuf + grpc written in golang

[youtube tutorial](https://www.youtube.com/watch?v=uolTUtioIrc&t=6s)


## generate pb.go with rpc
```sh
protoc -I . todo.proto --go_out=plugins=grpc:.
```

## how to run 
```sh
#run server (port 8888)
go run ./cmd/server/main.go

# use client
go run ./cmd/client/main.go add hello 
go run ./cmd/client/main.go list
```

