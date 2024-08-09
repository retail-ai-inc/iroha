# IROHA

## 1.About gRPC

1. The Protocol Buffers file in:

    ```proto/cryptographic.proto```

2. Generate files that support Golang

    ```shell
    protoc --go_out=. --go-grpc_out=. cryptographic.proto
    ```
