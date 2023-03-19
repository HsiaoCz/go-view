## K8s

由于现在现在部署服务基于docker，当docker多了之后，难以避免的出现容器的管理问题

k8s就是用来管理docker容器的

**k8s架构**

假如部署了三台k8s机器，有一台主节点
主节点上有三个东西:
APIServer:所有节点和manager通信的入口
Scheduler:调度节点
controller manager:控制节点

kubelet:节点客户端

kubctl：命令行，和APIServer进行沟通

网络：容器里面怎么通信
k8s里面有个CNI

服务发现：coredns\dashboard

ETCD:k8s里面的数据和状态存储的位置

ingress:给外部提供一个服务的入口


k8s搭建需要搭建的东西:APIServer、scheduler、controller manager、kubelet、kubctl、CNI、服务发现、ETCD、ingress

外部只需要部署Kubelet、和Kubctl

k8s部署的每一个节点都有一个节点的ip，nodeip,每个节点内部会部署很多的容器，容器内部又有自己的ip,容器内的ip由CNI来分发
用户访问的时候，通过nodeip来访问
### k8s搭建

**1、环境准备**

1、首先需要三台服务器，需要centos或者ubuntu
2、修改hostname 在/etc/hostname中修改,比如172.17.187.88 172.17.187.89  172.17.187.90，这三个ip作为nodeip，也就是节点ip
3、配置host主机名与ip对应关系，在/etc/hosts 172.17.187.88 master
172.17.187.89 node1 172.17.187.90 node2
4、配置免密登录
5、禁用防火墙:syetemctl stop firewalled systemctl disable firewalled
6、禁用selinux
7、禁用缓冲区：swapoff -a,/etc/fstab
8、内核：将流量传递到iptables链

**2、安装docker**
安装完docker之后，需要配置一下，镜像和cgroup
在etc/docker目录下的deamon.json下进行配置

**3、安装k8s**

安装kubeadm,kubelet,和kubctl

### k8s client介绍

### k8s管理

