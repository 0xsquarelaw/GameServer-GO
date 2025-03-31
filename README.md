# GameServer-GO
Golang Game Server Written in GO using Protobuf

## Some Commands
- `winget install protobuf`
- `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- [Link for gRPC guide](https://grpc.io/docs/languages/go/quickstart/)
- `protoc -I="shared" --go_out="server" "shared/packets.proto"`