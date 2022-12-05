# 1. golang变量
## 1.1 变量
### golang变量使用的三种方式：

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

### 多变量声明

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

### 定义全局变量（方法区外的变量）

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

## 1.2 数据类型

![](https://gitee.com/liujunrull/image-blob/raw/master/202211301605947.png)

byte只能用于保存单个字母，一个汉字占3个字节，最少应使用int32来保存。

结构体（struct)相当于Java中的class

管道（channel）多用于处理并发

切片（slice）类似动态数组。

###  整数类型
![](https://gitee.com/liujunrull/image-blob/raw/master/202211302122692.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302123110.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302128106.png)

- 如何查看某个变量占用字节大小和数量类型
````go
var n10 int64 = 10
fmt.Printf("n10的类型%T n10占用的字节数是%d",n10,unsafe.Sizeof(n10))
````

###  浮点类型

![](https://gitee.com/liujunrull/image-blob/raw/master/202211302200170.png)

**注意**

    - 浮点类型有固定的范围和字段长度，不受操作系统影响 
    - 浮点型默认声明为float64类型
    - 支持科学计数法形式

###  字符类型(char)

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

###  string

**注意**

        - 字符串的两种表示形式：
        1）双引号，会识别转义字符
        2）反引号 ``,以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果。

        - 用+拼接多个字符串换行时，+需放在上行行末

###  基本数据类型的默认值

![](https://gitee.com/liujunrull/image-blob/raw/master/202212011440164.png)

###  基本数据类型的转换

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

###  基本数据类型和string的互相转换

#### 基本类型转string

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

###  指针

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


#### 值类型和引用类型

**值类型**

变量直接存储值，内存通常在栈中分配

**引用类型**

变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在堆上分配。当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC回收。

#### 标识符
//TODO
标识符小写开头只能在本包中使用，大写开头可以被其他包引用（新版有其他包管理方式）

**go中保留关键字**

![](https://gitee.com/liujunrull/image-blob/raw/master/202212031400187.png)


**预定义标识符**

![](https://gitee.com/liujunrull/image-blob/raw/master/202212041221298.png)

## 1.3 运算符

### 算术运算符

##### 取模%

取模公式：a%b = a-a/b*b,符号位与a一致

````go
package main
import(
	"fmt"
) 
func main(){
	fmt.Println("10%3=",10%3)//10%3= 1
	fmt.Println("-10%3=",-10%3)//-10%3= -1
	fmt.Println("10%-3=",10%-3)//10%-3= 1
	fmt.Println("-10%-3=",-10%-3)//-10%-3= -1
}
````

go中的i++ i--只能独立使用,没有++i和--i

````go
package main
import (
    "fmt"
)
func main(){
    var i int = 8
    var a int
    a = i++//错误
    a = i--//错误
    if i++ >0{
        fmt.Println("ok")
    }
}
````

### 关系运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051129523.png)

### 逻辑运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051133741.png)

### 赋值运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051141713.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051149690.png)

#### 运算符优先级

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051200474.png)

### 位运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051202682.png)

### 其他运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202212051203042.png)

**go明确不支持三元运算符**

#### 二进制转八进制

规则：将二进制数每三位一组（从低位开始组合），转成对应的八进制数即可。

案例：将二进制11010101转成八进制

11 010 101 = 0 3 2 5

#### 二进制转成十六进制

规则：将二进制数每三位一组（从低位开始组合），转成对应的八进制数即可。

案例：将二进制11010101转成八进制

1101 0101 = 0x D 5

#### 八进制转二进制

规则：将八进制每1位转成对应的一个3位的二进制数即可

0237 =  10 011 111

#### 十六进制转二进制
 
规则：将十六进制每1位转成对应的一个4位的二进制数即可

0X237 = 10 0011 0111

#### 原码、反码、补码

对于有符号的而言：

- 二进制的最高位是符号位，0表示正数，1表示负数
````go
1=>[0000 0001] -1=>[1000 0001]
````
- 整数的原码、反码、补码都一样

- 负数的反码=它的原码符号位不变，其他位取反（0-1,1->0)

````go
 1=>原码[0000 0001] 反码[0000 0001] 补码[0000 0001]

-1=>原码[1000 0001] 反码[1111 1110] 补码[1111 1111]
````

- 负数的补码=它的反码+1
- 0的反码补码都是0
- 在计算机运算的时候，都是以补码的方式来运算的
````go
1+1 1-1 = 1 + (-1)
````

### 位运算符

- 按位与&：两位全为1，结果为1，否则为0
- 按位或|：两位有一个为1，结果为1，否则为0
- 按位异或^:两位一个为0，一个为1，结果为1，否则为0
  **移位运算符**
  - 右移运算符 >>:低位溢出，符号位不变，并用符号位补溢出的高位（除2的位移数）
  - 左移运算符 <<:符号位不变，低位补0（乘2的位移数）
````go
2&3
2的补码 0000 0010
3的补码 0000 0011
2&3     0000 0010 =》2

2|3
2的补码 0000 0010
3的补码 0000 0011
2|3    0000 0011 =》3

2^3
2的补码 0000 0010
3的补码 0000 0011
2^3     0000 0001 =》1

-2^2

-2的原码 1000 0010
-2的反码 1111 1101
-2的补码 1111 1110
2的补码  0000 0010
-2^2     1111 1100（补码）
         1111 1011（反码）
         1000 0100（原码）=》-4

````

# 2 流程控制

## 2.1 分支控制 if-else

### 单分支

基本语法：

    if 条件表示式{
        执行代码块
    }

**注意**：{}必须要有，即便只有一行代码

golang支持在if中，直接定义一个变量，比如下面

````go
if age := 20; age > 18 {
    fmt.Println（"你已年满18！")
}
````
### 双分支

基本语法：

    if 条件表达式{
        执行代码块
    }else{
        执行代码块
    }

### 多分支

    if 条件表达式{
            执行代码块
        }else if{
            执行代码块
        }else if{
            执行代码块
        }
        ...
        }else{
            执行代码块
        }

## switch分支

case后的表达式可以有多个，可以用","隔开。
匹配项后面不需要加break

基本语法：

    switch 表达式{
        case 表达式1，表达式2，...:
            语句块1
        case 表达式3，表达式4,...:
            语句块2
        //这里可以有多个case语句

        default:
            语句块
    }

**细节**

- switch表达式可以是常量，变量，甚至是函数
- case后面的表达式的数据类型要和switch表达式数据类型一致
- case后面的表达式如果是常量（字面量），则不能重复
- default语句可以没有
- switch后面可以不带表达式，可以看作switch true{},类似if-else分支来使用

````go
var age int = 10
switch{
    case age == 10:
        fmt.Println("age == 10")
    case age == 20:
        fmt.Println("age == 20")
    default:
        fmt.Println("没有匹配到")
}
````

- switch后可以定义/声明一个变量，分号结束，不推荐使用
- switch穿透，fallthrough。如果在case语句块后增加fallthrough，则后继续执行下一个case，也叫switch穿透

````go
var age int = 10
switch{
    case age == 10:
        fmt.Println("age == 10")
        fallthrough//默认只能穿透一层
    case age == 20:
        fmt.Println("age == 20")
    default:
        fmt.Println("没有匹配到")
}
//age == 10
//age == 20
````

- Type switch:switch语句话可以被用于type-switch来判断某个interface变量中实际指向的变量类型。

````go
package main
import(
	"fmt"
) 
func main(){
	var x interface{}
	var y = 10.0
	x = y
	switch i :=x.(type){
	case nil:
		fmt.Printf("x的类型是:%T",i)
	case int:
		fmt.Printf("x的类型是:int")
	case float64:
		fmt.Printf("x的类型是:float64")
	case func(int) float64:
		fmt.Printf("x的类型是:ifunc(int)nt")
	case bool,string:
		fmt.Printf("x的类型是:bool或者string")
	default:
		fmt.Printf("未知类型")
	}
	//x的类型是:float64
}
````

### for循环控制

基本语法：

    for 循环变量初始化；循环条件；循环变量迭代{
        循环操作（语句）
    }


**注意**：

    for{
        循环执行语句
    }


上面的写法等价for;;{}是一个无限循环，通常需要配合break语句使用

#### for循环遍历字符串

- 传统for循环方式,按字节来进行遍历，有中文时不转换会变成乱码

````go
str2 := "hekjndfr参数"
str3 := []rune(str2)
	for i:= 0; i < len(str3); i++{
		fmt.Printf("%c \n",str3[i])
	}
````
- for-range方式，按照字符来遍历，中文不会出现乱码
  
````go
package main
import(
	"fmt"
) 
func main(){
	str := "hello,worldv地方的色"
	for index, val := range str{
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
    // index=0,val=h
    // index=1,val=e
    // index=2,val=l
    // index=3,val=l
    // index=4,val=o
    // index=5,val=,
    // index=6,val=w
    // index=7,val=o
    // index=8,val=r
    // index=9,val=l
    // index=10,val=d
    // index=11,val=v
    // index=12,val=地
    // index=15,val=方
    // index=18,val=的
    // index=21,val=色
    // index=24,val=!
}
````

### while、do while循环

go中本身没有while、do while循环，可以通过for循环实现效果

while基本语法：

    for{
        if 循环条件表达式{
            break//跳出for循环
        }
        循环操作（语句）
        循环变量迭代
    }

````go
func main(){
    var i int = 1
    for{
        if i > 10{
            break
        }
        fmt.Println("hello,world",i)
        i++
    }
    fmt.Println("i=",i)
}
````

do while基本语法

    循环变量初始化
    for{
        循环操作（语句）
        循环变量迭代
        if 循环条件表达式{
            break//跳出for循环
        }
    }

````go
    var j int = 1
    for{
        fmt.Println("hello,ok",j)
        j++
        if j > 10{
            break
        }
    }
````

### break使用细节

break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的时哪一层语句块。不指定标签时默认跳出最近的for循环

````go
lable2:
for i = 0; i < 4; i++{
    lable1:
    for j := 0; j < 10; j++{
        break lable2
    }
    fmt.Println("j=",j)
}
````

### goto

- go语言的goto语句可以无条件的转移到程序中指定的行
- goto语句通常与条件语句配合使用，可以用来实现条件转移，跳出循环体等功能
- 在go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难

````go
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
// hello1
// hello5
// hello6
````

# 3. 函数（方法）

基本语法：

    func 函数名（形参列表）（返回值列表）{
        执行语句...
        return 返回值列表
    }

函数可以有返回值，也可以没有

## 包

1.19版本下go mod导包方式：
[1.19](https://juejin.cn/post/7075950668489424933)

**注意**

- 文件的包名通常和文件所在的文件夹名一致，一般为小写字母
- 在import包时，路径从$GOPATCH的src下开始，不用带src，编译器会自动从src下开始导入