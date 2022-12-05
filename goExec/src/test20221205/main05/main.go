package main
import(
	"fmt"
	"test20221205/utils"
) 
func main(){
	var x interface{}
	var y = 10.0
	x = y
	switch i :=x.(type){
	case nil:
		fmt.Printf("x的类型是:%T",i)
	case int:
		fmt.Printf("x的类型是:int")
	case float64:
		fmt.Printf("x的类型是:float64\n")
	case func(int) float64:
		fmt.Printf("x的类型是:ifunc(int)nt")
	case bool,string:
		fmt.Printf("x的类型是:bool或者string")
	default:
		fmt.Printf("未知类型")
	}
	result := utils.Cal(1,2);
	fmt.Println("result=",result)
	//x的类型是:float64
}