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
    - 字符型存储到计算机中，需要将字符对应的码值（整数）找出来：
  
         存储：字符->对应码值->二进制->存储
         读取：二进制->码值->字符->读取




