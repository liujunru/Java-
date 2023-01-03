package main
import(
	"fmt"
	"unsafe"
) 
func main(){
var n1,n2,n3 int
fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
//n1= 0 n2= 0 n3= 0

var n4,n5,n6 = 100,"tom",888
fmt.Println("n4=",n4,"n5=",n5,"n6=",n6)
//n4= 100 n5= tom n6= 888

n7,n8,n9 := 100,"tom~",88
fmt.Println("n7=",n7,"n8=",n8,"n9=",n9)
//n7= 100 n8= tom~ n9= 88

var n10 int64 = 10
fmt.Printf("n10的类型%T n10占用的字节数是%d",n10,unsafe.Sizeof(n10))
}