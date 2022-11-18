## Database
- 删除重复数据保留一条sql

```mysql
delete from testdelete where id not in (select a.id from(select min(id) as id from testdelete group by testname)as a);
```
- sql函数
```mybatis
LOCATE(substr,str) LOCATE(substr,str,pos):第一个语法返回字符串str第一次出现的子串substr的位置。

第二个语法返回第一次出现在字符串str的子串substr的位置，从位置pos开始。 substr不在str中，则返回0。

decode(columnname，值1,翻译值1,值2,翻译值2,...值n,翻译值n,缺省值):将查询结果翻译成其他值
```

## coreJava

- StringBuffer insert(int offset, String str):从offset开始插入str
- arraycopy(Object src,int srcPos,Object dest,int destPos,int length)：数组复制。src表示源数组，srcPos表示源数组要复制的起始位置，desc表示目标数组，length表示要复制的长度。

## 基本数据类型-boolean

Java语言表达式所操作的boolean值，在编译之后都使用Java虚拟机中的**int**数据类型来代替，而boolean数组将会被编码成Java虚拟机的**byte**数组，每个boolean元素占8位。这样我们可以得出boolean类型占了**单独使用是4个字节**，在**数组中又是1个字节**。使用int的原因是，对于当下32位的处理器（CPU）来说，一次处理数据是32位（这里不是指的是32/64位系统，而是指CPU硬件层面），具有高效存取的特点

## java自动装箱与拆箱

装箱就是自动将基本数据类型转换为包装器类型（int-->Integer）；调用方法：Integer的valueOf(int) 方法拆箱就是自动将包装器类型转换为基本数据类型（Integer-->int）。调用方法：Integer的intValue方法
````java
public static Integer valueOf(int i) {
if(i >= -128 && i <= IntegerCache.high)
return IntegerCache.cache[i + 128];
else
return new Integer(i);
}

private static class IntegerCache {
static final int high;
static final Integer cache[];
static {
final int low = -128;
// high value may be configured by property
int h = 127;
if (integerCacheHighPropValue != null) {
// Use Long.decode here to avoid invoking methods that
// require Integer's autoboxing cache to be initialized
int i = Long.decode(integerCacheHighPropValue).intValue();
i = Math.max(i, 127);
// Maximum array size is Integer.MAX_VALUE
h = Math.min(i, Integer.MAX_VALUE - -low);
}
high = h;
cache = new Integer[(high - low) + 1];
int j = low;
for(int k = 0; k < cache.length; k++)
cache[k] = new Integer(j++);
}
private IntegerCache() {}
}
````

从这2段代码可以看出，在通过valueOf方法创建Integer对象的时候，如果数值在[-128,127]之间，
便返回指向IntegerCache.cache中已经存在的对象的引用；否则创建一个新的Integer对象。
上面的代码中i1和i2的数值为100，因此会直接从cache中取已经存在的对象，所以i1和i2指向的是
同一个对象，而i3和i4则是分别指向不同的对象。
## final

被final修饰的类不可以被继承

被final修饰的方法不可以被重写

被final修饰的变量不可以被改变.如果修饰引用,那么表示引用不可变,引用指向的内容可变.

被final修饰的方法,JVM会尝试将其内联,以提高运行效率

被final修饰的常量,在编译阶段会存入常量池中.

**a+=b运算会进行隐性类型自动转换**

## try catch 里的finally和return

1、不管有木有出现异常，finally块中代码都会执行；

2、当try和catch中有return时，finally仍然会执行；

3、finally是在return后面的表达式运算后执行的（此时并没有返回运算后的值，而是先把要返回的值保存起来，管finally中的代码怎么样，返回的值都不会改变，任然是之前保存的值），所以函数返回值是在finally执行前确定的；

4、finally中最好不要包含return，否则程序会提前退出，返回值不是try或catch中保存的返回值。

## Java 创建对象有几种方式

- 使用 new 关键字，这也是我们平时使用的最多的创建对象的方式

- 使用反射方式创建对象，使用 newInstance()，但是得处理两个异常 InstantiationException、IllegalAccessException：
````java
User user=new User();User user=User.class.newInstance();
Object object=(Object)Class.forName("java.lang.Object").newInstance()
````

- 使用 clone 方法，前面题目中 clone 是 Object 的方法，所以所有对象都有这个方法。

- 使用反序列化创建对象，调用 ObjectInputStream 类的 readObject() 方法。
我们反序列化一个对象，JVM 会给我们创建一个单独的对象。JVM 创建对象并不会调用任何构造函数。一个对象实现了 Serializable 接口，就可以把对象写入到文件中，并通过读取文件来创建对象。

## HashMap 的长度为什么是 2 的 N 次方

为了能让 HashMap 存数据和取数据的效率高，尽可能地减少 hash 值的碰撞，也就是说尽量把数据能均匀的分配，每个链表或者红黑树长度尽量相等。
我们首先可能会想到 % 取模的操作来实现

下面是回答的重点哟：
取余（%）操作中如果除数是 2 的幂次，则等价于与其除数减一的与（&）操作（也就是说hash % length == hash &(length - 1) 的前提是 length 是 2 的 n 次方）。并且，采用二进制位操作 & ，相对于 % 能够提高运算效率。

这就是为什么 HashMap 的长度需要 2 的 N 次方了。

## JVM内存模型
![](https://gitee.com/liujunrull/image-blob/raw/master/202211171034916.png)

### 线程私有区
#### 1、程序计数器
记录下一条要运行的指令。如果执行的是JAVA方法，计数器记录正在执行的java字节码地址，如果执行的是native方法，则计数器为
空。
#### 2、虚拟机栈
线程私有的，与线程在同一时间创建。管理JAVA方法执行的内存模型。每个方法执行时都会创建一个桢栈来存储方法的的变量表、操作数栈、动态链接方法、返回值、返回地址等信息。栈的大小决定了方法调用的可达深度（递归多少层次，或嵌套调用多少层其他方法，-Xss参数可以设置虚拟机栈大小）。栈的大小可以是固定的，或者是动态扩展的。如果请求的栈深度大于最大可用深度，则抛出**stackOverflowError**；如果栈是可动态扩展的，但没有内存空间支持扩展，则抛出O**utofMemoryError**。 使用jclasslib工具可以查看class类文件的结构。
![](https://gitee.com/liujunrull/image-blob/raw/master/202211171036934.png)
#### 3、本地方法栈
与虚拟机栈作用相似。但它不是为Java方法服务的，而是本地方法（C语言）。由于规范对这块没有强制要求，不同虚拟机实现方法不同。

### 线程共享区
#### 1、方法区
线程共享的，用于存放被虚拟机加载的类的元数据信息，如常量、静态变量和即时编译器编译后的代码。若要分代，算是永久代（老年代），以前类大多“static”的，很少被卸载或收集，现回收废弃常量和无用的类。其中运行时常量池存放编译生成的各种常量。（如果hotspot虚拟机确定一个类的定义信息不会被使用，也会将其回收。回收的基本条件至少有：所有该类的实例被回收，而且装载该类的ClassLoader被回收）
#### 2、堆
存放对象实例和数组，是垃圾回收的主要区域，分为新生代和老年代。刚创建的对象在新生代的Eden区中，经过GC后进入新生代的S0区中，再经过GC进入新生代的S1区中，15次GC后仍存在就进入老年代。这是按照一种回收机制进行划分的，不是固定的。若堆的空间不够实例分配，则OutOfMemoryError。

## 什么时候会触发FullGC
除直接调用System.gc外，触发Full GC执行的情况有如下四种。

 1. 旧生代空间不足 旧生代空间只有在**新生代对象转入及创建为大对象、大数组**时才会出现不足的现象，当执行Full GC后空间仍然不足，则抛出如下错误： java.lang.OutOfMemoryError: Java heap space 为避免以上两种状况引起的FullGC，调优时应尽量做到让对象在Minor GC阶段被回收、让对象在新生代多存活一段时间及不要创建过大的对象及数组。
2. Permanet Generation空间满 PermanetGeneration中存放的为一些class的信息等，当系统中要**加载的类、反射的类和调用的方法较多**时，Permanet Generation可能会被占满，在未配置为采用CMS GC的情况下会执行Full GC。如果经过Full GC仍然回收不了，那么JVM会抛出如下错误信息： java.lang.OutOfMemoryError: PermGen space 为避免Perm Gen占满造成Full GC现象，可采用的方法为增大Perm Gen空间或转为使用CMS GC。
3. CMS GC时出现promotion failed和concurrent mode failure 对于采用CMS进行旧生代GC的程序而言，尤其要注意GC日志中是否有promotion failed和concurrent mode failure两种状况，当这两种状况出现时可能会触发Full GC。 promotionfailed是在进行Minor GC时，survivor space放不下、对象只能放入旧生代，而此时旧生代也放不下造成的；concurrent mode failure是在执行CMS GC的过程中同时有对象要放入旧生代，而此时旧生代空间不足造成的。 应对措施为：增大survivorspace、旧生代空间或调低触发并发GC的比率，但在JDK 5.0+、6.0+的版本中有可能会于JDK的bug29导致CMS在remark完毕后很久才触发sweeping动作。对于这种状况，可通过设XX:CMSMaxAbortablePrecleanTime=5（单位为ms）来避免。
4. 统计得到的Minor GC晋升到旧生代的平均大小大于旧生代的剩余空间 这是一个较为复杂的触发情况，Hotspot为了避免由于新生代对象晋升到旧生代导致旧生代空间不足的现象，在进行MinorGC时，做了一个判断，如果之前统计所得到的Minor GC晋升到旧生代的平均大小大于旧生代的剩余空间，那么就直接触发Full GC。 例如程序第一次触发MinorGC后，有6MB的对象晋升到旧生代，那么当下一次Minor GC发生时，首先检查旧生代的剩余空间是否大于6MB，如果小于6MB，
则执行Full GC。 当新生代采用PSGC时，方式稍有不同，PS GC是在Minor GC后也会检查，例如上面的例子中第一次Minor GC后，PS GC会检查此时旧生代的剩余空间是否大于6MB，如小于，则触发对旧生代的回收。 除了以上4种状况外，对于使用RMI来进行RPC或管理的Sun JDK应用而言，默认情况下会一小时执行一次Full GC。可通过在启动时通过- javaDsun.rmi.dgc.client.gcInterval=3600000来设置Full GC执行的间隔时间或通过-XX:+DisableExplicitGC来禁止RMI调用System.gc。

## Java程序是如何执行的
先把 Java 代码编译成字节码，也就是把 .java 类型的文件编译成 .class 类型的文件。这个过程的大致执行流程：Java 源代码 -> 词法分析器 -> 语法分析器 -> 语义分析器 -> 字符码生成器 ->最终生成字节码，其中任何一个节点执行失败就会造成编译失败；把 class 文件放置到 Java 虚拟机，这个虚拟机通常指的是 Oracle 官方自带的 Hotspot JVM；Java 虚拟机使用类加载器（Class Loader）装载 class 文件；类加载完成之后，会进行字节码效验，字节码效验通过之后 JVM 解释器会把字节码翻译成机器码交由操作系统执行。但不是所有代码都是解释执行的，JVM 对此做了优化，比如，以Hotspot 虚拟机来说，它本身提供了 JIT（Just In Time）也就是我们通常所说的动态编译器，它能够在运行时将热点代码编译为机器码，这个时候字节码就变成了编译执行。Java 程序执行
流程图如下
![](https://gitee.com/liujunrull/image-blob/raw/master/202211171724058.png)