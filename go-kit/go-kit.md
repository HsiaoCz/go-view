## go-kit

go-kit的三层架构

1、Transport
主要负责与FTTP、gRPC、Thrift等的通信

2、endpoint
定义Request和Response格式，并可以使用装饰器包装函数，以此来实现各种中间件嵌套

3、Service

这就是我们的业务类、接口等


创建的时候，先创建service业务类，后创建endpoint，最后创建transport