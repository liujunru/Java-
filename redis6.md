### 1. nosql数据库简介

#### 引入nosql数据库

##### 多台服务器登录信息session的处理，将session存在哪

  1. 存在cookie里
      - 不安全
      - 网络负担效率低
  2. 存在文件服务器或数据库里
      - 大量的IO效率问题
  3. session复制，将session在所有的服务器上都复制一份。
      - session数据冗余，节点越多浪费越大
  4. 缓存数据库
      -  完全在内存中
      - 数据结构简单

### 2. redis概述

##### redis是单线程+多路IO复用技术
多路复用是指使用一个线程来检查多个文件描述符（Socket）的就绪状态，比如调用select和poll函数，传入多个文件描述符，如果有一个文件描述符就绪则返回，否则阻塞直到超时。得到就绪状态进行真正的操作可以在同一个线程里执行，也可以启动线程执行（比如使用线程池）

![多路复用](https://gitee.com/liujunrull/image-blob/raw/master/QQ截图20220601095950.png)

mamcached是多线程+锁

### 3. redis常用数据类型

字符串String、列表list、集合set、哈希Hash、有序集合Zset

#### 3.1 redis键key命令

  - keys *：查看当前库所有key（匹配：keys * 1）
  - exists key:判断某个key是否存在，返回1表示存在，0不存在
  - type key：查询你的key是什么类型
  - del key：删除指定key数据
  - unlink key：根据value选择非阻塞删除，仅将key从keyspace元数据中删除，真正的删除会在后续异步操作
  - expire key 10：设置key的过期时间为10秒钟
  - ttl key：查看还有多少秒过期，-1表示永不过期，-2表示已过期

  其他命令

    - select命令 切换数据库
    - dbsize 查看当前数据库的key的数量
    - flushdb 清空数据库
    - flushall 清空所有数据库

#### 3.2 Redis字符串（String）

redis最基本的类型，一个redis中字符串value最多可以是512M。
String类型是二进制安全的。意味着Redis的String类型可以包含任何数据。比如jpg图片或者序列化的对象

##### 3.2.1 常用命令

![添加键值对](https://gitee.com/liujunrull/image-blob/raw/master/QQ截图20220601104357.png)

set key  value:添加键值对

   - NX：当数据库key不存在时，可以将key-value添加数据库
   - XX：当数据库key存在时，可以将key-value添加数据库，与NX参数互斥
   - EX：key的超时秒数
   - PX：key的超时毫秒数，与EX互斥

get key:查询对应键值

append key value：将给定的<value>追加到原值的末尾，返回总长度

strlen key：获得值的长度

setnx key value：只有当key不存在时，设置key的值。分布式锁

incr key：将key中存储的数字值增1,。只能对数字值操作，如果为空，新增值为1

decr key：将key中存储的数字值减1,。只能对数字值操作，如果为空，新增值为-1

incrby/decrby key 步长：将key中存储的数字值增减。自定义步长

![java多线程](https://gitee.com/liujunrull/image-blob/raw/master/QQ截图20220601110012.png)

mset key1 value1 key2 value2：同时设置一个或多个key-value对

mget key1 value1 key2 value2：同时获取一个或多个key-value对

msetnx key1 value1 key2 value2：同时设置一个或多个key-value对，当给定的key都不存在

**原子性，有一个失败则都失败**

getrange key 起始位置 结束位置：获得值的范围，蕾丝substring，前后值都包括

setrange key 起始位置 value:用value复写<key>所储存的字符串值，从<起始位置>开始（索引从0开始）

setex key 过期时间 value：设置键值的同事，设置过期时间，单位秒

getset key value：以新换旧，设置了新值同时获得旧值

##### 3.2.2 value数据结构

String的数据结构为简单动态字符串（Simple Dynamic String，缩写SDS）。是可以修改的字符串，内部结构实现上累死与java的ArrayList，采用**预分配冗余空间**的方式来减少内存的频繁分配。

![预分配](https://gitee.com/liujunrull/image-blob/raw/master/QQ截图20220601111722.png)

内部为当前字符串实际分配的空间capacity一般要高于实际字符串长度len。当字符串长度小于1M时，扩容都是加倍现有的空间，如果超过1M，扩容时一次只会多扩1M的空间。注意字符串最大长度512M。

#### 3.3 redis列表（List）

单键多值

redis列表时简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。

他的底层实际是个双向链表，对两段的操作性能很高，通过索引下标的操作中间的节点性能会较差。

##### 3.3.1 常用命令

lpush/rpush key value1 value2 value3 ...：从左边/右边插入一个或多个值

lpop/rpop key：从左边/右边吐出一个值。值在键在，值光键亡。

rpoplpush key1 key2：从key1列表右边吐出一个值，插到key2列表左边

lrange key start stop：按照索引下标获得元素（从左到右）

lrange mylist 0 -1:0左边第一个，-1右边第一个（0-1表示获取所有）

lindex key index:按照索引下标获得元素（从左到右）

llen key：获得列表长度

linsert key before value newvalue：在value后面插入newvalue，value有重置值在最左的value值处添加

lrem key n value：从左边删除n个value(从左到右)

lset key index value：将列表key下标为index的值替换为value

##### 3.3.2 数据结构

list的数据结构为快速链表**quickList**

首先在列表元素较少的情况下会使用一块连续的存储空间，这个结构是**ziplist**,也即是压缩列表。

它将所有的元素紧挨着一起存储，分配的是一块连续的内存。

当数据量较多的还是才会改成**quicklist**。

因为普通的链表需要的附加指针空间较大，会比较浪费空间。比如这个列表里存的只是int类型的数据，结构上海需要两个额外的指针prev和next。

![list数据结构](https://gitee.com/liujunrull/image-blob/raw/master/QQ截图20220601135021.png)

redis将链表和ziplist结合起来组成了quicklist，也就是将多个ziplist使用双向指针串起来使用。这一既5满足了快速的插入删除性能，又不会出现太大的空间冗余。

#### 3.4 redis集合（set）

redis的set是string类型的无序集合。它底层其实是一个value为null的hash表，所以添加、删除、查找的复杂度都是O(1)。

##### 3.4.1 常用命令

sadd key value1 value2 ...：将一个或多个元素加入到集合Key里，已经存在的member元素将被忽略。

smember key value：判断集合key是否为含有该value值，有1，没有0

scard key：返回该集合的元素个数

srem key value1 value2 ...：删除集合中的某个元素

spop key：随机从该集合中吐出一个值

srandmember key n：随机从该集合中取出n个值。不会从集合中删除。

smove source destination value:把集合中的一个value从source移动到destination

sinter key1 key2：返回两个集合的交集元素

sunion key1 key2:返回两个集合的并集元素

sdiff key1 key2：返回两个集合的差集元素（key1中的，不包含的key2中的）

##### 3.4.2 数据结构

dict字典，用哈希表实现

#### 3.5 redis哈希（Hash）

一个键值对集合

redis hash是一个string类型的field和value的映射表，hash特别适合用于存储对象，类似于java里面的map<String,Object>

用户ID为查找的key，存储的value用户对象包含姓名、年龄、生日等信息，如果用普通的key/value结构来存储

![hash存储结构](https://gitee.com/liujunrull/image-blob/raw/master/202206011531608.png)

##### 3.5.1 常用命令

hset user:1001 id 1

hset key field value：给key集合中的field键赋值value

hget key1 field ：从key1集合field去除value

hmset key1 field1 value1 field2 value2 ... ：批量设置hash的值。新版本hmset已弃用，hset也可实现

hexists key1 field：查看哈希表key中，给定域field是否存在

hkeys key：列出该hash集合所有field

hvals key：列出该hash集合所有的value

hinsrby key field incremnet：为哈希表key中的域field的值加上增加incremnet

hsettnx key field value ：将哈希表key中的域field的值设置为value，当且仅当域field不存在

##### 3.5.2 数据结构

两种数据结构：ziplist（压缩列表），hashtable（哈希表）。当field-value长度较短且个数较少时，使用ziplist，否则使用hashtable。

#### 3.6 redis有序集合Zset（sorted set)

与普通集合set非常相似，是一个没有重复元素的字符串集合

不同之处是有序集合的每个成员都关联了一个评分（score），这个评分被用来按照从最低分到最高分的方式排序集合中的成员。集合的成员是唯一的，但是评分可以是重复的。

因为元素是有序的，所以你可以很快的根据评分或者次序来获取一个范围的元素

访问有序集合的中间元素也是非常快的，因此你能够使用有序集合作为一个每天重复成员的智能列表。

##### 3.6.1 常用命令

zadd key score1 value1 score2 value2 ...：将一个或多个member元素及其score值加入到有序集key当中

zrange key start stop [WITHSCORES]：返回有序集key中，下标在start stop 之间的元素。带WITHSCORES，可以让分数一起和值返回到结果集。

zrangebyscore minmax [WITHSCORES] [limit offset count]：返回有序集key中，所有score值结余min和max之间（包括等于min和max）的成员。

zincrby key increment value：为元素的score加上增量increment

zrem key value ：删除该集合下，指定值的元素

zcount key min max：统计该集合，分数区间内的元素个数

zrank key value：返回该值在集合中的排名，从0开始

##### 3.6.2 数据结构

sortedset(zset)是redis提供的一个非常特别的数据结构，一方面它等价于java的数据结构Map<String,Double>，可以给每个原色value赋予一个权重score，另一方面他又类似于TreeSet，内部的元素会按照权重score进行排序，可以得到每个元素的名次，还以为通过score的范围来获取元素的列表。

zset底层使用了两个数据结构。

（1) **hash**,hash的作用就是关联元素value和权重score，保障元素value的唯一性，可以通过元素value找到相应的score值。

（2）**跳跃表**，跳跃表的木笔在于给元素value排序，根据score的范围获取元素列表。

![跳表查找51](https://gitee.com/liujunrull/image-blob/raw/master/202206011607606.png)

### 4. redis_jedis测试

#### 4.1 需要的jar包

```java
<dependency>
         <groupId>redis.clients</groupId>
         <artifactId>jedis</artifactId>
         <version>3.2.0</version>
     </dependency>
```







