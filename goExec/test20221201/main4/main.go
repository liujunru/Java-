package main
import(
	"fmt"
) 
func main(){
	//基本数据类型在内存布局
    var i int = 10
    //&i表示i的地址
    fmt.Println("i的地址=",&i)//i的地址= 0xc00000e0b8

    //下面的var ptr *int = &i
    //ptr时一个指针变量
    //ptr的类型是*int
    //ptr本身的值&i
    var ptr *int = &i
    fmt.Printf("ptr=%v\n",ptr)//ptr=0xc00000e0b8
    fmt.Printf("ptr 的地址=%v\n",  &ptr)//ptr 的地址=0xc00000a030
    fmt.Printf("ptr 的指向的值 =%v",*ptr)//ptr 的指向的值 =10
}


