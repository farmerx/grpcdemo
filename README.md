# grpcdemo
------

## gRPC简介
> gRPC是由Google主导开发的RPC框架，使用HTTP/2协议并用ProtoBuf作为序列化工具。其客户端提供Objective-C、Java接口，服务器侧则有Java、Golang、C++等接口，从而为移动端（iOS/Androi）到服务器端通讯提供了一种解决方案。 当然在当下的环境下，这种解决方案更热门的方式是RESTFull API接口。该方式需要自己去选择编码方式、服务器架构、自己搭建框架（JSON-RPC）。gRPC官方对REST的声音是：

* 和REST一样遵循HTTP协议(明确的说是HTTP/2)，但是gRPC提供了全双工流
* 和传统的REST不同的是gRPC使用了静态路径，从而提高性能
* 用一些格式化的错误码代替了HTTP的状态码更好的标示错误
> 至于是否要选择用gRPC。对于已经有一套方案的团队，可以参考下。如果是从头来做，可以考虑下gRPC提供的从客户端到服务器的整套解决方案，这样不用客户端去实现http的请求会话，JSON等的解析，服务器端也有现成的框架用。从15年3月到现在gRPC也发展了一年了，慢慢趋于成熟。下面我们就以gRPC的Golang版本看下其在golang上面的表现。至于服务端的RPC，感觉golang标准库的RPC框架基本够用了,没必要再去用另一套方案。

## 安装protobuf

* 安装 protoc ：[Protoc下载地址](https://github.com/google/protobuf/releases)，可以根据自己的系统下载相应的 protoc，windows 用户统一下载 win32 版本。

* 配置 protoc 到系统的环境变量中，执行如下命令查看是否安装成功：
  ```
  $ protoc --version
  ```
如果正常打印 libprotoc 的版本信息就表明 protoc 安装成功

* 安装 ProtoBuf 相关的 golang 依赖库
  ```
  $ go get -u github.com/golang/protobuf/{protoc-gen-go,proto}
  ```



