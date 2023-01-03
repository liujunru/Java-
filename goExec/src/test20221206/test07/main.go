package main
import(
	"fmt"
) 

func test(){
	defer func(){
	err := recover()
	if err != nil{
		fmt.Println("err=",err)
	}

}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main(){
	test()
	fmt.Println("下面的代码和逻辑。。。")
	//err= runtime error: integer divide by zero
	//下面的代码和逻辑。。。
}