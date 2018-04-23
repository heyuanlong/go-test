### 1.download protobuf 3.x
### 2.download cmake and use cmake-gui generate vs-project
### 3.use vs generate protoc.exe and *.dll
### 4. go get -a github.com/golang/protobuf/protoc-gen-go
### 5. protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto
### 6. cd greeter_server && go run main.go
### 7. cd greeter_client && go run main.go


golang gRPC 安装教程
https://studygolang.com/articles/4370 教程

1.下载protoc  https://github.com/google/protobuf/releases
2.go get -u github.com/golang/protobuf/protoc-gen-go #golang 插件
3.https://github.com/grpc/grpc-go/tree/v1.11.3

4.https://github.com/golang/text/tree/v0.3.0
5.https://github.com/google/go-genproto

6.protoc --go_out=plugins=grpc:. grpc.proto




