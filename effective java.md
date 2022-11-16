
如果类的构造器或者静态工厂中具有多个参数，设计这种类时， **Builder模式**是一种不错的选择.

### 用私有构造器或者枚举类型强化 Singleton 属性

私有构造器仅被调用一次 。 客户端的任何行为都不会改变这一点， 但要提醒一点：享有特权的客户端可以借助 AccessibleObject.setAccessible 方法，
通过反射机制调用私有构造器 。 如果需要抵御这种攻击，可以修改构造器，
让它在被要求创建第二个实例的时候抛出异常 。

#### 使用静态工厂模式创建单例
````java
public class Elvis{
  private sattic final Elvis INSTANCE = new Elvis();
  private Elvis(){}
  public static Elvis getInstance(){
    return INSTANCE;
  }
  public void leaveTheBuilding(){}
}
````

静态工厂方法的优势之一在于，它提供了灵活性 ：在不改变其 API 的前提下 ， 我们可以改变该类是否应该为 Singleton 的想法。工厂方法返回该类的唯一实例，但是，它很容易被修改， 比如改成为每个调用该方法的线程返回一个唯一的实例 。 第二个优势是，如果应用程序需要，可以编写一个泛型 S ingleton 工厂 （ generic singleton factory) 。使用静态工厂的最后一个优势是，可以通过方法引用（ method reference）作为提供者，比如Elvis : : instance 就是一个 Supplier<Elvis＞ 。除非满足以上任意一种优势 ，否则还是优先考虑公有域（ publi c-field ）的方法。

- 实现 Singleton 的第三种方法是声明一个包含单个元素的枚举类型：

````java
public enum Elvis{
  INSTANCE;

  public void leaveTheBuilding(){}
}
````

这种方法在功能上与公有域方法相似，但更加简洁，无偿地提供了序列化机制，绝对防止多次实例化，即使是在面对复杂的序列化或者反射攻击的时候 。虽然这种方法还没有广泛采用，但是单元素的枚举类型经常成为实现 Singleton 的最佳方法 。 注意，如果 Singleton必须扩展一个超类，而不是扩展 Enum 的时候，则不宜使用这个方法（虽然可以声明枚举去实现接口） 

### 优先考虑依赖注人来引用资源

依赖一个或多个底层资源的类，需要的是能够支持类的多个实例（在本例中是指 SpellChecker ），每一个实例都
使用客户端指定的资源（在本例巾是指同典） 。满足该需求的最简单的模式是， 当创建一个新的实例时 ， 就将该资源传到构造器中 。 这是依赖注入（ dependency injection ）的一种形式

````java
public class SpellChecker{
  private final Lexicon dictionary;

  public SpellChecker(Lexion dictionary){
    this.dictionary = Object.requireNonNull(dictionary);
  }
  public boolean isValid(String word){}
  public List<String> suggestions(String typo){}
}
````
### 避免创建不必要的对象
对于同时提供了静态 工厂方法 （ static factory method) （详见第 l 条）和构造器的不可变类，通常优先使用静态工厂方法而不是构造器，以避免创建不必要的对象。除了重用不可变的对象之外，也可以重用那些已知不会被修改的可变对象。

要优先使用**基本类型**而不是装箱基本类型，要当心无意识的自动装箱。

### 消除过期的对象引用

````java
public class Stack{
  private Object[] elements;
  private int size = 0;
  private static final int DEFAULT_INITIAL_CAPACITY = 16;

  public Stack(){
    elements = new Object[DEFAULT_INITIAL_CAPACITY ];
  }

  public void push(Object e){
    ensureCapacity();
    elements[size++] = e;
  }

  public Object pop(){
    if(size == 0)throw new EmptyStackException();
    return element[--size];
  }

  private void ensureCapacity(){
    if(elements.length == size) 
      elements = Arrays.copyOf(elements,2 * size + 1)
  }
}
````

如果一个栈先是增长 ，然后再收缩 ， 那么，从栈中弹出来的对象将不会被当作垃圾回收，即使使用栈的程序不再引用这些对象，它们也不会被回收 。 这是因为栈内部维护着对这些对象的过期引用 （ obsolete reference ） 。 所谓的过
期引用，是指永远也不会再被解除的引用 。 在本例中，凡是在 elements 数组的“活动部分”(active portion ）之外的任何引用都是过期的 。 活动部分是指 elements 中下标小于 size 的那些元素 。

永远不应该依赖终结方法或者清除方法来更新重要的持久状态。

如果类的对象中封装的资源（例如文件或者线程）确实需要终止，应该怎么做
才能不用编写终结方法或者清除方法呢？只需 让类实现 AutoCloseable，并要求其客户端在每个实例不再需要的时候调用 close 方法，一般是利用 try - with - resources 确保终止，即使遇到异常也是如此 。 值得提及的一个细节是，该实例必须记录下自己是否已经被关闭了： close 方法必须在一个私有域中记录下“该对象已经不再有效” 。 如果这些方法是在对象已经终止之后被调用，其他的方法就必须检查这个域，并抛出IllegalStateException 异常 。

