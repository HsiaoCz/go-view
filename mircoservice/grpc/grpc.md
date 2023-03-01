## grpc

grpc是一个高性能的，开源的rpc框架
使用HTTP2.0，可以认为这个所谓的2.0就是对tcp的简单封装但是又不损失它的性能

grpc是跨语言的,它使用了高序列化的protobuf协议

protobuf是一种数据存储格式

```protobuf
syntax="proto3"; // 指定protobuf的版本

package hello;  // 指定生成的包名

option go_package="./;hello" ; // 指定生成的go代码的包名

message Hello{
   string name=1;
   string password=2;
}
`
生成命令
```go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
```
这里需要注意的是，生成命令的时候需要在项目的根目录下
