## mux这是一个路由库

**安装**

```go
go get -u github.com/gorilla/gorilla/mux
```

**mux使用**

```go
func main() {
// 创建一个类型为*mux.Router的路由对象
// 这个路由对象注册处理器的方式与标准库是一样的
  r := mux.NewRouter()
  r.HandleFunc("/", BooksHandler)
//   /book/{icbc} 这种方式可以匹配路径中的特定部分
  r.HandleFunc("/books/{isbn}", BookHandler)
// *mux.Router也实现了http.Handler接口，可以把它作为请求的处理器
 http.Handle("/", r)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**mux有着灵活的匹配方式**

```go
// 指定路由的域名和子域名
r.Host("github.io")
r.Host("{subdomain:[a-zA-Z0-9]+}.github.io")

// 指定路径的前缀
// 只处理路径前缀为`/books/`的请求
r.PathPrefix("/books/")

// 指定请求的方法
// 只处理 GET/POST 请求
r.Methods("GET", "POST")

// 指定使用的协议
// 只处理 https 的请求
r.Schemes("https")

// 指定请求首部
// 只处理首部 X-Requested-With 的值为 XMLHTTPRequest 的请求
r.Headers("X-Requested-With", "XMLHTTPRequest")

// 指定查询的参数
// 查询参数（即 URL 中?后的部分）：
// 只处理查询参数包含key=value的请求
r.Queries("key", "value")

// 组合上述内容使用基本上就是这样

r.HandleFunc("/", HomeHandler)
 .Host("bookstore.com")
 .Methods("GET")
 .Schemes("http")
```

mux支持自定义匹配器
自定义的匹配器就是一个类型为func(r *http.Request, rm *RouteMatch) bool的函数，根据请求r中的信息判断是否能否匹配成功

http.Request结构中包含了非常多的信息：HTTP 方法、HTTP 版本号、URL、首部等。例如，如果我们要求只处理 HTTP/1.1 的请求可以这么写：

```go
r.MatchrFunc(func(r *http.Request, rm *RouteMatch) bool {
  return r.ProtoMajor == 1 && r.ProtoMinor == 1
})
```

需要注意的是,mux会根据路由注册的顺序依次匹配，所以，通常是将特殊的路由放在前面，一般的路由放在后面，如果反过来，特殊的路由就不会被匹配到了
```go
r.HandleFunc("/specific", specificHandler)
r.PathPrefix("/").Handler(catchAllHandler)
```

**mux子路由**

```go
r := mux.NewRouter()
bs := r.PathPrefix("/books").Subrouter()
bs.HandleFunc("/", BooksHandler)
bs.HandleFunc("/{isbn}", BookHandler)

ms := r.PathPrefix("/movies").Subrouter()
ms.HandleFunc("/", MoviesHandler)
ms.HandleFunc("/{imdb}", MovieHandler)
```

子路由一般通过路径前缀来限定，r.PathPrefix()会返回一个*mux.Route对象，调用它的Subrouter()方法创建一个子路由对象*mux.Router，然后通过该对象的HandleFunc/Handle方法注册处理函数。

[https://www.cnblogs.com/jiujuan/p/12768907.html]