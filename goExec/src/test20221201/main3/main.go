package main
import(
	"fmt"
    "strconv"
) 
func main(){
	var str string = "true"
    var b bool

    //_ 表示忽略某值
    b,_ = strconv.ParseBool(str)
    fmt.Printf("b type %T b=%v\n", b, b)
    //b type bool b=true

    var str2 string = "123456"
    var n1 int64
    var n2 int
    n1, _ = strconv.ParseInt(str2,10,64)
    n2 = int(n1)
    fmt.Printf("n1 type %T n1=%v\n", n1, n1)
    fmt.Printf("n2 type %T n2=%v\n", n2, n2)
    //n1 type int64 n1=123456
    //n2 type int n2=123456

    var str3 string = "123.45"
    var f1 float64
    f1, _ = strconv.ParseFloat(str3,64)
    fmt.Printf("f1 trye %T f1=%v\n",f1, f1)
    //f1 trye float64 f1=123.45
}


