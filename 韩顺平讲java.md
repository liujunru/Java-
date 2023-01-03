#### 数据类型

![数据类型](https://gitee.com/liujunrull/image-blob/raw/master/202210021027371.png)

- float单精度，double双精度。追求数据准确性采用双精度
- char类型存放汉字的话占用两个字节
- bit是计算机的最小存储单位，byte是计算机的基本存储单位

![整形数值范围](https://gitee.com/liujunrull/image-blob/raw/master/202210021032821.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202210021057119.png)

- byte short char 三者可以计算，在计算时首先转换为int类型

````java
byte b2 = 1;
short s1 = 1;
short s2 = b2 + s1;//错
int s3 = b2 + s1;//对
````
- 二进制转换为八进制

    - 从低位开始，将二进制每三位一组，转为对应的八进制数即可
    - 0b11(3)010(2)101(5) => 0325

- 二进制转换为十六进制

    - 从低位开始将二进制数每四位一组，转为对应的十六进制数即可
    - 0b1101(D)0101(5) => 0xD5


![](https://gitee.com/liujunrull/image-blob/raw/master/202210031735575.png)

- 用补码的方式来运算可以把正负数统一起来

![](https://gitee.com/liujunrull/image-blob/raw/master/202210031750937.png)

````java
//当a%b a是小数时，公式=a-(int)a/b*b
-10.5%3=-10.5-(-10)/3*3 = -10.5+9 = -1.5
````
![](https://gitee.com/liujunrull/image-blob/raw/master/202210032049604.png)

![](https://gitee.com/liujunrull/image-blob/raw/master/202210032229485.png)

- javap A.class:反编译指令
- 对象创建的流程分析：
 ````java
  class Person{
    int age = 90;
    String name;
    Person(String n,int a){
        name = n;
        age = a;
    }
  }
  Person p = nes Person("xxx",20);
 ````

  - 1.加载Person类信息（Person.class)，只会加载一次
  - 2.在堆中分配空间（地址）
  - 3.完成对象初始化
  - [3.1默认初始化age=0 name=null
  - 3.2 显式初始化age=90 name=null
  - 3.3构造器初始化age=20 name=xxx]
  - 4.将对象在堆中的地址，返回给p（p是对象名，也可以理解称对象的引用）

- 多态的注意事项：

    - 属性没有重写之说！属性的值看编译类型
    - 对象编译时看左边，运行时看右边。可以调用的方法看编译类型，为向上转型。调用子类特有的方法需要向下转型。

- Java的动态绑定机制
  1. 当调用对象方法时，该方法会和该对象的内存地址/运行类型绑定
  2. 当调用对象属性时，没有动态绑定机制，哪里声明，哪里使用

- 不需要创建对象的方法用static修饰
- 代码块的使用
  - 相当于另外一种形式的构造器（对构造器的补充机制），可以做初始化的操作
  - 如果多个构造器中都有重复的语句，可以抽取到初始化代码块中，提高代码的重用性
  - 执行先于构造器的初始化

- 代码块使用细节
  - 静态代码块作用是对类进行初始化，而且随着类的加载而执行，并且只会执行一次。普通代码块每创建一个对象，执行一次。
  - 类什么时候被加载：
    - 创建对象实例时
    - 创建子类对象实例，父类也会被加载
    - 使用类的静态成员时（静态属性，静态方法）
  - 普通代码块，在创建对象实例时，会被隐式的调用。被创建一次，就会调用一次。如果只是使用类的静态成员时，普通代码块并不会执行。

- 代码块使用细节：

  - 创建一个子类时，他们的静态代码块、静态属性初始化、普通代码块、普通属性初始化、构造方法的调用顺序如下：
    1. 父类的静态代码块和静态属性（优先级一样，按定义顺序执行）
    2. 子类的静态代码块和静态属性（优先级一样，按定义顺序执行）
    3. 父类的普通代码块和普通属性初始化（优先级一下，按定义顺序执行）
    4. 父类的构造方法
    5. 子类的普通代码块和普通属性初始化（优先级一样，按定义顺序执行）
    6. 子类的构造方法

- 静态代码块只能直接调用静态成员（静态属性和静态方法）。普通代码块可以调用任意成员

- 单例模式

  - 饿汉式

    ```java
    class SingleTon01{
    	//将构造器私有化，防止外部直接new
        private SingleTon01{}
        //提供一个静态属性类型就是SIngleTon01
        //先创建好这个SingleTon01对象实例
        private static SinleTon01 instance = new SingleTon01();
        
        //提供一个public静态方法，可以返回instance
        public static SingleTon01 getInstance(){
            return instance;
        }
    }
    ```

  - 懒汉式

    ```java
    class SingleTon02{
        //构造器私有化
        private SingleTon02(){}
        //提供一个静态属性类型是SingleTon02
        private static SingleTon02 instance;
        //提供一个public静态方法，可以返回instance
        public synchronized static SingleTon02 getInstance(){
            if(instance == null){
                instance = new SingleTon02();
            }
            return  instance;
        }
    }
    ```

    

- final 和static往往搭配使用，效率更高，不会导致类加载，底层编译器做了优化处理。

- 接口中的属性，只能是final的，而且是public satic final修饰符

  - 接口中的属性的访问形式：接口名.属性名
  - 接口不能继承其他的类，但可以继承多个别的接口

- 类的五大成员：属性、方法、构造器、代码块、内部类

- 内部类

  1. 局部内部类：

     - 可以直接访问外部类的所有成员，包括私有的

     - 不能添加访问修饰符，因为他的地位就是一个局部变量。局部变量是不能使用修饰符的。但是可以用final修饰，因为局部变量也可以使用final

     - 作用域：仅仅在定义他的方法或代码块中。

     - 局部内部类访问外部类的成员：直接访问

     - 外部类访问局部内部类的成员：创建对象再访问。

       ```java
       package test.test20230102;
       
       public class LocalInnerClass {
           public static void main(String[] args) {
            LocalInnerClass localInnerClass = new LocalInnerClass();
             Out02 out02 = localInnerClass.new Out02();
             out02.m1();
           }
           class Out02 {//外部类
               private int n1 = 100;
       
               private void m2() {//私有方法
                   System.out.println("out02 m2()");
               }
       
               public void m1() {
                   final class Inner02 {//局部内部类
       
                       public void f1() {
                           System.out.println("n1 = " + n1);
                           m2();
                       }
                   }
       
                   //外部类在方法中，可以创建Inner02对象，然后调用方法即可
                   Inner02 inner02 = new Inner02();
                   inner02.f1();
               }
           }
       }
       //n1 = 100
       //out02 m2()
       
       ```

       