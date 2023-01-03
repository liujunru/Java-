### Go 语言基础语法

#### 关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
![](https://gitee.com/liujunrull/image-blob/raw/master/202211141621379.png)

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

![](https://gitee.com/liujunrull/image-blob/raw/master/202211141622486.png)

### go语言变量

指定变量类型时，如果没有初始化，则变量默认为**零值**

- 数值类型（包括complex64/128）为 0

- 布尔类型为 false

- 字符串为 ""（空字符串）

- 以下几种类型为 nil：

````go
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口
````


````go
var f string = "Runoob" 简写为 f := "Runoob"：
````


你可以通过&i来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。

### go语言常量

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 **const**关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

### go语言运算符

![](https://gitee.com/liujunrull/image-blob/raw/master/202211142128518.png)

Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

### go语言循环语句


for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
````go
for key, value := range oldMap {
    newMap[key] = value
}
````

以上代码中的 key 和 value 是可以省略。

如果只想读取 key，格式如下：
````go
for key := range oldMap
//或者
for key, _ := range oldMap
````

如果只想读取 value，格式如下：

````go
for _, value := range oldMap
````

以下为 Go 语言嵌套循环的格式：
````go
for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
//以下实例使用循环嵌套来输出 2 到 100 间的素数：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var i, j int

   for i=2; i < 100; i++ {
      for j=2; j <= (i/j); j++ {
         if(i%j==0) {
            break; // 如果发现因子，则不是素数
         }
      }
      if(j > (i/j)) {
         fmt.Printf("%d  是素数\n", i);
      }
   }  
}
````
#### goto
Go 语言的 goto 语句可以无条件地转移到过程中指定的行。

goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。

````go
//goto 语法格式如下：
goto label;
..
.
label: statement;


//在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++    
   }  
}
````

### go语言函数

#### 值传递
 ````go
/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int

   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/

   return temp;
}

package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200

   fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )

   /* 通过调用函数来交换值 */
   swap(a, b)

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}

/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int

   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/

   return temp;
}
交换前 a 的值为 : 100
交换前 b 的值为 : 200
交换后 a 的值 : 100
交换后 b 的值 : 200
 ````

 #### 引用传递

 ````go
/* 定义交换值函数*/
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保持 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}

package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)

   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}

交换前，a 的值 : 100
交换前，b 的值 : 200
交换后，a 的值 : 200
交换后，b 的值 : 100
 ````

 ### go语言变量作用域

 Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
 ````go
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}

结果： g = 10
 ````

 #### 初始化局部和全局变量

不同类型的局部和全局变量默认值为：

 ![](https://gitee.com/liujunrull/image-blob/raw/master/202211161031500.png)

 ### Go 语言数组

 #### 初始化数组

 ````go 
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//也可以通过字面量在声明数组的同时快速初始化数组：
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数
//自行推断数组的长度：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
//如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
//初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
 balance[4] = 50.0
 ````

 #### 初始化二维数组
多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：
 ````go
a := [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11},   /* 第三行索引为 2 */
}
 ````

 多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：

 ````go
a := [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}}   /* 第三行索引为 2 */
 ````

 ### Go 语言指针
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
````go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}
````

### Go 语言结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

````go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}


func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
````

### Go 语言切片(Slice)

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

#### 定义切片

你可以声明一个未指定大小的数组来定义切片：

````go
var identifier []type
````

切片不需要说明长度。

或使用 make() 函数来创建切片:

````go
var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
````

也可以指定容量，其中 capacity 为可选参数。

````go
make([]T, length, capacity)
````

这里 len 是数组的长度并且也是切片的初始长度。

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

````go
package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

len=3 cap=5 slice=[0 0 0]
````

一个切片在未初始化之前默认为 nil，长度为 0

#### 切片截取

可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：

````go
package main

import "fmt"

func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}  
   printSlice(numbers)

   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers)

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
````

####  append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
````go
package main

import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)  
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
````

### Go 语言范围(Range)

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

````go
for key, value := range oldMap {
    newMap[key] = value
}
````

### Go 语言类型转换

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：

````go
type_name(expression)
//实例
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
````

### Go 并发

Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：
````go
go 函数名( 参数列表 )

//例如
go 函数名( 参数列表 )
````

### 通道（channel）
通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
````go
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
````
声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

````go
ch := make(chan int)
````
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

#### 通道缓冲区
通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
````go
ch := make(chan int, 100)
````
带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于**异步**状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。