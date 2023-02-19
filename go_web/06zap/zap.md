## zap

安装
```go
go get go.uber.org/zap
```

快速使用
```go
package main

import (
  "time"

  "go.uber.org/zap"
)

func main() {
  logger := zap.NewExample()
  defer logger.Sync()

  url := "http://example.org/api"
  logger.Info("failed to fetch URL",
    zap.String("url", url),
    zap.Int("attempt", 3),
    zap.Duration("backoff", time.Second),
  )

  sugar := logger.Sugar()
  sugar.Infow("failed to fetch URL",
    "url", url,
    "attempt", 3,
    "backoff", time.Second,
  )
  sugar.Infof("Failed to fetch URL: %s", url)
}
```

zap提供了两种类型的日志记录器(sugared logger)和(logger)

在性能很好但不是很关键的上下文中，使用SugaredLogger。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。

在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。它甚至比SugaredLogger更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。

结合到项目中
[https://www.liwenzhou.com/posts/Go/zap/]