package main
import(
	"fmt"
) 
func main(){
	var n1 int32 = 12
    var n3 int8
    //var n4 int8
    //n4 = int8(n1) +127
    n3 = int8(n1) + 128
    fmt.Println(n3)

}