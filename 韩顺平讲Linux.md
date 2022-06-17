![课程内容](https://gitee.com/liujunrull/image-blob/raw/master/202206151439749.png)

# 1. 软件安装

## 1.1 Vmware

  - 下载15.5版本
  - 激活。http://www.win7zhijia.cn/win10jc/win10_45016.html
  - 新建虚拟机，分配CPU和内核

## 1.2 centos 7.6 iso

  - 编辑虚拟机
  - 选择gnone安装，勾选传统x、兼容性程序库、开发工具。

  ![gnone安装](https://gitee.com/liujunrull/image-blob/raw/master/202206151700266.png)

  - 设置安装位置，选择我要配置分区，完成
  - 手动分区，点击+，选择boot分区，分配1G，文件系统选择ext4。点击+，选择swap，分配2G，文件系统swap。点击+，选择/，分配17G，文件系统选择ext4。设备类型标准分区。
  - 关闭KDUMP。内存泄露保护机制，测试环境为了节省内存将其关闭。
  - 设置网络和主机名。打开以太网
  - 关闭安全策略。

  ## 1.3 网络连接三种方式

  **桥接模式**

  在同一个号段内的，比如ip 192.168.0.10和192.168.0.20，好处是虚拟系统可以与外部系统通讯，但是ip地址有限，容易造成ip冲突

  **NAT模式**

  网络地址转换模式，虚拟机通过代理地址访问外部网络，外部网络不能直接访问虚拟机。既可以和外部系统通讯，又不造成ip冲突。

  **主机模式**

  独立系统，不和外部发生通讯。

#### 虚拟机克隆

1） 右键管理-克隆-完整克隆

2） 复制文件夹-打开-文件夹内虚拟机文件

## 1.4 安装vmtools-主机和虚拟机的共享文件夹

  1. 进入centos,弹出centos iso。
  2. 点击vm菜单的-install vmware tools
  3. centos会出现一个vm安装包，双击安装包，复制xx.tar.gz，点击其他位置-计算机-opt
  4. 拷贝到/opt
  5. 虚拟机桌面右键-打开终端。cd /opt/(进入opt目录）。tar -zxvf VM[Tab键补全文件名】，使用解压命令tar，得到一个安装文件
  6. 进入该vm解压目录，cd /vm...目录下
  7. 安装./vmware-install.pl
  8. 全部使用默认设置即可
  9. 安装vmtools需要有gcc

# 2. 目录结构

## 2.1 基本介绍

![主目录结构](https://gitee.com/liujunrull/image-blob/raw/master/202206171605987.png)


