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

## 8.1 解析通配符

<T extends E>和<? extends E>到底有什么关系？它们用的地方不一样，我们解释一下：

  1) <T extends E>用于定义类型参数，它声明了一个类型参数T，可放在泛型类定义中类名后面、泛型方法返回值前面。
  
  2) <? extends E>用于实例化类型参数，它用于实例化泛型变量中的类型参数，只是这个具体类型是未知的，只知道它是E或E的某个子类型。

形如DynamicArray<? >，称为无限定通配符。通配符形式更为简洁。虽然通配符形式更为简洁，但DynamicArray<? >和DynamicArray<T >都有一个重要的限制：只能读，不能写。

现在我们再来看泛型方法到底应该用通配符的形式还是加类型参数。两者到底有什么关系？我们总结如下。

  1) 通配符形式都可以用类型参数的形式来替代，通配符能做的，用类型参数都能做。
  
  2) 通配符形式可以减少类型参数，形式上往往更为简单，可读性也更好，所以，能用通配符的就用通配符。
  
  3) 如果类型参数之间有依赖关系，或者返回值依赖类型参数，或者需要写操作，则只能用类型参数。
  
  4) 通配符形式和类型参数往往配合使用。

泛型中的三种通配符形式<? >、<? super E>和<? extends E>，并分析了与类型参数形式的区别和联系，它们比较容易混淆，我们总结比较如下：

  1) 它们的目的都是为了使方法接口更为灵活，可以接受更为广泛的类型。
  
  2) <? super E>用于灵活**写入或比较**，使得对象可以写入父类型的容器，使得父类型的比较方法可以应用于子类对象，它不能被类型参数形式替代。
  3) <? >和<? extends E>用于灵活**读取**，使得方法可以读取E或E的任意子类型的容器对象，它们可以用类型参数的形式替代，但通配符形式更为简洁。

## 8.2 细节

在定义泛型类、方法和接口时，也有一些需要注意的地方，比如：

  ❑ 不能通过**类型参数创建对象**。
  
  ❑ 泛型类类型参数不能用于**静态变量和方法**。
  
  ❑ 了解多个类型限定的语法。


那如果确实希望根据类型创建对象呢？需要设计API接受类型对象，即Class对象，并使用Java中的**反射**机制。

Java中还支持多个上界，多个上界之间以&分隔。如果有上界类，类应该放在第一个，类型擦除时，会用第一个上界替换。

泛型与数组的关系：
  ❑ Java不支持创建泛型数组。
  
  ❑ 如果要存放泛型对象，可以使用**原始类型**的数组，或者使用**泛型容器**。
  
  ❑ 泛型容器内部使用Object数组，如果要转换泛型容器为对应类型的数组，需要使用**反射**。

# 9. 列表和队列

## 9.1 剖析ArrayList

### 9.1.1 基本方法
ArrayList的remove方法

```java
    public E remove(int index) {
        rangeCheck(index);

        modCount++;
        E oldValue = elementData(index);

        int numMoved = size - index - 1;//计算要移动的元素个数
        if (numMoved > 0)
            System.arraycopy(elementData, index+1, elementData, index,
                             numMoved);
        elementData[--size] = null; // clear to let GC do its work。将size-1,同时释放引用以便原对象被垃圾回收

        return oldValue;
    }
```

它也增加了modCount，然后计算要移动的元素个数，从index往后的元素都往前移动一位，实际调用System.arraycopy方法移动元素。elementData[--size] = null；这行代码将size减1，同时将最后一个位置设为null，设为null后不再引用原来对象，如果原来对象也不再被其他对象引用，就可以被垃圾回收。

### 9.1.2 迭代

#### Iterable和Iterator

❑ Iterable表示对象可以被迭代，它有一个方法iterator()，返回Iterator对象，实际通过Iterator接口的方法进行遍历；

❑ 如果对象实现了Iterable，就可以使用foreach语法；

❑ 类可以不实现Iterable，也可以创建Iterator对象。

listIterator()方法返回的迭代器从0开始，而listIterator(int index)方法返回的迭代器从指定位置index开始。

#### 在迭代的中间调用容器的删除方法

发生了并发修改异常，为什么呢？因为迭代器内部会维护一些索引位置相关的数据，要求在迭代过程中，容器不能发生结构性变化，否则这些索引位置就失效了。所谓结构性变化就是添加、插入和删除元素，只是修改元素内容不算结构性变化。

#### 可以使用迭代器的remove方法，可以避免此异常的原理

  - ArrayList中iterator方法的实现
```java
public Iterator<E> iterator() {
        return new Itr();
    }
```
  - 新建了一个Itr对象，Itr是一个成员内部类，实现了Iterator接口，声明为：

```java
private class Itr implements Iterator<E> 
```
    - 它有三个实例成员变量，为：

```java
int cursor;       // index of next element to return
int lastRet = -1; // index of last element returned; -1 if no such
int expectedModCount = modCount;
```

cursor表示下一个要返回的元素位置，lastRet表示最后一个返回的索引位置，expected-ModCount表示期望的修改次数，初始化为外部类当前的修改次数modCount，回顾一下，成员内部类可以直接访问外部类的实例变量。每次发生结构性变化的时候modCount都会增加，而每次迭代器操作的时候都会检查expectedModCount是否与modCount相同，这样就能检测出结构性变化。

  - 我们来具体看下，它是如何实现Iterator接口中的每个方法的，先看hasNext()，代码为：

```java
public boolean hasNext() {
            return cursor != size;
        }
```
  - cursor与size比较，比较直接，看next方法：

```java
  public E next() {
            checkForComodification();
            int i = cursor;
            if (i >= size)
                throw new NoSuchElementException();
            Object[] elementData = ArrayList.this.elementData;
            if (i >= elementData.length)
                throw new ConcurrentModificationException();
            cursor = i + 1;
            return (E) elementData[lastRet = i];
        }
```

   -  首先调用了checkForComodification，它的代码为：

```java
  final void checkForComodification() {
            if (modCount != expectedModCount)
                throw new ConcurrentModificationException();
        }
```
所以，next前面部分主要就是在检查是否发生了结构性变化，如果没有变化，就更新cursor和lastRet的值，以保持其语义，然后返回对应的元素。remove的代码为：

```java
     public void remove() {
            if (lastRet < 0)
                throw new IllegalStateException();
            checkForComodification();

            try {
                ArrayList.this.remove(lastRet);
                cursor = lastRet;
                lastRet = -1;
                expectedModCount = modCount;
            } catch (IndexOutOfBoundsException ex) {
                throw new ConcurrentModificationException();
            }
        }
```

它调用了ArrayList的remove方法，但同时更新了cursor、lastRet和expectedModCount的值，所以它可以正确删除。不过，需要注意的是，调用remove方法前必须先调用next

#### 迭代器的好处

迭代器表示的是一种**关注点分离**的思想，将数据的**实际组织方式**与数据的**迭代**遍历相分离，是一种常见的设计模式。需要访问容器元素的代码只需要一个Iterator接口的引用，不需要关注数据的实际组织方式，可以使用一致和统一的方式进行访问。

### 9.1.3 ArrayList实现的接口

Java 8对Collection接口添加了几个默认方法，包括removeIf、stream、spliterator等

#### RandomAccess

```java
public Interface RandomAccess{}
```
这种没有任何代码的接口在Java中被称为**标记接口**，用于声明类的一种属性。

主要用于一些通用的算法代码中，它可以根据这个声明而选择效率更高的实现。比如，Collections类中有一个方法binarySearch，在List中进行二分查找，它的实现代码就根据list是否实现了RandomAccess而采用不同的实现机制

```java
public static <T>
int  binarySearch(List<? extends Comparable<? super T>> list,T key){
  if(list instanceof RandomAccess || list.size()<BINARYSEARCH_THRESHOLD>)
  return Collections.indexedBinarySearch(list,key);
  else
  return Collections.iteratorBinarySearch(list,key);
}
```
Arrays中有一个静态方法asList可以返回对应的List,这个方法返回的List，它的实现类并不是本节介绍的ArrayList，而是Arrays类的一个内部类，在这个内部类的实现中，内部用的数组就是传入的数组，没有拷贝，也不会动态改变大小，所以对数组的修改也会反映到List中，对List调用add、remove方法会抛出异常。使用ArrayList完整的方法，应该新建一个ArrayList.

ArrayList还提供了两个public方法，可以控制内部使用的数组大小，一个是**ensureCapacity**,可以确保数组的大小至少为minCapacity，如果不够，会进行扩展。如果已经预知ArrayList需要比较大的容量，调用这个方法可以减少ArrayList内部分配和扩展的次数。另一个是**trimToSize**,它会重新分配一个数组，大小刚好为实际内容的长度。调用这个方法可以节省数组占用的空间。


## 9.2 剖析LinkedList

### 9.2.1 用法

  - 队列就类似于日常生活中的各种排队，特点就是**先进先出**，在尾部添加元素，从头部删除元素.

Queue扩展了Collection，它的主要操作有三个：

❑ 在尾部添加元素（add、offer）；

❑ 查看头部元素（element、peek），返回头部元素，但不改变队列；

❑ 删除头部元素（remove、poll），返回头部元素，并且从队列中删除。

每种操作都有两种形式，有什么区别呢？区别在于，对于特殊情况的处理不同。特殊情况是指队列为空或者队列为满，为空容易理解，为满是指队列**有长度大小限制，而且已经占满**了。LinkedList的实现中，队列长度没有限制，但别的Queue的实现可能有。在队列为空时，**element和remove**会抛出异常NoSuchElementException，而**peek和poll**返回特殊值null；在队列为满时，**add**会抛出异常IllegalStateException，而**offer**只是返回false。
 
    - 会抛异常的操作：element remove add
    - 返回null的操作：peek poll offer

  - 栈也是一种常用的数据结构，与队列相反，它的特点是**先进后出、后进先出**，类似于一个**储物箱**，放的时候是一件件往上放，拿的时候则只能从上面开始拿。

1）push表示入栈，在头部添加元素，栈的空间可能是有限的，如果栈满了，push会抛出异常IllegalStateException。

2）pop表示出栈，返回头部元素，并且从栈中删除，如果栈为空，会抛出异常NoSuch-ElementException。

3）peek查看栈头部元素，不修改栈，如果栈为空，返回null。

栈和队列都是在两端进行操作，**栈只操作头部，队列两端都操作，但尾部只添加、头部只查看和删除**

## 9.3 剖析ArrayDeque

### 9.3.1 实现原理

ArrayDeque的高效来源于head和tail这两个变量，它们使得物理上简单的从头到尾的数组变为了一个逻辑上循环的数组，避免了在头尾操作时的移动。

**构造方法** 

````
        public ArrayDeque(int numElements) {
            allocateElements(numElements);
        }
````
它主要就是在计算应该分配的数组的长度，计算逻辑如下：

1）如果numElements小于8，就是8。

2）在numElements大于等于8的情况下，分配的实际长度是严格大于numElements并且为2的整数次幂的最小数。比如，如果numElements为10，则实际分配16，如果num-Elements为32，则为64。

为什么要为2的幂次数呢？我们待会会看到，这样会使得很多操作的效率很高。为什么要严格大于numElements呢？因为循环数组必须时刻至少留一个空位，tail变量指向下一个空位，为了容纳numElements个元素，至少需要numElements+1个位置。

**ArrayDeque的基本原理**，内部它是一个动态扩展的循环数组，通过head和tail变量维护数组的开始和结尾，数组长度为2的幂次方，使用高效的位操作进行各种判断，以及对head和tail进行维护。

### 9.3.2 ArrayDeque特点分析

ArrayDeque实现了双端队列，内部使用循环数组实现，这决定了它有如下特点。

1）在两端添加、删除元素的效率很高，动态扩展需要的内存分配以及数组复制开销可以被平摊，具体来说，添加N个元素的效率为O(N)。

2）根据元素内容查找和删除的效率比较低，为O(N)。

3）与ArrayList和LinkedList不同，没有索引位置的概念，不能根据索引位置进行操作。

**ArrayDeque和LinkedList都实现了Deque接口，应该用哪一个呢？**

如果只需要Deque接口，从两端进行操作，一般而言，ArrayDeque效率更高一些，应该被优先使用；如果同时需要根据索引位置进行操作，或者经常需要在中间进行插入和删除，则应该选LinkedList。

# 10.Map和Set

## 10.1 剖析HashMap

### 10.1.1 Map接口

````
        public interface Map<K, V> { //K和V是类型参数，分别表示键(Key)和值(Value)的类型
            V put(K key, V value); //保存键值对，如果原来有key，覆盖，返回原来的值
            V get(Object key); //根据键获取值， 没找到，返回null
            V remove(Object key); //根据键删除键值对， 返回key原来的值，如果不存在，返回null
            int size(); //查看Map中键值对的个数
            boolean isEmpty(); //是否为空
            boolean containsKey(Object key); //查看是否包含某个键
            boolean containsValue(Object value); //查看是否包含某个值
            void putAll(Map<? extends K, ? extends V> m); //保存m中的所有键值对到当前Map
            void clear(); //清空Map中所有键值对
            Set<K> keySet(); //获取Map中键的集合
            Collection<V> values(); //获取Map中所有值的集合
            Set<Map.Entry<K, V>> entrySet(); //获取Map中的所有键值对
            interface Entry<K, V> { //嵌套接口，表示一条键值对
                K getKey(); //键值对的键
                V getValue(); //键值对的值
                V setValue(V value);
                boolean equals(Object o);
                int hashCode();
            }
            boolean equals(Object o);
            int hashCode();
        }
        boolean containsValue(Object value);
        Set<K> keySet();
````

key和value分别表示键和值，next指向下一个Entry节点，hash是key的hash值，直接存储hash值是为了在比较的时候加快计算

根据哈希值存取对象、比较对象是计算机程序中一种重要的思维方式，它使得存取对象主要依赖于自身Hash值，而不是与其他对象进行比较，存取效率也与集合大小无关，高达O(1)，即使进行比较，也利用Hash值提高比较性能。

## 10.2 剖析HashSet

HashSet内部是用HashMap实现的，它内部有一个HashMap实例变量

## 10.4 剖析TreeMap

TreeMap基于大致平衡的排序二叉树：红黑树，这决定了它有如下特点。

1）没有重复元素。

2）添加、删除元素、判断元素是否存在，效率比较高，为O(log2(N)), N为元素个数。

3）有序，TreeSet同样实现了SortedSet和NavigatableSet接口，可以方便地根据顺序进行查找和操作，如第一个、最后一个、某一取值范围、某一值的邻近元素等。

4）为了有序，TreeSet要求元素实现Comparable接口或通过构造方法提供一个Com-parator对象

## 10.7 剖析EnumMap

主要是因为枚举类型有两个特征：一是它可能的值是有限的且预先定义的；二是枚举值都有一个顺序，这两个特征使得可以更为高效地实现Map接口

## 10.8 剖析EnumSet

之前介绍的Set接口的实现类HashSet/TreeSet，它们内部都是用对应的HashMap/TreeMap实现的，但EnumSet不是，它的实现与EnumMap没有任何关系，而是用极为精简和高效的位向量实现的

# 第11章 堆与优先级队列

堆指的是内存中的区域，保存动态分配的对象，与栈相对应。

## 11.1 堆的概念与算法

堆首先是一棵二叉树，但它是完全二叉树。什么是完全二叉树呢？我们先来看另一个相似的概念：满二叉树。满二叉树是指除了最后一层外，每个节点都有两个孩子，而最后一层都是叶子节点，都没有孩子

![满二叉树](https://gitee.com/liujunrull/image-blob/raw/master/202209191057628.png)

满二叉树一定是完全二叉树，但完全二叉树不要求最后一层是满的，但如果不满，则要求所有节点必须集中在最左边，从左到右是连续的，中间不能有空的

![完全二叉树](https://gitee.com/liujunrull/image-blob/raw/master/202209191059641.png)

完全二叉树有一个重要的特点：给定任意一个节点，可以根据其编号直接快速计算出其父节点和孩子节点编号。如果编号为i，则父节点编号即为i/2，左孩子编号即为2× i，右孩子编号即为2× i+1。比如，对于5号节点，父节点为5/2即2，左孩子为2× 5即10，右孩子为2× 5+1即11。

![完全二叉树编号](https://gitee.com/liujunrull/image-blob/raw/master/202209191101830.png)

**从中间删除元素**

如果需要从中间删除某个节点，与从头部删除一样，都是先用最后一个元素替换待删元素。不过替换后，有两种情况：如果该元素大于某孩子节点，则需向下调整（sift-down）；如果小于父节点，则需向上调整（siftup）。我们来看个例子，删除值为21的节点

![第一步](https://gitee.com/liujunrull/image-blob/raw/master/202209191103297.png)

替换后，6没有子节点，小于父节点12，执行向上调整（siftup）过程，最后结果如图11-17所示。

![中间删除21](https://gitee.com/liujunrull/image-blob/raw/master/202209191104968.png)

将普通无序数组变为堆的过程称为heapify。基本思路是：从最后一个非叶子节点开始，一直往前直到根，对每个节点，执行向下调整（siftdown）。换句话说，是自底向上，先使每个最小子树为堆，然后每对左右子树和其父节点合并，调整为更大的堆，因为每个子树已经为堆，所以调整就是对父节点执行向下调整（siftdown），这样一直合并调整直到根。

堆是一种比较神奇的数据结构，概念上是树，存储为数组，父子有特殊顺序，根是最大值/最小值，构建/添加/删除效率都很高，可以高效解决很多问题。

## 11.2 剖析PriorityQueue

PriorityQueue是优先级队列，它首先实现了**队列接口（Queue）**，与LinkedList类似，它的队列长度也没有限制，与一般队列的区别是，它有**优先级**的概念，每个元素都有优先级，队头的元素永远都是优先级最高的。

优先级出队，内部是用**堆**实现的，有如下特点：

1）实现了优先级队列，最先出队的总是优先级**最高**的，即排序中的第一个。

2）优先级可以有相同的，内部元素不是完全有序的，如果遍历输出，除了第一个，其他没有特定顺序。

3）**查看头部元素的效率很高**，为O(1)，入队、出队效率比较高，为O(log2(N))，构建堆heapify的效率为O(N)。

4）根据值查找和删除元素的效率比较低，为O(N)。

# 12 通用容器类和总结

## 12.2 Collections

交换list中第i个和第j个元素的内容。实现代码为：

```
        public static void swap(List<? > list, int i, int j) {
            final List l = list;
            l.set(i, l.set(j, l.get(i)));
        }
```

如果需要实现类似“剪切”和“粘贴”的功能，可以使用rotate()方法。

批量填充固定值

````
        public static <T> void fill(List<? super T> list, T obj)
````

批量复制

````
        public static <T> void copy(List<? super T> dest, List<? extends T> src)
````

将列表src中的每个元素复制到列表dest的对应位置处，覆盖dest中原来的值，dest的列表长度不能小于src, dest中超过src长度部分的元素不受影响。

emptyList方法返回的是一个静态不可变对象，它可以节省创建新对象的内存和时间开销。

````
        public static final <T> List<T> emptyList() {
            return (List<T>) EMPTY_LIST;
        }
````

如果返回值只是用于读取，可以使用emptyList方法，但如果返回值还用于写入，则需要新建一个对象。

在Java 9中，可以使用List、Map和Set不带参数的of方法返回一个空的只读容器对象，也就是说，如下两行代码的效果是相同的：

````
        1. List list = Collections.emptyList();
        2. List list = List.of();
````

相比新建容器对象并添加元素，这些方法更为简洁方便，此外，它们的实现更为高效，它们的实现类都针对单一对象进行了优化。比如， singleton方法的代码：

````
        public static <T> Set<T> singleton(T o) {
            return new SingletonSet<>(o);
        }
````

## 12.3 容器类总结

Queue是Collection的子接口，表示先进先出的队列，在尾部添加，从头部查看或删除。Deque是Queue的子接口，表示更为通用的双端队列，有明确的在头或尾进行查看、添加和删除的方法。普通队列有两个主要的实现类：LinkedList和ArrayDeque。LinkedList基于链表实现，ArrayDeque基于循环数组实现。一般而言，如果只需要Deque接口，Array-Deque的效率更高一些。

#### 12.3.2 数据结构和算法

在容器类中，我们看到了如下数据结构的应用：

1）动态数组：**ArrayList**内部就是动态数组，**HashMap**内部的链表数组也是动态扩展的，**ArrayDeque**和**PriorityQueue**内部也都是动态扩展的数组。

2）链表：**LinkedList**是用双向链表实现的，**HashMap**中映射到同一个链表数组的键值对是通过单向链表链接起来的，**LinkedHashMap**中每个元素还加入到了一个双向链表中以维护插入或访问顺序。

3）哈希表：**HashMap**是用哈希表实现的，HashSet、LinkedHashSet和LinkedHashMap基于HashMap，内部当然也是哈希表。

4）排序二叉树：**TreeMap**是用红黑树（基于排序二叉树）实现的，**TreeSet**内部使用TreeMap，当然也是红黑树，红黑树能保持元素的顺序且综合性能很高。

5）堆：**PriorityQueue**是用堆实现的，堆逻辑上是树，物理上是动态数组，堆可以高效地解决一些其他数据结构难以解决的问题。

6）循环数组：**ArrayDeque**是用循环数组实现的，通过对头尾变量的维护，实现了高效的队列操作。

7）位向量：**EnumSet和BitSet**是用位向量实现的，对于只有两种状态，且需要进行集合运算的数据，使用位向量进行表示、位运算进行处理，精简且高效。

每种数据结构中往往包含一定的算法策略，这种策略往往是一种折中，比如：

1）动态扩展算法：动态数组的扩展策略，一般是指数级扩展的，是在两方面进行平衡，一方面是希望减少内存消耗，另一方面希望减少内存分配、移动和复制的开销。

2）哈希算法：哈希表中键映射到链表数组索引的算法，算法要快，同时要尽量随机和均匀。

3）排序二叉树的平衡算法：排序二叉树的平衡非常重要，红黑树是一种平衡算法， AVL树是另一种平衡算法。平衡算法一方面要保证尽量平衡，另一方面要尽量减少综合开销。

# 第13章 文件基本技术

硬盘的访问延时，相比内存，是很慢的。操作系统和硬盘一般是按**块**批量传输，而不是按字节，以摊销延时开销，块大小一般至少为512字节，即使应用程序只需要文件的一个字节，操作系统也会至少将一个块读进来。一般而言，应尽量减少接触硬盘，接触一次，就一次多做一些事情。对于网络请求和其他输入输出设备，原则都是类似的。

一般读写文件需要两次数据复制，比如读文件，需要先从硬盘复制到操作系统内核，再从内核复制到应用程序分配的内存中。操作系统运行所在的环境和应用程序是不一样的，操作系统所在的环境是内核态，应用程序是用户态，应用程序调用操作系统的功能，需要两次环境的切换，先从用户态切到内核态，再从内核态切到用户态。这种用户态/内核态的切换是有开销的，应尽量减少这种切换。

为了提升文件操作的效率，应用程序经常使用一种常见的策略，即使用**缓冲区**。读文件时，即使目前只需要少量内容，但预知还会接着读取，就一次读取比较多的内容，放到读缓冲区，下次读取时，如果缓冲区有，就直接从缓冲区读，减少访问操作系统和硬盘。写文件时，先写到写缓冲区，写缓冲区满了之后，再一次性调用操作系统写到硬盘。不过，需要注意的是，在写结束的时候，要记住将缓冲区的**剩余内容同步到硬盘**。操作系统自身也会使用缓冲区，不过，应用程序更了解读写模式，恰当使用往往可以有更高的效率。

操作系统操作文件一般有**打开和关闭**的概念。打开文件会在操作系统内核建立一个有关该文件的内存结构，这个结构一般通过一个整数索引来引用，这个索引一般称为**文件描述符**。这个结构是消耗内存的，操作系统能同时打开的文件一般也是有限的，在不用文件的时候，应该记住关闭文件。关闭文件一般会同步缓冲区内容到硬盘，并释放占据的内存结构。

操作系统一般支持一种称为**内存映射文件**的高效的随机读写大文件的方法，将文件直接映射到内存，操作内存就是操作文件。在内存映射文件中，只有访问到的数据才会被实际复制到内存，且数据只会复制一次，被操作系统以及多个应用程序共享。

基本的流按字节读写，没有缓冲区，这不方便使用。Java解决这个问题的方法是使用**装饰器**设计模式，引入了很多装饰类，对基本的流增加功能，以方便使用。一般一个类只关注一个方面，实际使用时，经常会需要多个装饰类。

# 14 常见文件类型处理

压缩一个文件或一个目录

````
        public static void zip(File inFile, File zipFile) throws IOException {
            ZipOutputStream out = new ZipOutputStream(new BufferedOutputStream(
                    new FileOutputStream(zipFile)));
            try {
                if(! inFile.exists()) {
                    throw new FileNotFoundException(inFile.getAbsolutePath());
                }
                inFile = inFile.getCanonicalFile();
                String rootPath = inFile.getParent();
                if(! rootPath.endsWith(File.separator)) {
                    rootPath += File.separator;
                }
                addFileToZipOut(inFile, out, rootPath);
            } finally {
                out.close();
            }
        }

                private static void addFileToZipOut(File file, ZipOutputStream out,
                String rootPath) throws IOException {
            String relativePath = file.getCanonicalPath().substring(
                        rootPath.length());
                if(file.isFile()) {
                    out.putNextEntry(new ZipEntry(relativePath));
                    InputStream in = new BufferedInputStream(new FileInputStream(file));
                    try {
                        copy(in, out);
                    } finally {
                        in.close();
                    }
                } else {
                    out.putNextEntry(new ZipEntry(relativePath + File.separator));
                    for(File f : file.listFiles()) {
                        addFileToZipOut(f, out, rootPath);
                    }
                }
            }
````

使用ZipInputStream解压文件，可以使用类似如下代码：

````
        public static void unzip(File zipFile, String destDir) throws IOException {
            ZipInputStream zin = new ZipInputStream(new BufferedInputStream(
                    new FileInputStream(zipFile)));
            if(! destDir.endsWith(File.separator)) {
                destDir += File.separator;
            }
            try {
                ZipEntry entry = zin.getNextEntry();
                while(entry ! = null) {
                    extractZipEntry(entry, zin, destDir);
                    entry = zin.getNextEntry();
                }
            } finally {
                zin.close();
            }
        }
````

调用extractZipEntry处理每个压缩条目，代码为：

````
        private static void extractZipEntry(ZipEntry entry, ZipInputStream zin,
                String destDir) throws IOException {
            if(! entry.isDirectory()) {
                File parent = new File(destDir + entry.getName()).getParentFile();
                if(! parent.exists()) {
                    parent.mkdirs();
                }
                  OutputStream entryOut = new BufferedOutputStream(
                          new FileOutputStream(destDir + entry.getName()));
                  try {
                      copy(zin, entryOut);
                  } finally {
                      entryOut.close();
                  }
              } else {
                  new File(destDir + entry.getName()).mkdirs();
              }
          }
````

## 14.5 使用Jackson序列化为JSON/XML/MessagePack

#### 引用同一个对象

我们看个简单的例子，有两个类Common和A, A中有两个Common对象，为便于演示，我们将所有属性定义为了public，它们的类定义如下：

````
        static class Common {
            public String name;
        }
        static class A {
            public Common first;
            public Common second;
        }
````

有一个A对象，如下所示：

````
        Common c = new Common();
        c.name= "common";
        A a = new A();
        a.first = a.second = c;
````

a对象的first和second都指向都一个c对象，不加额外配置，序列化a的代码为：

````
        ObjectMapper mapper = new ObjectMapper();
        mapper.enable(SerializationFeature.INDENT_OUTPUT);
        String str = mapper.writeValueAsString(a);
        System.out.println(str);
````

输出为：

````
        {
          "first" : {
            "name" : "abc"
          },
          "second" : {
            "name" : "abc"
          }
        }
````

在反序列化后，first和second将指向不同的对象，如下所示：

````
        A a2 = mapper.readValue(str, A.class);
        if(a2.first == a2.second){
            System.out.println("reference same object");
        }else{
            System.out.println("reference different objects");
        }
````

输出为：

````
        reference different objects
````

那怎样才能保持这种对同一个对象的引用关系呢？可以使用注解@JsonIdentityInfo，对Common类做注解，如下所示：

````
        @JsonIdentityInfo(
                generator = ObjectIdGenerators.IntSequenceGenerator.class,
                property="id")
        static class Common {
              public String name;
          }
````
@JsonIdentityInfo中指定了两个属性，property="id"表示在序列化输出中新增一个属性"id"以表示对象的唯一标示，generator表示对象唯一ID的产生方法，这里是使用整数顺序数产生器IntSequenceGenerator。加了这个标记后，序列化输出会变为：

````
        {
          "first" : {
            "id" : 1,
            "name" : "common"
          },
          "second" : 1
        }
````

# 15. 并发

## 15.1 线程

#### 为什么调用的是start，执行的却是run方法呢？

start表示启动该线程，使其成为一条单独的执行流，操作系统会分配线程相关的资源，每个线程会有单独的程序执行计数器和栈，操作系统会把这个线程作为一个独立的个体进行调度，分配时间片让它执行，执行的起点就是run方法。

关于线程，我们需要知道，它是有成本的。创建线程需要消耗操作系统的资源，操作系统会为每个线程创建必要的数据结构、栈、程序计数器等，创建也需要一定的时间。此外，线程调度和切换也是有成本的，当有大量可运行线程的时候，操作系统会忙于调度，为一个线程分配一段时间，执行完后，再让另一个线程执行，一个线程被切换出去后，操作系统需要保存它的当前上下文状态到内存，上下文状态包括当前CPU寄存器的值、程序计数器的值等，而一个线程被切换回来后，操作系统需要恢复它原来的上下文状态，整个过程称为上**下文切换**，这个切换不仅耗时，而且使CPU中的很多缓存失效。

synchronized实例方法实际保护的是同一个对象的方法调用，确保同时只能有一个线程执行。

synchronized保护的是对象而非代码，只要访问的是同一个对象的synchronized方法，即使是不同的代码，也会被同步顺序访问

synchronized保护的是对象，对实例方法，保护的是当前实例对象this，对静态方法，保护的是哪个对象呢？是类对象，这里是StaticCounter.class。实际上，每个对象都有一个锁和一个等待队列，类对象也不例外。

synchronized有一个重要的特征，它是**可重入**的，也就是说，对同一个执行线程，它在获得了锁之后，在调用其他需要同样锁的代码时，可以直接调用。比如，在一个syn-chronized实例方法内，可以直接调用其他synchronized实例方法。

#### 如何正确地取消/关闭线程

````
    //Future接口提供了如下方法以取消任务：
    boolean cancel(boolean mayInterruptIfRunning);
    //ExecutorService提供了如下两个关闭方法
    void shutdown();
    List<Runnable> shutdownNow();
````

# 16 并发包的基石

````
        public final boolean compareAndSet(int expect, int update)
````

compareAndSet是一个非常重要的方法，比较并设置，我们以后将简称为CAS。该方法有两个参数expect和update，以原子方式实现了如下功能：如果当前值等于expect，则更新为update，否则不更新，如果更新成功，返回true，否则返回false。

#### 可重入锁ReentrantLock

Lock接口的主要实现类是ReentrantLock，它的基本用法lock/unlock实现了与syn-chronized一样的语义，包括：

❑ 可重入，一个线程在持有一个锁的前提下，可以继续获得该锁；

❑ 可以解决竞态条件问题；

❑ 可以保证内存可见性。

ReentrantLock有两个构造方法：

        public ReentrantLock()
        public ReentrantLock(boolean fair)

参数fair表示是否保证公平，不指定的情况下，默认为false，表示不保证公平。所谓公平是指，等待时间最长的线程优先获得锁。保证公平会影响性能，一般也不需要，所以默认不保证，synchronized锁也是不保证公平的

使用显式锁，一定要记得调用unlock。一般而言，应该将lock之后的代码包装到try语句内，在finally语句内释放锁。比如，使用ReentrantLock实现Counter，代码可以为：

        public class Counter {
            private final Lock lock = new ReentrantLock();
            private volatile int count;
            public void incr() {
                lock.lock();
                try {
                    count++;
                } finally {
                    lock.unlock();
                }
            }
            public int getCount() {
                return count;
            }
        }

使用tryLock()，可以避免死锁。在持有一个锁获取另一个锁而获取不到的时候，可以释放已持有的锁，给其他线程获取锁的机会，然后重试获取所有锁。

# 18. 异步任务执行服务

- ExecutorService有两个关闭方法：shutdown和shutdownNow。区别是，shutdown表示不再接受新任务，但已提交的任务会继续执行，即使任务还未开始执行；shutdownNow不仅不接受新任务，而且会终止已提交但尚未执行的任务，对于正在执行的任务，一般会调用线程的interrupt方法尝试中断，不过，线程可能不响应中断，shutdownNow会返回已提交但尚未执行的任务列表。

- ExecutorService有两组批量提交任务的方法：invokeAll和invokeAny，它们都有两个版本，其中一个限定等待时间。invokeAll等待所有任务完成，返回的Future列表中，每个Future的isDone方法都返回true，不过isDone为true不代表任务就执行成功了，可能是被取消了。invokeAll可以指定等待时间，如果超时后有的任务没完成，就会被取消。而对于invokeAny，**只要有一个任务**在限时内成功返回了，它就会返回该任务的结果，其他任务会被取消；如果没有任务能在限时内成功返回，抛出TimeoutException；如果限时内所有任务都结束了，但都发生了异常，抛出ExecutionException。

- 线程池主要由两个概念组成：一个是任务队列；另一个是工作者线程。工作者线程主体就是一个循环，循环从队列中接受任务并执行，任务队列保存待执行的任务。

线程池的大小主要与4个参数有关：

❑ corePoolSize：核心线程个数。

❑ maximumPoolSize：最大线程个数。

❑ keepAliveTime和unit：空闲线程存活时间。

maximumPoolSize表示线程池中的最多线程数，线程的个数会动态变化，但这是最大值，不管有多少任务，都不会创建比这个值大的线程个数。corePoolSize表示线程池中的核心线程个数，不过，并不是一开始就创建这么多线程，刚创建一个线程池后，实际上并不会创建任何线程。

- 一般情况下，有新任务到来的时候，如果当前线程个数小于**corePoolSiz**，就会创建一个新线程来执行该任务，需要说明的是，即使其他线程现在也是空闲的，也会创建新线程。不过，如果线程个数大于等于corePoolSiz，那就不会立即创建新线程了，它会**先尝试排队**，需要强调的是，它是“尝试”排队，而不是“阻塞等待”入队，如果队列满了或其他原因不能立即入队，它就不会排队，而是检查线程个数是否达到了maximumPoolSize，如果没有，就会继续创建线程，直到线程数达到maximumPoolSize。

- ❑ LinkedBlockingQueue：基于链表的阻塞队列，可以指定最大长度，但默认是无界的。

    ❑ ArrayBlockingQueue：基于数组的有界阻塞队列。

    ❑ PriorityBlockingQueue：基于堆的无界阻塞优先级队列。

    ❑ SynchronousQueue：没有实际存储空间的同步阻塞队列。

如果用的是无界队列，需要强调的是，线程个数最多只能达到corePoolSize，到达core-PoolSize后，新的任务总会排队，参数maximumPoolSize也就没有意义了。对于SynchronousQueue，我们知道，它没有实际存储元素的空间，当尝试排队时，只有正好有空闲线程在等待接受任务时，才会入队成功，否则，总是会创建新线程，直到达到maximumPoolSize。

- 如果队列有界，且maximumPoolSize有限，则当队列排满，线程个数也达到了maxi-mumPoolSize，这时，新任务来了，如何处理呢？此时，会触发线程池的任务**拒绝策略**。

    拒绝策略是可以自定义的，ThreadPoolExecutor实现了4种处理方式。

    1）ThreadPoolExecutor.AbortPolicy：这就是默认的方式，抛出异常。

    2）ThreadPoolExecutor.DiscardPolicy：静默处理，忽略新任务，不抛出异常，也不执行。

    3）ThreadPoolExecutor.DiscardOldestPolicy：将等待时间最长的任务扔掉，然后自己排队。

    4）ThreadPoolExecutor.CallerRunsPolicy：在任务提交者线程中执行任务，而不是交给线程池中的线程执行。它们都是ThreadPoolExecutor的public静态内部类，都实现了RejectedExecutionHandler接口

    拒绝策略只有在队列有界，且maximumPoolSize有限的情况下才会触发。如果队列无界，服务不了的任务总是会排队，但这不一定是期望的结果，因为请求处理队列可能会消耗非常大的内存，甚至引发内存不够的异常。如果队列有界但maxi-mumPoolSize无限，可能会创建过多的线程，占满CPU和内存，使得任何任务都难以完成。所以，在任务量非常大的场景中，让拒绝策略有机会执行是保证系统稳定运行很重要的方面。

- 线程个数小于等于corePoolSize时，我们称这些线程为核心线程，默认情况下。

    ❑ 核心线程不会预先创建，只有当有任务时才会创建。

    ❑ 核心线程不会因为空闲而被终止，keepAliveTime参数不适用于它。

 - 只使用一个线程，使用无界队列LinkedBlockingQueue，线程创建后不会超时终止，该线程顺序执行所有任务。该线程池适用于需要确保所有任务被顺序执行的场合。

- 实际中，应该使用newFixedThreadPool还是newCachedThreadPool呢？在系统负载很高的情况下，newFixedThreadPool可以通过队列对新任务排队，保证有足够的资源处理实际的任务，而newCachedThreadPool会为每个任务创建一个线程，导致创建过多的线程竞争CPU和内存资源，使得任何实际任务都难以完成，这时， newFixedThreadPool更为适用。不过，如果系统负载不太高，单个任务的执行时间也比较短，newCachedThreadPool的效率可能更高，因为任务可以不经排队，直接交给某一个空闲线程。在系统负载可能极高的情况下，两者都不是好的选择，newFixedThreadPool的问题是队列过长，而newCachedThreadPool的问题是线程过多，这时，应根据具体情况自定义ThreadPoolExecutor，传递合适的参数。

#### 线程池的死锁

可以使用newCachedThreadPool创建线程池，让线程数不受限制。另一个解决方法是使用**SynchronousQueue**，它可以避免死锁，怎么做到的呢？对于普通队列，入队只是把任务放到了队列中，而对于SynchronousQueue来说，入队成功就意味着已有线程接受处理，如果入队失败，可以创建更多线程直到maximumPoolSize，如果达到了maximumPoolSize，会触发拒绝机制，不管怎么样，都不会死锁。

ThreadPoolExecutor实现了生产者/消费者模式，工作者线程就是消费者，任务提交者就是生产者，线程池自己维护任务队列。当我们碰到类似生产者/消费者问题时，应该优先考虑直接使用线程池，

# 第19章 同步和协作工具类

## 19.1 读写锁ReentrantReadWriteLock

多个线程的读操作完全可以并行，在读多写少的场景中，让读操作并行可以明显提高性能。怎么让读操作能够并行，又不影响一致性呢？答案是使用读写锁。在Java并发包中，接口ReadWriteLock表示读写锁，主要实现类是可重入读写锁**ReentrantReadWriteLock**。

通过一个ReadWriteLock产生两个锁：一个读锁，一个写锁。读操作使用读锁，写操作使用写锁。需要注意的是，只有“**读-读**”操作是可以并行的，“读-写”和“写-写”都不可以。**只有一个**线程可以进行写操作，在获取写锁时，只有没有任何线程持有任何锁才可以获取到，在持有写锁时，其他任何线程都获取不到任何锁。在没有其他线程持有写锁的情况下，多个线程可以获取和持有读锁。

## 19.2 信号量Semaphore

现实中，资源往往有多个，但每个同时只能被一个线程访问，比如，饭店的饭桌、火车上的卫生间。有的单个资源即使可以被并发访问，但并发访问数多了可能影响性能，所以希望限制并发访问的线程数。还有的情况，与软件的授权和计费有关，对不同等级的账户，限制不同的最大并发访问数。信号量类**Semaphore**就是用来解决这类问题的，它可以**限制对资源的并发访问数**

## 19.3 倒计时门栓CountDownLatch

CountDownLatch。它相当于是一个门栓，一开始是关闭的，所有希望通过该门的线程都需要等待，然后开始倒计时，倒计时变为0后，门栓打开，等待的**所有线程**都可以通过，它是一次性的，打开后就不能再关上了

## 19.4 循环栅栏CyclicBarrier

CyclicBarrier。它相当于是一个栅栏，所有线程在到达该栅栏后都需要等待其他线程，等**所有线程都到达后再一起通过**，它是循环的，可以用作重复的同步。CyclicBarrier特别适用于**并行迭代计算**，每个线程负责一部分计算，然后在栅栏处等待其他线程完成，所有线程到齐后，交换数据和计算结果，再进行下一次迭代。
CyclicBarrier与CountDownLatch可能容易混淆，我们强调下它们的区别。

1）CountDownLatch的参与线程是有不同角色的，有的负责倒计时，有的在等待倒计时变为0，负责倒计时和等待倒计时的线程都可以有多个，用于不同角色线程间的同步。

2）CyclicBarrier的参与线程角色是一样的，用于同一角色线程间的协调一致。

3）CountDownLatch是一次性的，而CyclicBarrier是可以重复利用的。

## 19.5 理解ThreadLocal

线程本地变量是说，每个线程都有同一个变量的独有拷贝

initialValue用于提供初始值，这是一个受保护方法，可以通过匿名内部类的方式提供，当调用get方法时，如果之前没有设置过，会调用该方法获取初始值，默认实现是返回null。remove删掉当前线程对应的值，如果删掉后，再次调用get，会再调用initialValue获取初始值。

使用场景：日期处理、随机数和上下文信息。

ThreadLocal对象一般都定义为static，以便于引用。

每个线程都有一个Map，对于每个ThreadLocal对象，调用其get/set实际上就是以ThreadLocal对象为键读写当前线程的Map，这样，就实现了每个线程都有自己的独立副本的效果

1）ThreadLocal使得每个线程对同一个变量有自己的独立副本，是实现线程安全、减少竞争的一种方案。

2）ThreadLocal经常用于存储上下文信息，避免在不同代码间来回传递，简化代码。

3）每个线程都有一个Map，调用ThreadLocal对象的get/set实际就是以ThreadLocal对象为键读写当前线程的该Map。

# 第20章 并发总结

## 20.1 线程安全的机制

线程表示一条单独的执行流，每个线程有自己的执行计数器，有自己的栈，但可以共享内存，共享内存是实现线程协作的基础，但共享内存有两个问题，竞态条件和内存可见性，之前章节探讨了解决这些问题的多种思路：

❑ 使用synchronized；

❑ 使用显式锁；

❑ 使用volatile；

❑ 使用原子变量和CAS；

❑ 写时复制；

❑ 使用ThreadLocal。

（1）synchronized

synchronized简单易用，它只是一个关键字，大部分情况下，放到类的方法声明上就可以了，既可以解决竞态条件问题，也可以解决内存可见性问题。需要理解的是，它保护的是对象，而不是代码，只有对同一个对象的synchronized方法调用，synchronized才能保证它们被顺序调用。对于实例方法，这个对象是this；对于静态方法，这个对象是类对象；对于代码块，需要指定哪个对象。另外，需要注意，它不能尝试获取锁，也不响应中断，还可能会死锁。不过，相比显式锁，synchronized简单易用，JVM也可以不断优化它的实现，应该被优先使用。

（2）显式锁

显式锁是相对于synchronized隐式锁而言的，它可以实现synchronized同样的功能，但需要程序员自己创建锁，调用锁相关的接口，主要接口是Lock，主要实现类是Reen-trantLock。相比synchronized，显式锁支持以非阻塞方式获取锁，可以响应中断，可以限时，可以指定公平性，可以解决死锁问题，这使得它灵活得多。在读多写少、读操作可以完全并行的场景中，可以使用读写锁以提高并发度，读写锁的接口是ReadWriteLock，实现类是ReentrantReadWriteLock。

（3）volatile

synchronized和显式锁都是锁，使用锁可以实现安全，但使用锁是有成本的，获取不到锁的线程还需要等待，会有线程的上下文切换开销等。保证安全不一定需要锁。如果共享的对象只有一个，操作也只是进行最简单的get/set操作，set也不依赖于之前的值，那就不存在竞态条件问题，而只有内存可见性问题，这时，在变量的声明上加上volatile就可以了。（

4）原子变量和CAS

使用volatile, set的新值不能依赖于旧值，但很多时候，set的新值与原来的值有关，这时，也不一定需要锁，如果需要同步的代码比较简单，可以考虑原子变量，它们包含了一些以原子方式实现组合操作的方法，对于并发环境中的计数、产生序列号等需求，考虑使用原子变量而非锁。原子变量的基础是CAS，一般的计算机系统都在硬件层次上直接支持CAS指令。通过循环CAS的方式实现原子更新是一种重要的思维。相比synchronized，它是乐观的，而synchronized是悲观的；它是非阻塞式的，而synchronized是阻塞式的。CAS是Java并发包的基础，基于它可以实现高效的、乐观、非阻塞式数据结构和算法，它也是并发包中锁、同步工具和各种容器的基础。

（5）写时复制

之所以会有线程安全的问题，是因为多个线程并发读写同一个对象，如果每个线程读写的对象都是不同的，或者，如果共享访问的对象是只读的，不能修改，那也就不存在线程安全问题了。我们在介绍容器类CopyOnWriteArrayList和CopyOnWriteArraySet时介绍了写时复制技术，写时复制就是将共享访问的对象变为只读的，写的时候，再使用锁，保证只有一个线程写，写的线程不是直接修改原对象，而是新创建一个对象，对该对象修改完毕后，再原子性地修改共享访问的变量，让它指向新的对象。

（6）ThreadLocal

ThreadLocal就是让每个线程，对同一个变量，都有自己的独有副本，每个线程实际访问的对象都是自己的，自然也就不存在线程安全问题了。

## 20.2 线程的协作机制

（1）wait/notify

wait/notify与synchronized配合一起使用，是线程的基本协作机制。每个对象都有一把锁和两个等待队列，一个是锁等待队列，放的是等待获取锁的线程；另一个是条件等待队列，放的是等待条件的线程，wait将自己加入条件等待队列，notify从条件等待队列上移除一个线程并唤醒，notifyAll移除所有线程并唤醒。需要注意的是，wait/notify方法只能在synchronized代码块内被调用，调用wait时，线程会释放对象锁，被notify/notifyAll唤醒后，要重新竞争对象锁，获取到锁后才会从wait调用中返回，返回后，不代表其等待的条件就一定成立了，需要重新检查其等待的条件。wait/notify方法看上去很简单，但往往难以理解wait等的到底是什么，而notify通知的又是什么，只能有一个条件等待队列，这也是wait/notify机制的局限性，这使得对于等待条件的分析变得复杂

（2）显式条件

显式条件与显式锁配合使用，与wait/notify相比，可以支持多个条件队列，代码更为易读，效率更高。使用时注意不要将signal/signalAll误写为notify/notifyAll。

（3）线程的中断

Java中取消/关闭一个线程的方式是中断。中断并不是强迫终止一个线程，它是一种协作机制，是给线程传递一个取消信号，但是由线程来决定如何以及何时退出，线程在不同状态和IO操作时对中断有不同的反应。作为线程的实现者，应该提供明确的取消/关闭方法，并用文档清楚描述其行为；作为线程的调用者，应该使用其取消/关闭方法，而不是贸然调用interrupt。

（4）协作工具类

除了基本的显式锁和条件，针对常见的协作场景，Java并发包提供了多个用于协作的工具类。信号量类Semaphore用于限制对资源的并发访问数。倒计时门栓CountDownLatch主要用于不同角色线程间的同步，比如在裁判/运动员模式中，裁判线程让多个运动员线程同时开始，也可以用于协调主从线程，让主线程等待多个从线程的结果。倒计时门栓CountDownLatch主要用于不同角色线程间的同步，比如在裁判/运动员模式中，裁判线程让多个运动员线程同时开始，也可以用于协调主从线程，让主线程等待多个从线程的结果。循环栅栏CyclicBarrier用于同一角色线程间的协调一致，所有线程在到达栅栏后都需要等待其他线程，等所有线程都到达后再一起通过，它是循环的，可以用作重复的同步。

（5）阻塞队列

对于最常见的生产者/消费者协作模式，可以使用阻塞队列，阻塞队列封装了锁和条件，生产者线程和消费者线程只需要调用队列的入队/出队方法就可以了，不需要考虑同步和协作问题。阻塞队列有普通的先进先出队列，包括基于数组的ArrayBlockingQueue和基于链表的LinkedBlockingQueue/LinkedBlockingDeque，也有基于堆的优先级阻塞队列PriorityBlock-ingQueue，还有可用于定时任务的延时阻塞队列DelayQueue，以及用于特殊场景的阻塞队列SynchronousQueue和LinkedTransferQueue。

（6）Future/FutureTask

在常见的主从协作模式中，主线程往往是让子线程异步执行一项任务，获取其结果。手工创建子线程的写法往往比较麻烦，常见的模式是使用异步任务执行服务，不再手工创建线程，而只是提交任务，提交后马上得到一个结果，但这个结果不是最终结果，而是一个Future。Future是一个接口，主要实现类是FutureTask。Future封装了主线程和执行线程关于执行状态和结果的同步，对于主线程而言，它只需要通过Future就可以查询异步任务的状态、获取最终结果、取消任务等，不需要再考虑同步和协作问题。

# 第21章 反射

## 21.1 Class类

### 1．名称信息

Class有如下方法，可以获取与名称有关的信息

        public String getName()
        public String getSimpleName()
        public String getCanonicalName()
        public Package getPackage()

![](https://gitee.com/liujunrull/image-blob/raw/master/202210111116056.png)

### 2．字段信息

类中定义的静态和实例变量都被称为字段，用类Field表示，位于包java.lang.reflect下

        //返回所有的public字段，包括其父类的，如果没有字段，返回空数组
        public Field[] getFields()
        //返回本类声明的所有字段，包括非public的，但不包括父类的
        public Field[] getDeclaredFields()
        //返回本类或父类中指定名称的public字段，找不到抛出异常NoSuchFieldException
        public Field getField(String name)
        //返回本类中声明的指定名称的字段，找不到抛出异常NoSuchFieldException
        public Field getDeclaredField(String name)

        //获取字段的名称
        public String getName()
        //判断当前程序是否有该字段的访问权限
        public boolean isAccessible()
        //flag设为true表示忽略Java的访问检查机制，以允许读写非public的字段
        public void setAccessible(boolean flag)
        //获取指定对象obj中该字段的值
        public Object get(Object obj)
        //将指定对象obj中该字段的值设为value
        public void set(Object obj, Object value)

在get/set方法中，对于静态变量，obj被忽略，可以为null，如果字段值为基本类型， get/set会自动在基本类型与对应的包装类型间进行转换；对于private字段，直接调用get/set会抛出非法访问异常IllegalAccessException，应该先调用setAccessible(true)以关闭Java的检查机制

### 3．方法信息

        //返回所有的public方法，包括其父类的，如果没有方法，返回空数组
        public Method[] getMethods()
        //返回本类声明的所有方法，包括非public的，但不包括父类的
        public Method[] getDeclaredMethods()
        //返回本类或父类中指定名称和参数类型的public方法，
        //找不到抛出异常NoSuchMethodException
        public Method getMethod(String name, Class<? >... parameterTypes)
        //返回本类中声明的指定名称和参数类型的方法，找不到抛出异常NoSuchMethodException
        public Method getDeclaredMethod(String name, Class<? >... parameterTypes)
        //获取方法的名称
        public String getName()
        //flag设为true表示忽略Java的访问检查机制，以允许调用非public的方法
        public void setAccessible(boolean flag)
        //在指定对象obj上调用Method代表的方法，传递的参数列表为args
        public Object invoke(Object obj, Object... args) throws
            IllegalAccessException, Illegal-ArgumentException, InvocationTargetException

### 4．创建对象和构造方法

        public T newInstance() throws InstantiationException, IllegalAccessException

它会调用类的默认构造方法（即无参public构造方法），如果类没有该构造方法，会抛出异常InstantiationException。

newInstance只能使用默认构造方法。Class还有一些方法，可以获取所有的构造方法：

        //获取所有的public构造方法，返回值可能为长度为0的空数组
        public Constructor<? >[] getConstructors()
        //获取所有的构造方法，包括非public的
        public Constructor<? >[] getDeclaredConstructors()
        //获取指定参数类型的public构造方法，没找到抛出异常NoSuchMethodException
        public Constructor<T> getConstructor(Class<? >... parameterTypes)
        //获取指定参数类型的构造方法，包括非public的，没找到抛出异常NoSuchMethodException
        public Constructor<T> getDeclaredConstructor(Class<? >... parameterTypes)

类Constructor表示构造方法，通过它可以创建对象，方法为：

        public T newInstance(Object ... initargs) throws InstantiationException,
        IllegalAccessException, IllegalArgumentException, InvocationTargetException

### 5．类型检查和转换

我们之前介绍过instanceof关键字，它可以用来判断变量指向的实际对象类型。instanceof后面的类型是在代码中确定的，如果要检查的类型是动态的，可以使用Class类的如下方法：

        public native boolean isInstance(Object obj)

        Class cls = Class.forName("java.util.ArrayList");
        if(cls.isInstance(list)){
            System.out.println("array list");
        }

强制转换到的类型是在写代码时就知道的。如果是动态的，可以使用Class的如下方法：

        public T cast(Object obj)

        public static <T> T toType(Object obj, Class<T> cls){
        return cls.cast(obj);
    }

isInstance/cast描述的都是对象和类之间的关系，Class还有一个方法，可以判断Class之间的关系：

        //检查参数类型cls能否赋给当前Class类型的变量
        public native boolean isAssignableFrom(Class<? > cls);

### 6．类的加载

Class有两个静态方法，可以根据类名加载类：

        public static Class<? > forName(String className)
        public static Class<? > forName(String name, boolean initialize,
            ClassLoader loader)
ClassLoader表示类加载器，第24章会进一步介绍，initialize表示加载后，是否执行类的初始化代码（如static语句块）。第一个方法中没有传这些参数，相当于调用：

        Class.forName(className, true, currentLoader)
    
currentLoader表示加载当前类的ClassLoader。

基本类型不支持forName方法

反射虽然是灵活的，但一般情况下，并不是我们优先建议的，主要原因是：1）反射更容易出现运行时错误，使用显式的类和接口，编译器能帮我们做类型检查，减少错误，但使用反射，类型是运行时才知道的，编译器无能为力。2）反射的性能要低一些，在访问字段、调用方法前，反射先要查找对应的Field/Method，要慢一些。

# 第22章 注解

### 框架和库的注解

声明式编程风格，在这种风格中，程序都由三个组件组成：❑ 声明的关键字和语法本身。❑ 系统/框架/库，它们负责解释、执行声明式的语句。❑ 应用程序，使用声明式风格写程序。

### 创建注解

        @Target(ElementType.METHOD)
        @Retention(RetentionPolicy.SOURCE)
        public @interface Override {
        }

**@Target**

表示注解的目标，@Override的目标是方法（ElementType.METHOD）。ElementType是一个枚举，主要可选值有：

❑ TYPE：表示类、接口（包括注解），或者枚举声明；

❑ FIELD：字段，包括枚举常量；

❑ METHOD：方法；

❑ PARAMETER：方法中的参数；

❑ CONSTRUCTOR：构造方法；

❑ LOCAL_VARIABLE：本地变量；

❑ MODULE：模块（Java 9引入的）。

目标可以有**多个**，用{}表示，比如@SuppressWarnings的@Target就有多个。如果没有声明@Target，默认为适用于**所有类型**。

**@Retention**

表示注解信息保留到什么时候，取值只能有一个，类型为RetentionPolicy，它是一个枚举，有三个取值。

❑ SOURCE：只在**源代码**中保留，编译器将代码编译为字节码文件后就会丢掉。

❑ CLASS：保留到**字节码**文件中，但Java虚拟机将class文件加载到内存时不一定会在内存中保留。

❑ RUNTIME：一直保留到**运行**时。

如果没有声明@Retention，则默认为**CLASS**。

可以为注解定义一些参数，定义的方式是在注解内定义一些方法

注解内参数的类型不是什么都可以的，合法的类型有基本类型、String、Class、枚举、注解，以及这些类型的数组。

参数定义时可以使用default指定一个默认值

        @Target({ METHOD, CONSTRUCTOR, FIELD })
        @Retention(RUNTIME)
        @Documented
        public @interface Inject {
            boolean optional() default false;
        }

# 第23章 类加载机制

## 24.1 类加载的基本机制和过程

负责加载类的类就是类加载器，它的输入是完全限定的类名，输出是Class对象。类加载器不是只有一个，一般程序运行时，都会有三个（适用于Java 9之前，Java 9引入了模块化，基本概念是类似的，但有一些变化，限于篇幅，就不探讨了）。

1）启动类加载器（Bootstrap ClassLoader）：这个加载器是Java虚拟机实现的一部分，不是Java语言实现的，一般是C++实现的，它负责加载Java的基础类，主要是<JAVA_HOME>/lib/rt.jar，我们日常用的Java类库比如String、ArrayList等都位于该包内。

2）扩展类加载器（Extension ClassLoader）：这个加载器的实现类是sun.misc.Laun-cher$ExtClassLoader，它负责加载Java的一些扩展类，一般是<JAVA_HOME>/lib/ext目录中的jar包。

3）应用程序类加载器（Application ClassLoader）：这个加载器的实现类是sun.misc.Launcher$AppClassLoader，它负责加载应用程序的类，包括自己写的和引入的第三方法类库，即所有在类路径中指定的类。

Application ClassLoader的父亲是Extension ClassLoader, Extension的父亲是Bootstrap ClassLoader。注意不是父子继承关系，而是**父子委派**关系，子ClassLoader有一个变量**parent**指向父ClassLoader，在子Class-Loader加载类时，一般会首先通过父ClassLoader加载，具体来说，在加载一个类时，基本过程是：

1）判断是否已经加载过了，加载过了，直接返回Class对象，一个类只会被一个Class-Loader加载一次。

2）如果没有被加载，先让父ClassLoader去加载，如果加载成功，返回得到的Class对象。

3）在父ClassLoader没有加载成功的前提下，自己尝试加载类。

“双亲委派”模型，即优先让父ClassLoader去加载。为什么要先让父ClassLoader去加载呢？这样，可以避免**Java类库被覆盖**的问题。比如，用户程序也定义了一个类java.lang.String，通过双亲委派，java.lang.String只会被Bootstrap ClassLoader加载，避免自定义的String覆盖Java类库的定义。

“双亲委派”虽然是一般模型，但也有一些例外，比如：

1）**自定义的加载顺序**：尽管不被建议，自定义的ClassLoader可以不遵从“双亲委派”这个约定，不过，即使不遵从，以java开头的类也不能被自定义类加载器加载，这是由Java的安全机制保证的，以避免混乱。

2）**网状加载顺序**：在OSGI框架和Java 9模块化系统中，类加载器之间的关系是一个网，每个模块有一个类加载器，不同模块之间可能有依赖关系，在一个模块加载一个类时，可能是从自己模块加载，也可能是委派给其他模块的类加载器加载。

3）**父加载器委派给子加载器加载**：典型的例子有JNDI服务（Java Naming and DirectoryInterface），它是Java企业级应用中的一项服务，具体我们就不介绍了。

## 24.2 自定义ClassLoader

自定义Class-Loader是Tomcat实现应用隔离、支持JSP、OSGI实现动态模块化的基础。

继承类ClassLoader，重写findClass就可以了

不把BASE_DIR放到classpath中，而是使用MyClassLoader加载，还有一个很大的好处，那就是可以创建多个MyClassLoader，对同一个类，每个MyClassLoader都可以加载一次，得到同一个类的不同Class对象

1）可以实现隔离。一个复杂的程序，内部可能按模块组织，不同模块可能使用同一个类，但使用的是不同版本，如果使用同一个类加载器，它们是无法共存的，不同模块使用不同的类加载器就可以实现隔离，Tomcat使用它隔离不同的Web应用，OSGI使用它隔离不同模块。

2）可以实现热部署。使用同一个ClassLoader，类只会被加载一次，加载后，即使class文件已经变了，再次加载，得到的也还是原来的Class对象，而使用MyClassLoader，则可以先创建一个新的ClassLoader，再用它加载Class，得到的Class对象就是新的，从而实现动态更新

# 25.函数式编程

## lambda表达式

        File[] files = f.listFiles((File dir, String name) -> name.endsWith(".txt"));

没有括号的时候，主体代码是一个表达式，这个表达式的值就是函数的返回值，结尾不能加分号，也不能加return语句

        List<String> names = map(students, Student::getName);
Student::getName这种写法是Java 8引入的一种新语法，称为方法引用。它是Lambda表达式的一种简写方法，由：：分隔为两部分，前面是类名或变量名，后面是方法名。

## 函数式数据处理

实际上，调用filter()和map()都不会执行任何实际的操作，它们只是在构建操作的流水线，调用collect才会触发实际的遍历执行，在一次遍历中完成过滤、转换以及收集结果的任务。

像filter和map这种不实际触发执行、用于构建流水线、返回Stream的操作称为中间操作（intermediate operation），而像collect这种触发实际执行、返回具体结果的操作称为终端操作（terminal operation）

流定义了很多数据处理的基本函数，对于一个具体的数据处理问题，解决的主要思路就是组合利用这些基本函数，以声明式的方式简洁地实现期望的功能，这种思路就是函数式数据处理思维，相比直接利用容器类API的命令式思维，思考的层次更高。

Stream API的这种思路也不是新发明，它与数据库查询语言SQL是很像的，都是声明式地操作集合数据，很多函数都能在SQL中找到对应，

Stream API也与各种基于Unix系统的管道命令类似。熟悉Unix系统的都知道，Unix有很多命令，大部分命令只是专注于完成一件事情，但可以通过管道的方式将多个命令链接起来，完成一些复杂的功能




