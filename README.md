# Install
1. Download [latest release of `protobuf`](https://github.com/protocolbuffers/protobuf/releases). (Note: `proto/include` in this repo is copied from `protobuf`.)
2. Install `grpc`: `go get -u google.golang.org/grpc`.
3. Install `protoc-gen-go`: `go get -u github.com/golang/protobuf/protoc-gen-go`.
4. Regenerate `MathService.pb.go`: `protoc --proto_path=proto --go_out=plugins=grpc:proto MathService.proto`.

# Links
- Building an Basic API with gRPC and Protobuf by Tensor Programming ([video](https://www.youtube.com/watch?v=Y92WWaZJl24), [repo](https://github.com/tensor-programming/grpc_tutorial))
