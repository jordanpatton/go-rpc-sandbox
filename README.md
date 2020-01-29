# Install
1. Download latest release of `protobuf` ([releases](https://github.com/protocolbuffers/protobuf/releases)). (Note: `proto/include` files in this repo are just `include` files from `protobuf`.)
2. Install `grpc`: `go get -u google.golang.org/grpc`.
3. Install `protoc-gen-go`: `go get -u github.com/golang/protobuf/protoc-gen-go`.
4. Regenerate `MathService.pb.go`: `protoc --proto_path=proto --go_out=plugins=grpc:proto MathService.proto`.

# Links
- [Building an Basic API with gRPC and Protobuf](https://www.youtube.com/watch?v=Y92WWaZJl24) by Tensor Programming ((repo)[https://github.com/tensor-programming/grpc_tutorial])
