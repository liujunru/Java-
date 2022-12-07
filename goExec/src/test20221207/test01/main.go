package main
import(
	"fmt"
) 
func main(){
	var slice3 []int = []int{100,200,300}
	//通过append直接给slice3追加具体的元素
	slice7 := append(slice3, 400, 500, 600)
	fmt.Println("slice7", slice7)
	//slice7 [100 200 300 400 500 600]

	//通过append将切片slice3追加给slice3
	slice6 := append(slice3,slice7...)
	fmt.Println("slice6", slice6)
	//slice6 [100 200 300 100 200 300 400 500 600]
	
	//使用copy内置函数完成拷贝
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int,10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)//slice4= [1 2 3 4 5]
	fmt.Println("slice5=", slice5)//slice5= [1 2 3 4 5 0 0 0 0 0]
}