## go语言编写Makefile

Makefile简单理解为它定义了一个项目文件的编译规则

```Makefile
# 声明一个虚拟的目标
# 如果不写这行，make会判断当前项目文件夹下是否有同名文件，如果有，那么就不执行了
.PHONY: all build run gotool clean help

# 这是一个变量
# 在makefile中使用变量${BINARY}
BINARY="bluebell"

# 输入make什么都不干的时候执行的命令
all: gotool build

build:
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
    @go run ./main.go cong/config.yaml

gotool:
    go fmt ./
    go vet ./

clean:
    @if [-f ${BINARY}] ;then rm ${BINARY} ;fi

help:
    @echo "make - 格式化go代码"
    @echo "make build -"
    @echo "make run -直接运行go代码"
    @echo "make clean - 移除二进制文件"

```