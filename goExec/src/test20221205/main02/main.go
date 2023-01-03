package main
import(
	"fmt"
) 
func main(){
	fmt.Println("hello1")
	goto label1
	fmt.Println("hello2")
	fmt.Println("hello3")
	fmt.Println("hello4")
	label1:
	fmt.Println("hello5")
	fmt.Println("hello6")
	
}