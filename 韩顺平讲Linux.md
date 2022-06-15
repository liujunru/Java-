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
