package main
import(
	"fmt"
) 
func main(){
	var intArr [3]int64
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p\tintArr[0]地址%p\tintArr[1]的地址=%p\tintArr[2]的地址=%p",&intArr,
	&intArr[0],&intArr[1],&intArr[2])	
}