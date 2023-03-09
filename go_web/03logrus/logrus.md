## 日志库 logrus

**安装:**

```go
go get github.com/sirupsen/logrus

```

**使用**

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.SetLevel(logrus.TraceLevel)

  logrus.Trace("trace msg")
  logrus.Debug("debug msg")
  logrus.Info("info msg")
  logrus.Warn("warn msg")
  logrus.Error("error msg")
  logrus.Fatal("fatal msg")
  logrus.Panic("panic msg")
}
```

**logrus 的日志级别**

```
logrus的使用非常简单，与标准库log类似。logrus支持更多的日志级别：

Panic：记录日志，然后panic。
Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
Error：错误日志，需要查看原因；
Warn：警告信息，提醒程序员注意；
Info：关键操作，核心流程的日志；
Debug：一般程序中输出的调试信息；
Trace：很细粒度的信息，一般用不到；

日志级别从上向下依次增加，Trace最大，Panic最小。logrus有一个日志级别，高于这个级别的日志不会输出。 默认的级别为InfoLevel。所以为了能看到Trace和Debug日志，我们在main函数第一行设置日志级别为TraceLevel。
```

日志输出有三个信息：time、level、msg

**定制 logrus**

1、输出文件名

调用 logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息：

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.SetReportCaller(true)

  logrus.Info("info msg")
}
```

2、添加字段

有时候需要在输出中添加一些字段，可以通过调用 logrus.WithField 和 logrus.WithFields 实现。 logrus.WithFields 接受一个 logrus.Fields 类型的参数，其底层实际上为 map[string]interface{}：

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.WithFields(logrus.Fields{
    "name": "dj",
    "age": 18,
  }).Info("info msg")
}
```

如果在一个函数中的所有日志都需要添加某些字段，可以使用 WithFields 的返回值。例如在 Web 请求的处理器中，日志都要加上 user_id 和 ip 字段：

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  requestLogger := logrus.WithFields(logrus.Fields{
    "user_id": 10010,
    "ip":      "192.168.32.15",
  })

  requestLogger.Info("info msg")
  requestLogger.Error("error msg")
}
```

实际上，WithFields 返回一个 logrus.Entry 类型的值，它将 logrus.Logger 和设置的 logrus.Fields 保存下来。 调用 Entry 相关方法输出日志时，保存下来的 logrus.Fields 也会随之输出。

**重定向输出**

默认情况下，日志输出到 io.Stderr。可以调用 logrus.SetOutput 传入一个 io.Writer 参数。后续调用相关方法日志将写到 io.Writer 中。 现在，我们就能像上篇文章介绍 log 时一样，可以搞点事情了。传入一个 io.MultiWriter， 同时将日志写到 bytes.Buffer、标准输出和文件中：

```go
  writer1 := &bytes.Buffer{}
  writer2 := os.Stdout
  writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
  if err != nil {
    log.Fatalf("create file log.txt failed: %v", err)
  }

  logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
  logrus.Info("info msg")
```

**日志格式**

支持文本和 JSON 格式，默认为文本格式。可以通过 logurs.SetFormatter 设置日志格式

```go
package main

import (
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.SetLevel(logrus.TraceLevel)
  logrus.SetFormatter(&logrus.JSONFormatter{})

  logrus.Trace("trace msg")
  logrus.Debug("debug msg")
  logrus.Info("info msg")
  logrus.Warn("warn msg")
  logrus.Error("error msg")
  logrus.Fatal("fatal msg")
  logrus.Panic("panic msg")
}
```

**logrus 钩子**

所谓的钩子就是在执行前后执行的方法
还可以为 logrus 设置钩子，每条日志输出前都会执行钩子的特定方法。所以，我们可以添加输出字段、根据级别将日志输出到不同的目的地。 logrus 也内置了一个 syslog 的钩子，将日志输出到 syslog 中。这里我们实现一个钩子，在输出的日志中增加一个 app=awesome-web 字段。

设置钩子需要实现 logrus.Hook 接口:

```go
type Hook interface {
  Levels() []Level
  Fire(*Entry) error
}
```

Levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子。Fire 是日志输出前调用的钩子方法。

自定义一个钩子方法

```go
package main

import (
  "github.com/sirupsen/logrus"
)

type AppHook struct {
  AppName string
}

func (h *AppHook) Levels() []logrus.Level {
  return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
  entry.Data["app"] = h.AppName
  return nil
}

func main() {
  h := &AppHook{AppName: "awesome-web"}
  logrus.AddHook(h)

  logrus.Info("info msg")
}
```
