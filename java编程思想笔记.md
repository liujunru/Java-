# java编程思想

## 对象导论

### 继承

在被继承类添加新方法时，应该仔细考虑是否存在基类也需要这些额外方法的可能性。减少迭代的多余代码量

#### 向上转型：

将导出类看作是他的基类的过程

![image-20220424213722943](https://gitee.com/liujunrull/image-blob/raw/master/202204242137997.png)

#### 单根继承结构

所有的类都继承自单一的基类-Object类。好处：

1）保证所有的对象都具备某些功能，object自带方法

2）所有对象都可以很容易在堆上创建，而参数传递也得到了极大的简化

3）使垃圾会受到额实现变得容易得多。垃圾回收器基于C++的改进，由于对象都保证具有类型信息，因此不会因无法确定对象类型陷入僵局。这对于系统级操作（如异常处理）显得尤为重要，也给变成带来了更大的灵活性。

## 一切都是对象

#### 程序运行时，存储对象的地方都有哪些？

- **寄存器**

  这是最快的存储区，这是最快的存储区，因为它位于不同于其他存储区的地方-处理器内部。数量极其有限，所以需要按需分配。

- **堆栈**

  位于通用Ram(随机访问存储器)中，但通过堆栈指针可以从处理器那里获得直接支持。基本数据类型

- **堆**

  一种通用的内存池（也位于RAM区），用于存放所有的Java对象。跟堆栈相比，编译器不需要知道存储的数据在堆里存活多长时间，比较灵活。但是用堆进行内存分配和清理可能比用堆栈进行存储分配需要的更多的时间。

- **常量存储**

  常量值存放。ROM中

- **非RAM存储**

  流对象和持久化对象。这种存储方式的技巧在于，把对象转化成可以存放在其他媒介上的事物，在需要时可恢复成常规的、基于RAM的对象。

#### Java中数组的安全性体现在哪里

Java确保数组会被初始化，而且不能在他的范围之外被访问。这种范围检查，是以每个数组上少量的内存开销以及运行时的下标检查为代价的。

 boolean默认false,char默认'\u0000'(null)

## 初始化与清理

#### 构造器

为了保证编译器知道该调用哪个方法进行初始化，构造器采用与类相同的名称。

#### 垃圾回收器如何工作

垃圾回收器的思想：对任何活的对象，一定能追溯到其存活的堆栈或者静态存储区之间的引用。

java虚拟机中内存分配以较大的“块”为单位。如果对象较大，就会占用单独的块。有了块之后，垃圾回收器在回收的时候就可以往废弃的块里老被对象了。每个块都用相应的代数来记录他是否还存活。通常，如果块在某处被引用，其代数会增加；垃圾回收器将对上次回收动作之后新分配的块进行整理。这对处理大量短命的临时对象很有帮助。垃圾回收器会定期进行完整的清理动作--大型对象仍然不会被复制（只是代数会增加），内含小型对象的那些块则被复制并整理。java虚拟机会进行监视，如果所有对象都很稳定，垃圾回收器的效率降低的话就切换到“标记--清扫”方式；同样java虚拟机会跟踪“标记--清扫”效果，要是堆空间出现很多碎片，就会切换回“停止--复制”方式。这就是“自适应、分代的、停止--复制、标记--清扫”式垃圾回收器。

##### 标记--清扫方式

从堆栈和静态存储区触发，遍历所有的引用，进而找出所有的存活对象。每当找到一个存活对象就对对象做一个标记。当全部标记完成之后，清理才会开始。清理时，没有标记的对象会被释放，不会发生任何肤质。所以剩下的堆空间是不连续的，垃圾回收器要是希望得到连续空间的话，就要重新整理剩下的对象

##### 停止--复制方式

垃圾回收的时候，程序会被暂停，然后将所有存活的对象从当前堆复制到另一个堆，没有复制的都是垃圾。当对象被复制到新堆时，他们是一个挨着一个，所以新堆保持紧凑排列。

#### 静态数据是何时被初始化的

静态初始化只有在必要时刻才会进行。第一次访问静态数据的时候才会被初始化，此后静态对象不会再被初始化。

初始化的顺序是先静态对象（如果之前未被初始化），然后是“非静态”对象。

#### 对象的创建过程

1. 假设有个Dog的类。即使没有显式的使用static关键字，构造器实际上也是静态方法。因此，当首次创建类型为Dog的对象时（构造器可以看成静态方法），或者Dog类的静态方法、静态域首次被访问时，Java解释器必须查找类路径，以定位Dog.class文件
2. 然后加载Dog.class，创建一个class对象，有关静态初始化的所有动作都会执行。因此，静态初始化只在class对象首次加载的时候进行一次
3. 当用new Dog()创建对象的时候，首先在堆上为Dog对象分配足够的存储空间
4. 这块存储空间会被清零，这就自动的将Dog对象中的所有基本类型设置为默认值，引用设为null
5. 执行所有出现于字段定义处的初始化动作
6. 执行构造器

#### 数组初始化

数字赋值是复制的引用

## 访问控制权限

#### java解释器的运行过程

首先找出环境变量CLASSPATH（可以通过操作系统来设置，有时也可以通过安装程序——用来在你的机器上安装java或者java工具）。classpath包含一个或多个目录，用来查找.calsspath文件的根目录。从根目录开始，解释器获取包的名称并将每个据点替换为反斜杠，以从classpath根中产生一个路径名称（浴室，package.foo.bar就变成了foo\bar\bar）。得到的路径会与classpath中的各个不同的项项连接，解释器就在这些目录中查询与你要创建的类名相关的.class文件



继承的执行顺序：从基类开始到继承类

垃圾清理的顺利：与清理动作生成顺序相反

##### 继承使用场景

是否需要从新类向基类进行向上转型。

## 多态

##### 动态绑定（后期绑定、运行时绑定）

除了static和final方法，java中其他方法都是后期绑定。运行时根据对象的类型进行绑定。

## 接口

##### 策略设计

创建一个能够根据所传递的参数对象不同而具有不同行为的方法。这类方法包含所要执行的算法中固定不变的部分，而“策略”包含变化的部分。策略就是要传递进去的参数对象，他包含要执行的代码。

## 内部类

##### 创建内部类对象

从外部类的非静态方法之外的任意位置创建某个**内部类**的对象：

1)OuterClassName.INnerClassName

2)使用new语法

```java
public class DotNew(){
    public class Inner{}
    public static void main(String[] args){
        DotNew dn = new DotNew();
        DotNew.Inner dbi = dn.new Inner();
    }
}
```

生成对**外部类对象**的引用，可以使用外部类的名字后面紧跟圆点和this

```java
public class DotThis{
    void f(){
        System.out.printlin("DotThis.f()");
    }
    public class Inner{
    public DotThis outer(){
        return DotThis.this;
    	}
    }
    public Inner inner(){
        return new Inner();
    }
    public static void main(String[] args){
        DotThis dt = new DotThis();
        DotThis.Inner dt1 = dt.inner();
        dt1.outer.f();
    }
}
//output:DotThis.f()

```

#### 嵌套类

如果不需要内部类对象与其外围类对象之间有联系，那么可以将内部声明为static：

1）要创建嵌套类的对象， 并不需要其外围类的对象

2）不能从嵌套类的对象中访问非静态的外围类对象

##### 接口内部的类

如果想要创建某些公共代码，使得他卖的可以被某个接口的所有不同实现所共用，那么可以使用接口内部的嵌套类

```java
public interface ClassTnterface{
    void howdy();
    class Test implements classInterface{
        public void howdy(){
            sout("hoddy");
        }
        public static void main(String[] args){
            new Test().howdy();
        }
    }
}
//Output:hoddy
```

#### 持有对象

HashSet最快的获取元素的方式，TreeSet按照比较结果的升序保存对象，LinkedHashSet按照被添加的顺序保存对象。

HashMap提供了最快的查找技术，TreeMap按照比较结果的升序保存键，LinkedHashMap按照插入顺序保存键，同时还保留了HashMap的查询速度。

##### ListIteraror

Iterator的子类型。可以双向移动，可以产生相对于迭代器在列表中指向的当前位置的前一个和后一个元素的索引，并且可以用set()替换他访问过的最后一个元素。可以通过调用listIterator()方法产生一个指向List开始处的ListIterator，并且可以通过调用listIterator(n)方法创建一个一开始就指向列表索引为n的元素处的ListIterator

##### LinkedList

getFirst() /element():返回列表的头（第一个元素），并不移除他，list为空报错。

peek() list为空时返回null

removeFirst() /remove():移除并返回列表的头，list为空报错

poll()list为空时返回null

#### Queue

LinkedList提供了方法以支持队列的行为，并且他实现了Queue接口，因此LinkedList 可以用作Queue的一种是实现。

##### PriorityQueue

优先级队列。调用offer()方法来插入一个对象时，这个对象会在队列中被排序。默认的排序将使用对象在队列中的自然顺序，但是你可以通过提供自己的Comparator来修改这个顺序。

### 字符串

##### stringBuilder循环

```java
StringBUilder result = new StringBUilder("[");
for(int i= 0; i < 25; i++){
    result.append(rand.nexTnt(100));
    result.append(",");
}
result.delete(result.length() - 1,result.length());
result.append("]");
sout(result);
//[56,33,45,67,78,9,...4,56,3,78,89]
```

### 类型信息

#### Class对象

Class.forName("全限定名"):创建class引用

Class getName:产生全限定的的类名

Class getSImpleName():产生不含包名的类名

Class getCanoincalName()：产生全限定的类名

Class.newInstance():创建虚拟构造器

#### 动态代理

通过调用静态方法Proxy.newProxyInstance()可以创建动态代理，这个方法需要得到一个类加载器，一个你希望该代理实现的接口列表，以及InvocationHandler接口的一个实现。

invoke()方法传递进来了代理对象，以防你需要区分请求的来源。在invoke()内部，在代理上调用方法时需要格外小心，因为对接口的调用将被重定向为对代理的调用。

通常，你会执行被代理的操作，然后使用Method.invoke()将请求转发给被代理对象，并传入必须的参数。

```java
public static void main(String[] args){
    SomeMethods proxy = (SomeMethods)Proxy.newProxyIntance(SomeMEthods.class.getClassLoader(),
                                                          new Class[]{SomeMethods.class},
                                                           new MethodSeletor(new Implementation()));
}
```

### 泛型

##### extends关键字

边界<T extends HasF>声明T必须具有类型HasF或者从HasF导出的类型

#### 散列码的原理
查询一个值的过程首先就是计算散列码，然后使用散列码查询数组。如果能够保证没有冲突（如果值得数量是固定得，那么就有可能），那就有了一个完美得散列函数，但是这种情况只是特例，通常，冲突由外部链接处理：数组并不直接保存值，而是保存值得list。然后对list中得值使用equals()方法进行线性得查询。这部分查询自然会笔记满，但是如果散列函数好的话，数组得每个位置只有较少得值。因此，不是查询整个list，而是快速跳到数组得某个位置，只对很少得元素进行比较，这就是hsahMap会如此快得原因。

##### 职责链模式

以多种不同的方式来解决一个问题，然后将他们链接在一起。当一个请求到来时，它遍历这个链，直到链中的某个解决方案能够处理该请求。

### 注解

每当你创建描述符性质的类或者接口时，一旦其中包含了**重复性**的工作，那就可以考虑使用**注解**来简化与自动化该过程

##### 元注解：负责注解其他的注解

|@Target|表示该注解可以用于什么地方。ConStructor:构造器声明。Field:域声明。（包括enum实例）。Loacl_variable:局部变量声明。Method:方法声明。package：包声明。parameter:参数声明。type；类、接口或enum声明|
|-------|--------------------|
|@Retention|表示需要在什么级别保存该注解信息。可选的参数包括Source:注解将被编译器丢弃。Class：注解在class文件中使用，但会被Vm丢弃。Runtime：JVM运行期也会保留此1注解，因此可以通过反射机制读取注解的信息
|@Documented|将此注解包含在javadoc中|
|@Inherited|允许子类继承父类中的注解|

##### 注解元素
@UserCase(id=47,descption="passwpord")
注解元素可用的类型包括

- 所有基本类型（int,float,boolean等）
- String
- Class
- enum
- Annotation
- 以上类型的数组

### 并发

实现并发最直接的方式是在操作系统级别使用**进程**，进程是运行在它自己的地址空间内的自包容的程序。
多任务操作系统可以通过**周期性**的将CPU将一个进程**切换**到另一个进程，来实现同时运行多个进程。
JAVA所使用的这种并发系统会共享诸如内存和I/O这样的资源，因此编写多线程程序最基本的困难在于**协调**不同线程驱动的任务之间对这些资源的使用，以使得这些资源不会**同时**被多个任务访问。
某些编程语言被设计为可以将并发任务彼此隔离，这些语言通常被称为函数型语言，其中每个函数调用不会产生任何副作用，并因此可以当作独立的任务来驱动，比如Erlang语言。当程序的某个部分必须使用大量的并发，并且你在试图构建这个部分时遇到了很多问题，可以考虑这种专门处理并发的语言。
Java采取了更加传统的方式，在顺序型语言的基础上提供对现成的支持。在与多任务操作系统中分叉外部程序的不同，线程机制是在由执行程序表示的单一进程中创建任务。这种任务产生的一个好处是操作系统的透明性。
Java的线程机制是**抢占式**的，这表示调度机制会周期性的中断线程，将上下文切换到另一个线程，从而为每一个线程都提供时间片，使得每个线程都会分配到数量合理的时间去驱动她的任务。在协作式系统中，每个任务都会自动的放弃控制，这要求程序员要有意识的在每个任务中插入某种类型的让步语句。协作式系统的优势是双重的：上下文切换的开销通常要比抢占式系统低廉的多，并且对可以同时执行的线程数量在理论上没有任何限制。

##### 使用Executor
java.util.concurrent包中的执行器将为你管理Thread对象，从而简化了并发编程。Executor在客户端和任务执行之间提供了一个间接层；与客户端直接执行任务不同，这个中介对象将执行任务。Executor允许管理异步任务的执行 ，而无需显式的管理线程的生命周期。

##### 线程池种类选择
CachedThreadPool在程序执行过程中通常会创建与所需数量相同的线程，然后再它回收旧线程的时候停止创建新线程，因此是合理的Executor的首选。只有当这种方式引发问题时，才需要切换到FixedThreadPool。
SingleThreadExecutor就像是线程数量为1的FixedThreadPool。这ui对于你希望在另一个线程连续运行的任何事物（长期存活的任务）来说非常有用。例如监听进入的套接字链接的任务。它对于希望在线程中运行的短任务也同样很方便。比如更新本地或远程日志的小任务，或者是事件分发线线程。他会序列化所有提交给它的任务，并会维护它自己（隐藏）的悬挂任务队列。通过序列化任务，你可以消除对序列化对象的需求。
**程序设计的基本目标**

将保持不变的事物和发生改变的事物相分离。可以使用策略模式，将会发生改变的代码封装成单独的类（策略对象），可以将策略对象传递给总是相同的代码。例如，用不同的对象来表示不同的比较方式，然后将他们传递给相同的排序代码
#### lis实用方法
* indexOfSubList(List source,List target):返回target在source中第一次出现的位置，找不到返回-1
* lastOfSubList(List source,List target)：返回最后一次出现的位置
* rotate(List,int distance):所有元素向后移动distance个位置，末尾的元素循环到前面来
* swap(List,int i,int j):交换list中位置i与位置j的元素。通常比自己写的代码快
* fill(List<? super T>,T x):用对象x替换list中所有元素
* frequency(Collection,Object x):返回Collection中等于x的元素个数

##### 设置Collection或Map为不可修改
```java
List<String> a = Collections.unmodifiableList(new ArrayList<String> data);
```
使用场景：将容器设为只读前，填入有意义的数据，装在数据后，使用“不可修改的”方法返回的引用替换原来的引用，这样就不用担心无意中修改了只读的内容。另一方面，此方法允许你保留一份可修改的容器，作为类的private成员，然后通过某个方法调用返回对该容器的“只读”的引用。这样一来，只有你可以修改容器的内容，而别人只能读取。
##### Collection或Map的自有同步控制
```java
Collection<String> c = Collections,synchronizedList(new ArrayList<String> data);
```
### IO系统
任何自Inputstream或Reader派生而来的类都含有名为read()的基本方法，用于读取单个字节或者字节数组。同样，任何自OutputStream或Writer派生而来的类都含有名为write()的基本方法，用于写单个字节或者字节数组。但是我们通常不会用到这些方法，他们之所以存在是因为别的类可以使用他们，以便提供更有用的接口。因此，我们很少使用单一的类来创建流对象，而是通过叠合多个对象来提供所期望的功能（这是装饰器设计模式）
#### 标准IO重定向
```java
PrintStream console = System.out;
BufferedInputStream in = new BufferedInputStream(new FileINputStream("readrectiong.java"));
PrintStream out = new PrintStream(new BUfferedOutputStream(new FileOutputStream("test.out")));
System.setIn(in);
System.setOut(out);
System.err(out);
BUfferedReader br = new BufferedReader(new InputStreamReader(System.in));
String s;
while((s = br.readLine()) != null)
sout(s);
out.close();
System.setOut(console);
```
这个程序将标准输入附接到文件上，并将标准输出和标准错误重定向到另一个文件。注意，他在程序开头处存储了队最初的System.out对象的引用，并且在结尾处将系统恢复到了该对象上。