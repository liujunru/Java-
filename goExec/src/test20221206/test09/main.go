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
	//[10 20 30]
	// intArr的地址=0xc000012138  
	// intArr[0]地址0xc000012138  
	// intArr[1]的地址=0xc000012140
	// intArr[2]的地址=0xc000012148
}