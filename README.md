# Install
1. Download [latest release of `protobuf`](https://github.com/protocolbuffers/protobuf/releases). (Note: `proto/include` in this repo is copied from `protobuf`.)
2. Install `grpc`: `go get -u google.golang.org/grpc`.
3. Install `protoc-gen-go`: `go get -u github.com/golang/protobuf/protoc-gen-go`.
4. Regenerate `service.pb.go`: `protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto`.
5. Install `gin`: `go get -u github.com/gin-gonic/gin`.

# Run
1. Run server: `go run server/main.go`.
2. Run client: `go run client/main.go`.
3. Test api endpoints:
    - `curl http://localhost:4002/calculate/4/+/2`.
    - `curl http://localhost:4002/calculate/4/-/2`.
    - `curl http://localhost:4002/calculate/4/*/2`.
    - `curl http://localhost:4002/calculate/4/divide/2`.
4. Test errors:
    - `curl http://localhost:4002/calculate/4/asdf/2` (bad operator).

# Links
- Building an Basic API with gRPC and Protobuf by Tensor Programming ([video](https://www.youtube.com/watch?v=Y92WWaZJl24), [repo](https://github.com/tensor-programming/grpc_tutorial))
