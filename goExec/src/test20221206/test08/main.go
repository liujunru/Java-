package main
import(
	"fmt"
	"errors"
) 
func readConf(name string) (err error){
	if name == "config.ini" {
		//读取
		return nil
	}else{
		//返回一个自定义错误
		return errors.New("读取文件错误。。")
	}
}

func test(){
	err := readConf("config2.ini")
	if err != nil{
		//panic:捕获自定义异常并退出程序
		panic(err)
	}
	fmt.Println("程序正常运行")
}

func main(){
	test()
	fmt.Println("main程序正常运行")
	// panic: 读取文件错误。。

	// goroutine 1 [running]:
	// main.test()
	//         E:/know/笔记/goExec/src/test20221206/test08/main.go:20 +0x49
	// main.main()
	//         E:/know/笔记/goExec/src/test20221206/test08/main.go:26 +0x19
	// exit status 2
}