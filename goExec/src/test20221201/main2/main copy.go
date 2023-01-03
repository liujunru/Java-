package main
import(
	"fmt"
    "strconv"
) 
func main(){
	var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var str string

    //返回i的base进制的字符串表示。base 必须在2到36之间，
    //结果中会使用小写字母'a'到'z'表示大于10的数字。
    str = strconv.FormatInt(int64(num1),10)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="99"
    str = strconv.FormatFloat(num2,'f',10,64)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="23.4560000000"
    str = strconv.FormatBool(b)
    fmt.Printf("str type %T str=%q\n",str,str)
    //str type string str="true"
   
}