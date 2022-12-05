package main1
import(
	"fmt"
) 
func main(){
	var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var mychar byte = 'h'
    var str string

    //%d 格式化为十进制位 %q格式化为"value"
    str = fmt.Sprintf("%d",num1)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="99"
    str = fmt.Sprintf("%f",num2)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="23.456000"
    str = fmt.Sprintf("%t",b)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="true"
    str = fmt.Sprintf("%c",mychar)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="h"
}