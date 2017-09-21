### 1.download protobuf 3.x
### 2.download cmake and use cmake-gui generate vs-project
### 3.use vs generate protoc.exe and *.dll
### 4. go get -a github.com/golang/protobuf/protoc-gen-go
### 5. protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto