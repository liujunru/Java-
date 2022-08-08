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

![主目录结构](https://gitee.com/liujunrull/image-blob/raw/master/202206171605987.png

# 3. 指令

## 3.1 vim指令
     vim xxx.xml:进入某文件,i进入编辑模式，esc退出编辑模式，:wq保存并退出

  ![vim模式切换](https://gitee.com/liujunrull/image-blob/raw/master/202206220911296.png)

     yy复制行，3yy复制3行，p复制
     dd删除行,3dd删除3行
     /关键字，查找关键词，输入n查找该关键词的下一个位置
     :set nu:显示行号，:set nonu:取消行号显示
     G：跳到末尾一行，gg：跳到首行
     u：撤销
     20 shift+g:跳到第20行

## 3.2 登录注销、切换用户

     shutdown 关机
     reboot 重启
     关机重启之前先sync，将内存数据同步到磁盘

     用普通用户身份登录时，切换root用户指令 su - root,登出root用户指令logout,最后一个用户登出logout，此时会退出系统

## 3.3 用户管理

在root用户下操作

     useradd xxx：添加用户xxx，未指定组时默认放在同用户名组下
     useradd -d 指定目录 新的用户名：给新创建的用户指定家目录
     passwd 用户名:修改指定用户密码
     userdel 用户名：删除用户，保留家目录，建议保留家用户操作
     userdel -r 用户名:删除用户及家目录
     id 用户名：查看用户
     groupadd xxx:增加组，组类似于角色的概念
     groupdel xxx:删除组
     useradd -g 用户组 用户名:增加用户时将他指定组
     usermod -g 用户组 用户名:修改用户到某用户组
     init 3:切换centos运行状态3，多用户有网络状态，5为图形化状态
     systemctl get-default:查看当前默认运行状态
     systemctl set-default graphical.target:设置默认运行状态为3，multi-user.target为5

  **修改root密码**

  1. 开机界面输入e进入编辑界面
  2. 找到linux16所在行，在行末尾输入init=/bin/sh,然后Ctrl+x进入单用户模式。
  3. 在光标闪烁的位置输入mount -o remount,rw /,然后回车
  4. 在新一行最后输入passwd,然后回车输入密码，再次回车确认密码
  5. 密码修改成功后，会显示passwd...的样式，证明修改成功
  6. 在鼠标闪烁位置输入touch /.autorelabel，回车
  7. 输入exec /sbin/init,回车等待系统自动重启

  ## 3.4 文件目录指令

     ls l:以列的形式查看目录
     ls a：查看目录包括隐藏文件
     cd ~：回到自己的家目录
     cd ..:回到当前目录的上一级目录
     pwd：显示当前所在目录
     mkdir -p /home/dog:创建多级目录
     rmdir:删除空目录
     rm -rf:删除目录
     touch 文件名：更新文件状态，比如操作时间
## 文件目录类

### cp指令

拷贝指定文件到指定目录。

#### 常用语法
cp [选项] source dest

-r ：递归复制整个文件夹

#### 案例：

将/home/hello.txt拷贝到/home/bbb目录下：

cp hello.txt /home/bbb

将home/bbb整个目录，拷贝到/opt

cp -r/home/bbb /opt

如果重复拷贝提示是否强制覆盖，可以用\cp直接强制覆盖不提示。

### rm指令

移除文件或目录

#### 基本语法

rm [选项] 要删除的文件或目录

#### 常用选项

-r:递归删除整个文件夹

-f:强制删除不提示

rm -rf /home/bbb [删除整个文件夹，不提示]

### mv指令

移动文件与目录或者重命名

基本语法：

 - mv oldNameFIle newNameFile (统一目录下移动为从重命名)
 - mv oldNameFIle /temp/movefile/targetFolder(不同目录下为移动文件)
 - mv  oldNameFIle /temp/movefile/targetFolder newNameFile(不同目录下为移动文件并重命名)
 - mv opt/bbb /home/uuu(移动bbb文件夹到home文件夹并改名uuu文件夹名)

### cat指令

 - cat -n 文件（查看文件并显示行号）

### echo指令

输出内容到控制台

    - 使用echo指令输出环境变量，比如输出$PATH $HOSTNAME,echo $HOSTNAME
    - 使用echo指令输出hello,world

### head指令

显示文件的开头部分内容，默认情况下head指令显示文件的前10行内容

    - head 文件（查看文件头10行内容）
    - head -n  5 文件（查看文件头5行内容）

### tail指令

用于输出文件中尾部的内容，与head指令对应

    - tail 文件（查看文件尾10行内容）
    - tail -n  5 文件（查看文件尾5行内容）
    - tail -f 文件（实时追踪该文档所有更新）

### >指令和>>指令

>指令 覆盖写
>>指令 追加

    - ls -l > 文件（将列表的内容覆盖写入文件a.txt中）
    - ls -al >> 文件（将列表的内容覆盖写入文件a.txt中末尾）
    - cat 文件1 > 文件2 （将文件覆盖文件2）
    - echo “内容” >> 文件（将文件追加到文件尾）
  
### history指令

查看已经执行过的历史指令，也可以执行历史指令

## 2.3 时间日期类

### date指令-显示当前时间

1） date:显示当前时间

2）date %Y:显示当前年份

3）date %m:显示当前月份

4）date &d:显示当前是哪一天

5）date "+%Y-%m-%d %H:%M:%S":显示年月日时分秒

date -s "2020-11-03 20：02：10" :设置系统当前时间为2020-11-03 20：02：10

### cal 指令

查看日历指令

#### 基本语法

    - 显示当前日历 cal
    - 显示2020年日历： cal 2020

## 2.4 搜索查找类

### find指令

将从指令目录向下递归的遍历其各个子目录，将满足条件的文件或目录显示在终端

**指令**

-  -name <查询方式>：按照指令的文件名查找模式查找文件
-   -user <用户名>：查找属于指定用户名所有文件
-  -size <文件大小>：按照指定的文件大小查找文件
### locate 指令

locate指令可以快速定位文件路径。locate指令利用事先建立的系统中所有文件名及路径的locate数据库实现快速定位给定的文件、locate指令无需遍历整个文件系统，查询速度较快。为了保证查询结果的准确度，管理员必须定期更细locate时刻。
   
**基本语法**

locate <搜索文件>

**特别说明**

由于locate语法基于数据库进行查询，所以第一次运行前，必须使用updatedb指令创建locate数据库。

**应用实例**

请用locate指令快速定位hello.txt文件所在目录
