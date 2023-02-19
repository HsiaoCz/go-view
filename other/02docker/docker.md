## docker

**1、docker是什么？**

docker是一种容器技术，是一种沙盒技术。它提供了一种非常遍历的打包机制，这种机制直接打包了应用运行所需要的整个操作系统，从而能够保证本地环境（开发环境）和生产环境（运行环境）的高度一致

镜像与容器：
镜像是一个静态的概念，容器是一个动态的概念，容器是镜像的实例，通俗的讲，镜像就是放在硬盘上的，容器时基于镜像跑起来的东西。

配置docker可以修改/etc/docker/daemon.json文件

比如可以设置镜像存储位置:

docker镜像默认的存储位置在根目录，可以改一下
```json
{
    "data_root":"myownpath", ### 存储位置
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}

// 设置镜像源
{
    "data_root":"myownpath", ### 存储位置
     "registry-mirrors": ["https://6kx4zyno.mirror.aliyuncs.com"], ### 镜像源，可以设置多个
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "/usr/bin/nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}
```

*设置docker的代理*

docker是一个C/S架构，我们执行的docker命令实际是一种客户端，它会发起REST API到daemon(Server端），由daemon去拉取需要的镜像。此节设置的就是daemon的代理。几乎所有的daemon相关设置都可以在daemon.json中完成，但代理是个例外，这个设置需要创建：
/etc/systemd/system/docker.service.d/http-proxy.conf 文件。

```
[Service]
Environment="HTTP_PROXY=http://proxy.example.com:80"
Environment="HTTPS_PROXY=https://proxy.example.com:443"
Environment="NO_PROXY=localhost,127.0.0.1,docker-registry.example.com,.corp"  ### 设置一些ip跳过代理
```

*容器代理*

这个需要起一个配置文件
创建~/.docker/config.json

```json
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://192.168.1.12:3128",
     "httpsProxy": "http://192.168.1.12:3128",
     "noProxy": "*.test.example.com,.example2.com,127.0.0.0/8"
   }
 }
}
```

**2、docker的使用**

*启动容器*

```bash
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
```
docker run有很多参数可以设置

```
-d: 后台运行容器，并返回容器ID；
-i: 以交互模式运行容器，通常与 -t 同时使用；
-P: 随机端口映射，容器内部端口随机映射到主机的端口
-p: 指定端口映射，格式为：主机(宿主)端口:容器端口
-t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；
--name="nginx-lb": 为容器指定一个名称；
--cpuset="0-2" or --cpuset="0,1,2": 绑定容器到指定CPU运行；
-m :设置容器使用内存最大值；
--net=bridge: 指定容器的网络连接类型，支持 bridge/host/none/container: 四种类型；
--expose=[]: 开放一个端口或一组端口；
--volume , -v: 绑定一个卷
--rm ,退出容器后删除名字 
--restart ,重启选项，有no/always/on-failure/unless-stopped
--entrypoint ,重写容器进程的入口
```

比如我们执行:

```bash
sudo docker run -it --name test ubuntu:16.04
```
这种是以前台交互式的允许，会启动并且进入容器，我们像使用普通终端一样去安装工具
如果以后台形式运行：

```bash
sudo docker run -d --name test ubuntu:16.04
```
进入容器：

```bash
docker exec -it 容器ID/名称 bash
```

使用`docker update`可以修改容器运行时指定的参数

**3、docker的进阶使用**

1、持久化(挂载主机硬盘)
启动时通过-v 主机目录：容器目录选项即可讲主机的目录挂载到容器中

```go
sudo docker run -d --name test -v /home/xxx:/root/xxx ubuntu:16.04
```

2、端口映射
有时候容器内启动一个网络服务，这个服务去监听一个接口，但它监听的实际是容器的内部端口，直接访问是不行的
通过-p 主机端口：容器端口 或直接使用主机网络--net=host

```bash
docker run -d -p 5000:5000 ubuntu:16.04 
docker run -d --net=host ubuntu:16.04 
```

3、自定义启动命令

截止到目前，我们都没有指定过容器启动后运行什么命令，其实run的最后一个参数可以用于在启动容器后运行的命令

```bash
docker run -d --name test ubuntu:16.04  /bin/bash
docker run -d --name test ubuntu:16.04  sh -c “/run.sh && /bin/bash” ### 多条命令拼接
```

4、容器状态/日志查看

```bash
docker logs [-f等选项] 容器名/ID
```

5、对容器修改的提交

很多时候我们基于一个镜像启动了容器，在容器中我们安装了我们需要的软件，想在容器删除后也能够使用，而不是再装一次。这时就需要我们能够提交这个修改。和git类似，也是通过commit指令去提交。
```bash
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

Create a new image from a container's changes

Options:
  -a, --author string    Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")
  -c, --change list      Apply Dockerfile instruction to the created image
  -m, --message string   Commit message
  -p, --pause            Pause container during commit (default true)
  
```

例如我们修改了ubuntu,我们装了一些软件

```bash
docker commit test ubuntu:my16.04
```

现在生成了一个新的镜像，我们可以使用它

```bash
docker run -d --name test ubuntu:my16.04  /bin/bash
```

**4、制作镜像**

1、docker file编写

通过在容器内修改再提交的方式虽然能够生成镜像，但手动操作太多，而且不便于自动化。更常用的制作镜像的方式是Dockerfile。Dockerfile的基本使用比较简单，只需要掌握几个关键字：

```Dockerfile
FROM ubuntu:16.04 ### FROM: 基础镜像
ENV LANG C.UTF-8 
ENV TZ=Asia/Shanghai ### 设置容器的时区, ENV用于设置环境变量

RUN mkdir /opt/alg ### RUN: 执行一条命令，多个命令可以通过&&

ADD config/ /opt/alg/config/ ### ADD: 除有COPY的功能外，还能通过URL下载文件，并且会自动解压缩
COPY Dependency/ /opt/alg/Dependency/ ### COPY: 拷贝宿主机的文件或文件夹到镜像
COPY bin/ /opt/alg/bin/
COPY models/ /opt/alg/models/

ENTRYPOINT ["/opt/alg/config/start_service.sh" ]  ### 设置容器启动的入口，类似于main函数，在docker run中可以通过 --entrypoint=XXX 覆盖，如果有这个，那么docker run时设置的command就会被当作它的参数

```

可以用entrypoint去指定入口，也可以使用CMD指定，这二者之间的差异
[https://blog.csdn.net/wuce_bai/article/details/88997725]

2、生成镜像

在dockerfile的目录执行

```bash
docker build . -t 镜像名:标签
例如：
docker build . -t myapp:v1
```

生成镜像之后可以像之前那样去使用

3、镜像保存、载入

镜像既可以上传至官方的DockerHub供人pull,也可以自行搭建私有化的镜像仓库（如harbor)。但对于普通人或日常使用，更多的可能是想将镜像保存成一个可传输的文件，然后放到其他机器，再载入。这个docker也是有对应命令支持的

```bash
docker save [OPTIONS] IMAGE [IMAGE...]
> docker save -o my_ubuntu_v3.tar runoob/ubuntu:v3 ###将镜像runoob/ubuntu:v3 保存成my_ubuntu_v3.tar
docker load
--input , -i : 指定导入的文件，代替 STDIN。
--quiet , -q : 精简输出信息。
> docker load -i my_ubuntu_v3.tar
```

也可以结合其他压缩软件的命令，直接保存出压缩包

```bash
docker save <myimage>:<tag> | gzip > <myimage>_<tag>.tar.gz
gunzip -c <myimage>_<tag>.tar.gz | docker load
```

4、显卡使用

对于深度学习部署，很多可能需要显卡，使用docker时，需要保证显卡驱动安装，同时按上述步骤安装了nvidia-docker2。
启动容器时，增加--gpus选项即可：

```bash
sudo docker run --rm --gpus all nvidia/cuda:11.0-base nvidia-smi ### all: 所有显卡都可用
sudo docker run --rm --gpus device=0,2 nvidia/cuda:11.0-base nvidia-smi ### 0,2 卡可用
也可以用下列方式：
sudo docker run --rm --gpus '"device=0"' nvidia/cuda:11.0-base nvidia-smi ### 0卡可用
```

5、其他使用命令

```bash
docker images 列出所有镜像
docker rmi 删除镜像
docker cp 宿主机和容器间拷贝文件
```