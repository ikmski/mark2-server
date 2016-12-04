# Protocol Buffers for Mark2

## Protocol Buffers v3
https://developers.google.com/protocol-buffers/

## コンパイル方法

```
// GO 用にコンパイルする
$ protoc -I <INPUT DIR> --go_out=plugins=grpc:<OUTPUT DIR> <INPUT PROTO>

// exp)
$ protoc -I ./ --go_out=plugins=grpc:./ ./mark2.proto
```

## 生成ファイル

* mark2.pb.go
