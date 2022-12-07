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
2&3     0000 0010 => 2

2|3
2的补码 0000 0010
3的补码 0000 0011
2|3    0000 0011 => 3

2^3
2的补码 0000 0010
3的补码 0000 0011
2^3     0000 0001 => 1

-2^2

-2的原码 1000 0010
-2的反码 1111 1101
-2的补码 1111 1110
2的补码  0000 0010
-2^2     1111 1100(补码)
         1111 1011(反码)
         1000 0100(原码) => -4

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
    fmt.Println("你已年满18！")
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

## 2.2 switch分支

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

## 2.3  for循环控制

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

## 2.4 while、do while循环

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

## 3.1 包

1.19版本下go mod导包方式：
[1.19](https://juejin.cn/post/7075950668489424933)

**注意**

- 文件的包名通常和文件所在的文件夹名一致，一般为小写字母
- 在import包时，路径从$GOPATCH的src下开始，不用带src，编译器会自动从src下开始导入
- 在同一个包下，不能有相同的函数名（也不能有相同的全局变量名）。否则报重复定义 ->没有 重载

## 3.2 函数

**函数细节**

- 基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值
- 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，效果上看类似引用。
- 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以堆函数调用

````go
package main
import(
	"fmt"
) 

func getSum(n1 int,n2 int) int {
	return n1 + n2
}
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func main(){
	a := getSum
	fmt.Printf("a的类型%T，getSum类型是%T\n", a,getSum)
	
	res := a(10,40)
	fmt.Println("res=", res)

	//a的类型func(int, int) int，getSum类型是func(int, int) int
	//res= 50

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)//res2= 110
}
````

- go支持自定义数据类型

        基本语法：type 自定义数据类型名 数据类型
        案例：type myInt int //这时myInt等价int来使用
              type mySum func(int,int) int//这时mySum等价一个函数类型func(int,int)

    ````go
    package main
    import(
        "fmt"
    ) 

    func getSum(n1 int,n2 int) int {
        return n1 + n2
    }
    func myFun(funvar func(int, int) int, num1 int, num2 int) int {
        return funvar(num1, num2)
    }
    func main(){
        //给int取了别名，在go中myInt和int虽然都是int类型，但是go
        //认为myInt和int是两个类型
        type myInt int
        var num1 myInt

        var num2 int
        num1 = 40
        //num2 = num1//报错
        fmt.Println("num1=", num1)//num1= 40 
    }
    ````

- 支持对函数返回值命名

    ````go
    func getSumAndSub(n1 int,n2 int) (sum int, sub int){
        sum = n1 + n2
        sub = n1 - n2
        return 
    }
    ````

- go支持可变参数，放在形参列表的最后

    ````
    func sum(args... int) sum int{}
    ````

args是slice，通过args[index]可以访问到各个值

### init函数

每个源文件都可以包含一个init函数，该函数会在main函数执行前，被go运行框架调研，也就是说init会在main函数前被调用。

**注意**

- 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是变量定义->init函数->main函数
- init函数主要用于初始化工作

### 匿名函数

如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用

 - 使用方式1

在定义匿名函数时就直接调用，这种调用只能使用一次

````go
package main
import(
	"fmt"
) 
func main(){
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10,20)

	fmt.Println("res1=", res1)//res1= 30
}
````

- 使用方式2

将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数。

````go
    a := func (n1 int, n2 int) int {
		return n1 + n2
	}

	res2 := a(10, 30)

	fmt.Println("res2=", res2)//res2= 40
````


- 全局匿名函数

如果将匿名函数赋给一个全局变量，那么这个匿名函数就成为一个全局匿名函数，可以在程序有效

````go
    var (
        Fun1 = func (n1 int, n2 int) int {
            return n1 * n2
        }
    )

	res4 := Fun1(10, 30)
	fmt.Println("res4=", res4)//res4= 300
````

### 闭包

闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）

````go
package main
import(
	"fmt"
) 

func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}
func main(){
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	// str= helloa
	// 11
	// str= helloaa
	// 13
	// str= helloaaa
	// 16

}
````

        var n int = 10
        var str = "hello"
        return func (x int) int {
            n = n + x
            str += "a"
            fmt.Println("str=", str)
            return n
        }

- 此匿名函数和引用到函数外的n str形成一个整体，构成闭包
- 可以这样理解，闭包是类，函数时操作，n是字段。函数和他使用到的n str构成闭包
- 当我们返回调用f函数时，因为n是初始化一次，因此每调用一次就 进行累计

### defer

在函数中，程序员经常需要创建资源，比如数据库连接、文件句柄、锁等，为了在函数执行完毕后，及时的释放资源，go的设计者提供了defer（延时机制）

````go
package main
import(
	"fmt"
) 

func sum(n1 int,n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)

	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
func main(){
	res := sum(10,20)
	fmt.Println("res=", res)
	// ok3 res= 32
	// ok2 n2= 20
	// ok1 n1= 10
	// res= 32
}
````

**注意**

- 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
- 当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈，执行
- defer将语句放入到栈时，也会将相关的值拷贝同时入栈。

### 函数参数的传递方式

**值类型**

基本数据类型int系列，float系列，bool，string、数组和结构体struct

**引用类型**

指针、slice切片、map、管道chan、interface等

#### 变量作用域

- 函数内部声明/定义的变量叫局部变量，作用于仅限于函数内部
- 函数外部声明/定义的变量叫全局变量，作用于在整个包都有效。如果首字母是大写，则作用域在整个程序有效

## 3.3 常用函数

### go字符串函数

- 统计字符串长度，按字节len(str)
- 字符串便利，同时处理有中文的问题r := []rune(str)
- 字符串转整数：n,err := strconv.Atoi("12")
- 整数转字符串str = strconv.Itoa(123456)
- 字符串转[]byte:var byte = []byte("hello go")
- []byte转字符串：str = string([]byte{97,98,99})
- 10进制转2 8 16进制：str = strconv.FormaInt(123,2)
- 查找子串是否在指定的字符串中：strings.Contains("ceheee","e")
- 统计一个字符串有几个指定的子串：strings.Count("ceheee","e")
- 不区分大小写的字符串比较：fmt.Println(strings.EuqalFold("abc","Abc))//true
- 返回子串在字符串中第一个出现的index值，如果没有返回-1:strings.Index("Nlsabc","abc")//4
- 返回子串在字符串最后一次出现的index，如果没有返回-1：strings.LastIndex("golang","go")
- 将指定的子串替换成另外一个子串：strings.Replace("gogohello","go","go语言",n),n指定替换几个，n=1表示全部替换
- 按照指定的某个字符，为分割表示，将一个字符串拆分为字符串数组：strings.Splot("helloworld,ok",",")
- 将字符串的字母进行大小写的转换：strings.ToLower("Go")转小写；strings.ToUpper("Go")转大写
- 将字符串作用两边的空格去掉：strings.TrimSpace(" tn m jhk  f ")
- 将字符串左右两边指定的字符去掉：strings.Trim("!hello!","!")
- 将字符串左边指定的字符去掉：strings.TrimLeft("!hello","!")
- 将字符串右边指定的字符去掉：strings.TrimRight("hello!","!")
- 判断字符串是否以指定的字符串开头：strings.HasPrefix("ftp://192.168.88.25","ftp")
- 判断字符串是否以指定的字符串结尾：strings.HasSuffix("adc.jpg","jpg")

### 时间和日期函数

需要导入time包

````go
package main
import(
	"fmt"
	"time"
) 
func main(){
	//获取当前时间
	now := time.Now();
	fmt.Printf("now=%v now type=%T\n", now, now)
	//now=2022-12-06 19:21:08.279572 +0800 CST m=+0.001794201 now type=time.Time

	//通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	// 年=2022
	// 月=December
	// 月=12
	// 日=6
	// 时=19
	// 分=25
	// 秒=58

	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	//当前年月日 2022-12-6 19:28:54

	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	//2022-12-06 19:30:52
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	//2022-12-06
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	//19:30:52
}
````

- 时间的常量

        const(
            Nanosecond Duration//纳秒
            Microsecond //微秒
            Millisecond //毫秒
            Second //秒 
            Minute //分钟 
            Hour //小时 
        )
        //获得1000毫秒
        1000 * time.Millisecond

- 休眠

        func sleep(d Duration)

        案例：time.Sleep(100 * time.Millisecond)//休眠100毫秒

- 获取时间戳：time.Now().Unix()

### 内置函数（buildin)

golang设计者为了编程方便，提供了一些函数可以直接使用，我们称为go的内置函数

- len:用来求长度
- new:用来分配内存，主要用来分配值类型，比如int、float32、struct...返回的是指针
- make:用来分配内存，主要用来分配引用类型
## 3.4 错误处理

### go错误处理机制

go中可以抛出一个panic异常，然后在defer中通过recover捕获这个异常，然后正常处理

````go
package main
import(
	"fmt"
) 

func test(){
	defer func(){
	err := recover()
	if err != nil{
		fmt.Println("err=",err)
	}

}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main(){
	test()
	fmt.Println("下面的代码和逻辑。。。")
	//err= runtime error: integer divide by zero
	//下面的代码和逻辑。。。
}
````

#### 自定义错误

- error.New("错误说明"),会返回一个error类型的值，表示一个错误
- panic内置函数，接收一个interface()类型的值（也就是任何值）作为参数。可以接收error类型的变量，输出错误信息，并退出程序

````go
package main
import(
	"fmt"
	"errors"
) 
func readConf(name string) (err error){
	if name == "config.ini" {
		//读取
		return nil
	}else{
		//返回一个自定义错误
		return errors.New("读取文件错误。。")
	}
}

func test(){
	err := readConf("config2.ini")
	if err != nil{
		//panic:捕获自定义异常并退出程序
		panic(err)
	}
	fmt.Println("程序正常运行")
}

func main(){
	test()
	fmt.Println("main程序正常运行")
	// panic: 读取文件错误。。

	// goroutine 1 [running]:
	// main.test()
	//         E:/know/笔记/goExec/src/test20221206/test08/main.go:20 +0x49
	// main.main()
	//         E:/know/笔记/goExec/src/test20221206/test08/main.go:26 +0x19
	// exit status 2
}
````

# 4. 数组和切片
 
 数组的定义：

        var 数组名 [数组大小]数据类型
        var a [5]int
        赋值 a[0] =1
            a[1] = 30

````go
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
````

**注意**

- 数组的地址可以通过数组名来获取 &intArr
- 数组的第一个元素的地址就是数组的地址
- 数组的各个元素的地址间隔是依据数组的类型决定，比如int65->8 int32->4 ...

#### 初始化数组的方式

    var numArr01 [3]int = [3]int{1,2,3}

    var numArr02 = [3]int{5,6,7}

    var numArr03 = [...]int{8,9,10}

    var numArr04 = [...]int{1:800,0:300,2:666}

#### for-range遍历

基本语法：

    for index,value := range array01{
        ...
    }

**说明**

- 第一个返回值index是数组的下标
- 第二个value是在该下标位置的值
- 他们都是仅在for循环内部可见的局部变量
- 遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线

**注意**

- go数组是值类型，默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响。如果想在其他函数中，修改原来的数组，可以使用引用传递（指针方式）

### 切片

**基本介绍**

- 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
- 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len(slice)都一样
- 切片的长度是可以变换的，因此切片是一个可以动态变化数组
- 切片基本语法：var 变量名 []类型  var a []int

#### 切片的使用

- 定义一个切片，然后让切片去引用一个数组
- 通过make创建切片

基本语法：

    var 切片名 []type = make([]type,len,cap)

    type:数据类型 len大小 cap:指定切片容量，可选

- 定义一个切片，直接就指定具体数组，使用原理类似make方式

**说明**

- 切片初始化时  var slice = arr   [startIndex:endIndex] (左闭右开)
- 切片初始化时仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
  - var slice = arr[o:end] 可以简写 var slice = arr[:end]
  - var slice = arr[start:len(arr)] 可以简写： var slice = arr[start:]
  - var slice = arr[0,len(arr)] 可以简写：var slice = arr[:]

- cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
- 切片定义完之后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
- 切片可以继续切片
- 使用append内置函数，可以对切片进行动态追加。append需要用新数组来接收，silce1 append另一个切片slice2时为append(silce1,slice2...)
  1. 切片append操作的底层原理分析：
  - 本质是对数组扩容
  - 如果原来的数组容量够的话直接追加，不够的话，go底层会创建一下新的数组newArr
  - 将slice原来包含的元素拷贝到新的数组newArr
  - slice重新引用到newArr
  - 注意newArr是在底层来维护的，程序员不可见

- copy(par1,par2):par1 par2为切片，将par2 copy到par1,超出par1长度的部分放弃copy

#### string和slice

- string底层是一个byte数组，因此string可以进行切片处理
- string是不可变的，也就是说不能通过str[0] = 'w'来修改字符串
- 如果需要修改字符串，可以先将string -> []byte或者 []rune 修改重写为string

### map

基本语法：

    var map变量名 map[keytype]valuetype

key可以是bool、数字、string、指针、channel、接口、结构体、数组，通常是int、string

注意：slice、map和function不可以，因为这几个没法用==判断

value 类型基本和key一样，通常为数字、string、map、struct

map声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。

````go
package main
import(
	"fmt"
) 
func main(){
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "aaa"
	a["no2"] = "bbb"
	a["no1"] = "ccc"
	fmt.Println(a)//map[no1:ccc no2:bbb]
}
````