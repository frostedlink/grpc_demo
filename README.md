# grpc_demo
gRPC server built in go for demonstration purposes 

## Requirements

Install protobuf compiler

```
sudo apt install -y protobuf-compiler
```

```
go install
```

## Create protobuf

```
protoc --go_out=plugins=grpc:chat chat.proto
```

### References:

[gRPC go beginners guide](https://www.youtube.com/watch?v=BdzYdN_Zd9Q)