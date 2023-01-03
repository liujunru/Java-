1. debug失效：delve版本太老

   [Delve版本太老导致Intellij IDEA无法debug GO程序的解决方案 - 首席CTO笔记 (shouxicto.com)](https://www.shouxicto.com/article/321.html)

2. go版本和开发工具都选用最新版，防止版本对不起来

3. 包管理方式，go mod和gopath选用一种即可，建议gomod模式

   ````go
   //生成go.mod文件
   go mod init 项目名
   //初始化包下载
   go mod tidy
   ````

4. 其他包引用本包变量，本包报错“变量没有使用”，把变量删掉重写

