package main
import(
	"fmt"
) 

func getSum(n1 int,n2 int) int {
	return n1 + n2
}
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func main(){
	//给int取了别名，在go中myInt和int虽然都是int类型，但是go
	//认为myInt和int是两个类型
	type myInt int
	var num1 myInt

	var num2 int
	num1 = 40
	//num2 = num1//报错
	fmt.Println("num1=", num1,"num2=",num2)//res2= 110
}