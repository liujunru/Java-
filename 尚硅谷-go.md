## 1. golang变量
### 1.1 变量
#### golang变量使用的三种方式：

 1） 指定变量类型，声明后若不复制，使用默认值
````go
    func main(){
        var i int
        fmt.Println("i=",i)
    }
    //i=0
````
2）根据值自行判断变量类型（类型推导）

````go
var num = 10.11
fmt.Println("num=",num)
````
3）省略var，注意:=左侧的变量不应该是已经声明过的否则会编译错误。

````go
name :="tom"
fmt.Println("name=",name)
````

#### 多变量声明

````go
package main
import "fmt"

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
}
````

#### 定义全局变量（方法区外的变量）

````go
var n1 = 100
var n2 = 200
var name = "jack"

//上面的声明方式，也可以改成一次性声明
var(
    n3 = 100
    n4 = 200
    name2 = "mary"
)
````
全局变量可以定义后不使用，不可以使用:=定义

### 1.2 数据类型

![](https://gitee.com/liujunrull/image-blob/raw/master/202211301605947.png)

byte只能用于保存单个字母，一个汉字占3个字节，最少应使用int32来保存。

结构体（struct)相当于Java中的class

管道（channel）多用于处理并发

切片（slice）类似动态数组。

#### 1.2.1 整数类型
![](https://gitee.com/liujunrull/image-blob/raw/master/202211302122692.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302123110.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302128106.png)

- 如何查看某个变量占用字节大小和数量类型
````go
var n10 int64 = 10
fmt.Printf("n10的类型%T n10占用的字节数是%d",n10,unsafe.Sizeof(n10))
````

#### 1.2.2 浮点类型

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302200170.png)

**注意**

    - 浮点类型有固定的范围和字段长度，不受操作系统影响 
    - 浮点型默认声明为float64类型
    - 支持科学计数法形式

#### 1.2.3 字符类型(char)

golang中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte保存。

传统的字符串是由字符连接起来的字符序列，而go的字符串是由字节组成的。

````go
package main
import(
	"fmt"
) 
func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们直接输出byte值，就是输出对应的字符的码值
	fmt.Println("c1=",c1)//c1= 97
	fmt.Println("c2=",c2)//c2= 48

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1,c2)//c1=a c2=0

	//注意'北'为字符北，不是字符串"北"，此时格式化输出的是北对应的unicode码
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应的码值=%d",c3,c3)//c3=北 c3对应的码值=21271
}
````

**注意**：

     - go语言的字符使用UTF-8编码
     - 在go中，字符的本质是一个整数，直接输出时该字符对应的时UTF-8编码的码值，可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的Unicode字符。
     字符型存储到计算机中，需要将字符对应的码值（整数）找出来：
         存储：字符->对应码值->二进制->存储
         读取：二进制->码值->字符->读取

#### 1.2.4 string

**注意**

        - 字符串的两种表示形式：
        1）双引号，会识别转义字符
        2）反引号 ``,以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果。

        - 用+拼接多个字符串换行时，+需放在上行行末

#### 1.2.5 基本数据类型的默认值

![](https://gitee.com/liujunrull/image-blob/raw/master/202212011440164.png)

#### 1.2.6 基本数据类型的转换

go在不同类型的变量之间赋值时需要显式转换。也就是说golang中数据类型不能自动转换

表达式T(v)将值v转换为类型T

T就是数据类型，比如int32,int64,float32等等

v就是需要转换的变量

````go
package main
import(
	"fmt"
) 
func main(){
	var i int32 = 100
	//希望将i =>float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)

	fmt.Printf("i=%v  n1=%v  n2=%v\n",i,n1,n2)
	fmt.Printf("i的数据类型%T  n1的数据类型%T n2的数据类型%T",i,n1,n2)
	//i=100 n1=100 n2=100
	//i的数据类型int32  n1的数据类型float32 n2的数据类型int8
} 
````

被转换的是变量存储的数据（即）值，变量本身的数据类型并没有变化。

在转换时，比如将int64转为int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。

````go
package main
import(
	"fmt"
) 
func main(){
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2)//num2= 63

}
````

#### 1.2.7 基本数据类型和string的互相转换

##### 基本类型转string

- fmt.Sprintf("%"参数,表达式)：参数需要和表达式的数据类型相匹配

````go
package main
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
````
- 使用strconv包的函数
- 使用Itoa函数
````go
//Itoa是FormatInt(i, 10) 的简写。
func Itoa(i int) string
````
  
##### string转基本数据类型
- 使用strconv包的函数

````go
//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
func ParseBool(str string) (value bool, err error)
//返回字符串表示的整数值，接受正负号。

//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；

//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
func ParseInt(s string, base int, bitSize int) (i int64, err error)

//解析一个表示浮点数的字符串并返回其值。

//如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
func ParseFloat(s string, bitSize int) (f float64, err error)

//ParseUint类似ParseInt但不接受正负号，用于无符号整型。
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
````

**注意**
在将string类型转为基本数据类型时，要确保string类型能够转成有效的数据，如果不能转换，
golang会将其转为默认值。

#### 1.2.8 指针

- 获取变量的地址，用&
- 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值。比如：var ptr * int = &num
- 获取指针类型所指向的值，用*。比如：var ptr *int

````go
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
````

**内存图**
![](https://gitee.com/liujunrull/image-blob/raw/master/202212011911760.png)


