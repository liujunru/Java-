## 1. 编程基础

  数据在计算机内部都是二进制表示的，不方便操作，为了方便操作数据，高级语言引入了**数据类型**和**变量**的概念
### 1.1 数据类型和变量

日期在Java中也是一个对象，内部表示为整型long

所谓内存在程序看来就是一块有地址编号的连续的空间，数据放到内存中的某个位置后，为了方便地找到和操作这个数据，需要给这个位置起一个名字。编程语言通过**变量**这个概念来表示这个过程。

在给long类型赋值时，如果常量超过了int的表示范围，需要在常量后面加大写或小写字母L，即L或l，之所以需要加L或l，是因为数字常量默认为是int类型。

日期在Java中也是一个对象，内部表示为整型long

### 1.2 赋值

一个基本类型变量，内存中只会有一块对应的内存空间。但数组有两块：一块用于存储数组内容本身，另一块用于存储内容的位置

### 1.3 条件执行

switch值的类型可以是byte、short、int、char、枚举和String,但不能是long类型，因为跳转表大部分是32位存储

### 1.4 函数调用的基本原理

栈一般是从高位地址向低位地址扩展，换句话说，栈底的内存地址是最高的，栈顶的是最低的。

计算机系统主要使用栈来存放函数调用过程中需要的数据，包括参数、返回地址，以及函数内定义的局部变量。

对于对象和数组，他们都有两块内存，一块存在实际的内容，一块存放实际内存的地址，实际的内容空空一般不是分配在栈上，而是在**堆**，存放地址的空间是分配在**栈**上。

函数调用主要通过栈来存储相关的数据，系统就函数调用者和函数如何使用栈做了约定，返回值可以简单认为是通过一个专门的返回器存储器存储的。

## 2. 理解数据背后的二进制

### 2.1 整数的二进制表示与位运算

整数有4种类型byte short int long,分布占 1 2 4 8个字节，即分别占8 16 32 64位，每种类型的符号位都是其最左边的一位。整数是原码表示法，负数时补码表示法，即在原码表示的基础上取反再加1。

```java
int a = 25;
sout(Integer.toBinaryString(a));//二进制
sout(Integer.toHexString(a));//十六进制
```

### 2.2 位运算

1) 左移：<<,向左移动，右边的低位补0，高位舍弃，相当于乘以2.
2) 无符号右移：>>>，向右移动，右边的舍弃掉，左边补0.
3) 有符号右移：>>，向右移动，右边的舍弃掉，左边补什么取决于原来最高位是什么，相当于除以2

```java
int a = 4;
a = a>> 2;//001，等于1
a = a << 3;//1000,=8
```

### 2.3 char的真正含义

char本质上是一个固定占用两个字节的无符号正整数，这个正整数对应于unicode编号，用于表示那个Unicode编号对应的字符。

# 3. 类的基础

## 3.1 类的基本概念

**Arrays api**

void sort(int[] a):int数组升序排序

int binarySearch(long[] a,long key):二分查找key，数组a需要升序排序

void fill(int[] a,int val):给所有数组元素赋值val

int copyOf(int[] original,int newLength):数组复制

boolean equals(char[] a,char[] a2):判断两个数组是否相同

### 3.1.1 类、方法、变量

类方法只能访问类变量，不能访问实例变量，可以调用其他的类方法，不能调用实例方法。

实例方法既能访问实例变量，也能访问类变量。既可以调用实例方法，也可以调用类方法。

**私有化变量的原因：**

一般而言，不应该将实例变量声明为Public，而值应该通过对象的方法对实例变量进行操作，这也是为了减少误操作，直接访问变量没有办法进行参数检查和控制，而通过方法修改，可以在方法中进行检查。

### 3.1.2 构造方法

**私有化构造方法的场景**

1.  不能创建类的实例，类只能被静态访问，如Math和Arrays类，他们的构造方法就是私有的。
2. 能创建类的实例，但只能被类的静态方法调用，如单例模式
3. 只是用来被其他多个构造方法调用，用于减少重复代码。

### 3.1.3 类和对象的生命周期

对象和数组一样，有两块内存，保存地址的部分分配在栈中，而保存实际内容的部分分配在堆中。栈中的内存是自动管理的，函数调用入栈就会分配，而出栈就会释放。

堆中的内存是被垃圾回收机制管理的，当没有活跃变量指向对象的时候，对应的堆空间就可能被释放，具体释放实际是JVM决定的。活跃变量就是已加载的类的类变量，以及栈中所有的变量。

### 3.1.4 小结

通过**类**实现自定义数据类型，**封装**该类型的数据所具有的属性和操作，**隐藏实现细节**，从而在更高的层次（类和对象的层次，而非基本数据类型和函数的层次）上考虑和操作数据，是计算机程序解决复杂问题的一种重要思维方式。

# 4. 类的继承

## 4.1 基本概念

```java
public class ShapeManager{
  private static final int MAX_NUM = 100;
  private Shape[] shapes = new Shape[MAX_NUM];
  private int ShapeNum = 0;
  public void addShape(Shape shape){
    if(shapeNum < MAX_NUM>){
      shapes[shapeNum++] = shape;
    }
  }
  public void draw(){
    for(int i = 0; i < shhapeNum; i++>){
      shapes[i].draw();
    }
  }
}

  public static void main(String[] args) {
     ShapeManager manager = new ShapeManager();
     manager.addShape(new Cicrle(new Point(4,2)));
     manager.addShape(new Line(new Point(4,2)));
     manager.addShape(new ArrowLine(new Point(4,2)));
     manager.draw();
    }
```

  变量shape可以引用任何Shape子类类型的对象，这叫**多态**，即一种类型的变量，可以引用多种实际类对象。这样对于变量shape，他就有两个类型：类型Shape,称之为shape的**静态类型**；类型Circle/Line/ArrowLine。我们称之为shape的**动态类型**。在ShapeManager的draw方法中，shape[i].draw()调用的是其对应动态类型的draw方法，这称之为方法的**动态绑定**。

  **多态和动态绑定**是计算机程序的一种重要思维方式，使得操作对象的程序不需要关注对象的实际类型，从而可以从统一处理不同对象，但又能实现每个对象的特有行为。

## 4.2 继承的细节

静态绑定在程序编译阶段即可决定，而动态绑定则要等到程序运行时。实例变量、静态变量、静态方法、private方法，都是静态绑定的。

父类方法定义实现的模板，具体实现由子类提供，这种应用位模板方法。模板方法在很多框架中有广泛的应用，这是使用protected的一种常见场景。

重写时，子类方法不能降低父类方法的可见性，子类方法可以升级父类方法的可见性。

## 4.3 继承实现的基本原理

类加载过程：
  1. 分配内存保存类的信息
  2. 给类变量赋默认值
  3. 加载父类
  4. 设置父子关系
  5. 执行类初始化代码

创建对象的过程包括：
  1. 分配内存
  2. 对所有实例变量赋默认值
  3. 执行实例初始化代码。

动态绑定实现的机制就是根据对象的实际 类型查找要执行的方法，子类型中找不到的时候再查找父类。

如果继承的层次比较深，需要调用的方法位于比较上层的父类，而调用的的效率时比较低的，因为每次调用都要进行多次查找。大多数系统使用一种称为**虚方法表**来优化调用的效率。

所谓虚方法表，就是在类加载的的时候为每个类创建一个表，记录该类的对象所有**动态绑定的方法（包括父类的方法）及其地址**，但一个方法只有一条记录，子类重写了父类方法后只会保留**子类**的。

对变量的访问是**静态绑定**的，无论是类变量还是实例变量。

继承具有破坏力。因为一方面继承可能破坏封装，封装是程序设计的第一原则。子类和父类直接可能存在着实现细节的依赖。子类在继承父类的时候，往往不得不关注父类的实现字节，而父类在修改其内部实现，如果不考虑子类，也往往会影响子类。另一方面，继承可能没有反映出is-a关系。

**正确使用继承**

1. 如果基类是别人写的，子类是我们自己写的

   - 重写方法不要改变预期的行为
   - 阅读文档说明，理解可重新方法的实现机制，尤其是方法之间的依赖关系
   - 在基类修改的情况下，阅读其修改说明，相应修改子类。

2. 基类是自己写的，子类是别人写的
   
    - 使用继承反映真正的额is-a关系，只将真正公共的部分放到基类。
    - 对不希望被重写的公开方法添加final修饰符
    - 写文档，说明可重写方法的实现机制，为子类提供指导，告诉子类应该如何重写。
    - 在基类修改可能影响子类时，写修改说明。

# 5. 类的扩展

## 5.1 接口的本质

接口声明了一组能力，但他自己没有实现这个能力，他只是一个约定。接口设计交互两方对象，一方需要实现这个接口，一方使用这个接口，但双方对象并不直接相互依赖，他们只是通过接口间接交互。

与类一样，接口也可以使用instanceof关键字，判断一个对象是否实现了某接口

Java8允许在接口中定义两类新方法：静态方法和默认方法

```java
public interface Idemo{
  void hello();
  public static void test(){
    sout("hello");
  }
  default void hi(){
    sout("hi");
  }
}
```

**针对接口编程**是一种重要的程序思维方式，这种方法不仅可以复用代码，还可以降低耦合，提高灵活性，是分解复杂问题的一种重要工具。

## 5.2 抽象类

#### 为什么需要抽象类

使用抽象方法而非空方法体，子类就知道它必须要实现该方法，而不可能忽略，若忽略Java编译器会提示错误。使用抽象类，类的使用者创建对象的时候，就知道必须要使用某个具体子类，而不可能误用不完整的父类。

抽象类可以减少误用。抽象类经常和接口配合使用，接口定义能力，抽象类提供默认实现，方便子类实现接口。

## 5.3 内部类

内部类与包含它的外部类有比较密切的关系，而与其他类关系不大，定义在类内部，可以实现对外部完全隐藏，可以有更好的封装性，代码实现上也往往更为简洁。

每个内部类最后都会被编译成一个独立的类，生成一个独立的字节码文件。

### 5.3.1 静态内部类

```java
public class Outer {
    private static int shared = 100;
    public static class StaticInner{
        public void innerMethod(){
            System.out.println("inner" + shared);
        }
        
    }
    public void test(){
        StaticInner si = new StaticInner();
        si.innerMethod();
    }
}
```
 静态内部类与外部类的联系也不大（与其他内部类相比）。它可以访问外部类的**静态变量和方法**，如innerMethod直接访问shared变量，但不可以访问**实例变量和方法**。在类内部，可以直接使用内部静态类，如test()所示。

 public静态内部类可以被外部使用，只是需要通过“外部类.静态内部类”的方式使用
 ```java
 Outer.SatticInner si = new Outer.SatticInner();
 si.innerMethod();
 ```

 静态内部类的使用场景是很多的，如果它与外部类关系密切，且不依赖于外部类实例，则可以考虑定义为静态内部类。

 Java Api中使用静态内部类的例子：

  - Integer类内部有一个私有静态内部类IntegerCache，用于支持整数的自动装箱
  - 表示链表的Linked List类内部Node，表示链表中的每个节点
  - Character类内部有一个Public静态内部类UnicodeBlock，用于表示一个Unicode block。

### 5.3.2 成员内部类

```java
public class Outer {
    private  int a = 100;
    public  class Inner{
        public void innerMethod(){
            System.out.println("outer a" + a);
        }

    }
    private void action(){
        System.out.println("action");
    }
    public void test(){
       Inner inner = new Inner();
       inner.innerMethod();
    }
}
```

Inner就是成员内部类，于静态内部类不同，除了**静态变量和方法**，成员内部类还可以直接访问外部类的**实例变量和方法**。

成员内部类对象总是于一个外部类对象相连的，在外部使用时，它要先创建一个Outer类对象，如：

```java
Outer outer = new Outer();
OUter.Inner inner = outer.new Inner();
inner.innerMethod();
```

**成员内部类中不可以定义静态变量和方法（final变量除外），方法内部类和匿名内部类也不可以**，原因：

这些内部类是与外部实例相连的，不应独立使用，而静态变量和方法作为类型的数属性和方法，一般是独立使用的，在内部类中意义不大，而如果内部类确实需要静态变量和方法，那么也可以挪到外部类中。

**应用场景**

成员内部类有哪些应用场景呢？如果内部类和外部类关系密切，需要访问外部类的实例变量或方法，则可以考虑定义为成员内部类。外部类的一些方法的返回值可能是u某个接口，为了返回这个接口，外部类方法可能使用内部类实现这个接口，这个内部类可以被设为private，对外完全隐藏。

Java Api：

LinkedList,他的两个方法listIterator和descendingIterator的返回值都是接口Iterator，调用者可以通过Iterator接口对链表遍历，这两个内部类都实现了接口Iterator。

### 5.3.3 方法内部类

```java
public class Outer {
    private  int a = 100;
    public void test(final int param){
        final String str = "hello";
        class Inner{
            public void innerMethod(){
                System.out.println("outer a" + a);
                System.out.println("param" + param);
                System.out.println("local str" + str);
            }

        }
        Inner inner = new Inner();
        inner.innerMethod();
    }
}
```

类Inner定义在外部类方法test中，方法内部类只能在定义的方法内被使用。如果方法是实例方法，则除了静态变量和方法，内部类还可以直接访问外部类的实例变量和方法，如innerMethod直接访问了外部私有实例变量a。如果方法是静态方法，则方法内部类只能访问外部类的静态变量和方法。方法内部类还可以直接访问方法的参数和方法中的局部变量。

**为什么方法内部类访问外部方法中的参数和局部变量时，这些变量必须声明为final**

  - 方法内部类操作的并不是外部的变量，而是他自己的实例变量，只是这些变量的值和外部一样，对这些变量赋值，并不会改变外部的值，为避免混淆，所以声明为final。java8之后不需要加final 也不会报错。java 加了一个语法糖，声明那个变量时可以不加final修饰，但是效果和加了是一样的，同时也要按照final变量的方式使用它，不然会报错。

### 5.3.4 匿名内部类

```java
public class Outer {
   public void test(final int x,final int y){
       Point p = new Point(2,3){
           @Override
           public double distance() {
               return distance(new Point(x,y));
           }
       };
       System.out.println(p.distance());
   }
}
```
匿名内部类是与new关联的，在创建对象的时候定义类，new后面是父类或者父接口，然后是圆括号（），里面是可以传递给父类构造方法的参数，最后是大括号{}，里面是类的定义。

匿名内部类只能被使用一次，用来创建一个对象。他没有名字，没有构造方法，但可以根据参数列表，调用对应的父类构造方法。与方法内部类一样，匿名内部类可以访问外部类的所有变量和方法，可以访问方法中的final参数和局部变量。

匿名内部类可以做的，方法内部类都能做做。但如果对象只创建一次，且不需要构造方法来接受参数，则可以使用匿名内部类，这样代码书写上更简洁。

将程序分为保持不变的主体框架和针对具体情况的可变逻辑，通过回调的方式进行协作，是计算机程序的一种常用实践。匿名内部类是实现回调接口的一种简便方式。

## 5.4.0 枚举的本质

枚举值是有顺序的，可以比较大小。枚举类型都有一个方法int ordinal()，表示枚举值在声明时的顺序，从0开始。

枚举类型都有一个静态的valueOf(String)方法，可以返回字符串对应的枚举值，例如，以下代码输出为true：

```java
Size.SMALL = Size.valueOf("SMALL");
```

枚举类型也都有一个静态的values方法，返回一个包括所有枚举值的数组。

枚举的好处体现在以下几方面。

  ❑ 定义枚举的语法更为简洁。

  ❑ 枚举更为安全。一个枚举类型的变量，它的值要么为null，要么为枚举值之一，不可能为其他值，但使用整型变量，它的值就没有办法强制，值可能就是无效的。

  ❑ 枚举类型自带很多便利方法（如values、valueOf、toString等），易于使用。

枚举值的定义需要放在最上面，枚举值写完之后，要以分号（; ）结尾，然后才能写其他代码。

枚举还有一些高级用法，比如，每个枚举值可以有关联的类定义体，枚举类型可以声明抽象方法，每个枚举值中可以实现该方法，也可以重写枚举类型的其他方法。此外，枚举可以实现接口，也可以在接口中定义枚举。

## 6. 异常类

![java异常类体系](https://gitee.com/liujunrull/image-blob/raw/master/202206131057457.png)

# 7.常用基础类

## 7.1 包装类

Float有一个静态方法floatToIntBits()，将float的二进制表示看作int。需要注意的是，只有两个float的二进制表示完全一样的时候，equals才会返回true。

包装类都是**不可变类**。所谓不可变是指实例对象一旦创建，就没有办法修改了。这是通过如下方式强制实现的：

  ❑ 所有包装类都声明为了final，不能被继承。
  
  ❑ 内部基本类型值是私有的，且声明为了final。
  
  ❑ 没有定义setter方法。
  
  为什么要定义为不可变类呢？不可变使得程序更为简单安全，因为不用操心数据被意外改写的可能，可以安全地共享数据，尤其是在多线程的环境下。

  IntegerCache表示Integer缓存，其中的cache变量是一个静态Integer数组，在静态初始化代码块中被初始化，默认情况下，保存了-128～127共256个整数对应的Integer对象。在valueOf代码中，如果数值位于被缓存的范围，即默认-128～127，则直接从Integer-Cache中获取已预先创建的Integer对象，只有不在缓存范围时，才通过new创建对象。通过共享常用对象，可以节省内存空间，由于Integer是不可变的，所以缓存的对象可以安全地被共享。Boolean、Byte、Short、Long、Character都有类似的实现。这种共享常用对象的思路，是一种常见的设计思路，它有一个名字，叫**享元模式**，英文叫Flyweight，即共享的轻量级元素。

  ## 7.2 String

  String类的hashCode方法，代码如下：

```java
  public int hashCode() {
        int h = hash;
        if (h == 0 && value.length > 0) {
            char val[] = value;

            for (int i = 0; i < value.length; i++) {
                h = 31 * h + val[i];
            }
            hash = h;
        }
        return h;
    }
```

## 7.3 StringBuilder

new StringBuilder()代码内部会创建一个长度为16的字符数组，count的默认值为0

默认长度为16，长度不够时，会先扩展到16*2+2即34，然后扩展到34*2+2即70，然后是70*2+2即142，这是一种指数扩展策略。为什么要加2？这样，在原长度为0时也可以一样工作。为什么要这么扩展呢？这是一种折中策略，一方面要减少内存分配的次数，另一方面要避免空间浪费。在不知道最终需要多长的情况下，**指数扩展**是一种常见的策略，广泛应用于各种内存分配相关的计算机程序中。

insert方法的实现代码：
```java
    public AbstractStringBuilder insert(int index, char[] str, int offset,
                                        int len)
    {
        if ((index < 0) || (index > length()))
            throw new StringIndexOutOfBoundsException(index);
        if ((offset < 0) || (len < 0) || (offset > str.length - len))
            throw new StringIndexOutOfBoundsException(
                "offset " + offset + ", len " + len + ", str.length "
                + str.length);
        ensureCapacityInternal(count + len);
        System.arraycopy(value, index, value, index + len, count - index);
        System.arraycopy(str, offset, value, index, len);
        count += len;
        return this;
    }

```

String可以直接使用+和+=运算符，这是Java编译器提供的支持，背后，Java编译器一般会生成StringBuilder, +和+=操作会转换为append

传递比较器Comparator给sort方法，体现了程序设计中一种重要的思维方式。将不变和变化相分离，排序的基本步骤和算法是不变的，但按什么排序是变化的，sort方法将不变的算法设计为主体逻辑，而将变化的排序方式设计为参数，允许调用者动态指定，这也是一种常见的设计模式，称为策略模式，不同的排序方式就是不同的策略。

二分查找既可以针对基本类型数组，也可以针对对象数组，对对象数组，也可以传递Comparator，也可以指定查找范围。如果能找到，binarySearch返回找到的元素索引。如果没找到，返回一个负数，这个负数等于-（插入点+1）。插入点表示，如果在这个位置插入没找到的元素，可以保持原数组有序。

多维数组只是一个假象，只有一维数组，只是数组中的每个元素还可以是一个数组，这样就形成二维数组。

```java
    private static int binarySearch0(long[] a, int fromIndex, int toIndex,
                                     long key) {
        int low = fromIndex;
        int high = toIndex - 1;

        while (low <= high) {
            int mid = (low + high) >>> 1;
            long midVal = a[mid];

            if (midVal < key)
                low = mid + 1;
            else if (midVal > key)
                high = mid - 1;
            else
                return mid; // key found
        }
        return -(low + 1);  // key not found.
    }
```

# 8. 泛型与容器



















