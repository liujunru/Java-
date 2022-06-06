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

#### 4.2 测试

```
    public static void main(String[] args) {
        //创建Jedis对象,param1要访问的redis服务的主机地址，param2端口号
        Jedis jedis = new Jedis("192.168.53.10",6379);
        //测试
        String ping = jedis.ping();
        System.out.println(ping);


    }
```

connect time out(连接超时)：需要打开防火墙或者6379端口

```
1)systemctl status firewall 查看防火墙状态 active是否为running,若是
systemctl stop firewall 关闭防火墙
firewall-cmd --reload 重启防火墙

2)firewall-cmd --query-port=6379/tcp 查看6379端口号是否打开
firewall-cmd --query-port=6379/tcp 打开6379端口号
```

命令行有的指令jedis都有对应方法

### 5、springboot整合redis

#### 5.1 依赖

```
    <!--redis-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-redis</artifactId>
        </dependency>
        <!--spring2.X集成redis所需common-pool2-->
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-pool2</artifactId>
            <version>2.6.0</version>
        </dependency>
```

#### 5.2 配置文件

```
#redis服务器地址
spring.redis.host=192.168.53.10
#redis服务器连接端口
spring.redis.port=6379
#redis数据库索引（默认为0）
spring.redis.database=0
#连接超时时间（毫秒）
spring.redis.timeout=1800000
#连接池最大连接数（负值表示没有限制）
spring.redis.lettuce.pool.max-active=20
#最大阻塞等待时间（负数表示没限制）
spring.redis.lettuce.pool.max-wait=-1
#连接池中的最大空闲连接
spring.redis.lettuce.pool.max-idle=5
#连接池中的最小空闲连接
spring.redis.lettuce.pool.min-idle=0
```

#### 5.3 配置类

```
package com.atguigu.config;

import com.fasterxml.jackson.annotation.JsonAutoDetect;
import com.fasterxml.jackson.annotation.PropertyAccessor;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.jsontype.impl.LaissezFaireSubTypeValidator;
import org.springframework.cache.CacheManager;
import org.springframework.cache.annotation.CachingConfigurerSupport;
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.cache.RedisCacheConfiguration;
import org.springframework.data.redis.cache.RedisCacheManager;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.RedisSerializationContext;
import org.springframework.data.redis.serializer.RedisSerializer;
import org.springframework.data.redis.serializer.StringRedisSerializer;

import java.time.Duration;

@EnableCaching
@Configuration
public class RedisConfig extends CachingConfigurerSupport {
    @Bean
    public RedisTemplate<String,Object> redisTemplate(RedisConnectionFactory factory){
        RedisTemplate<String,Object> template = new RedisTemplate<>();
        RedisSerializer<String> redisSerializer = new StringRedisSerializer();
        Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new Jackson2JsonRedisSerializer(Object.class);
        ObjectMapper om = new ObjectMapper();
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        om.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
        jackson2JsonRedisSerializer.setObjectMapper(om);
        template.setConnectionFactory(factory);
        //key序列化方式
        template.setKeySerializer(redisSerializer);
        //value序列化方式
        template.setHashValueSerializer(jackson2JsonRedisSerializer);
        return template;
    }

    @Bean
    public CacheManager cacheManager(RedisConnectionFactory factory){
        RedisTemplate<String,Object> template = new RedisTemplate<>();
        RedisSerializer<String> redisSerializer = new StringRedisSerializer();
        Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new Jackson2JsonRedisSerializer(Object.class);
        //解决查询缓存转换异常的问题
        ObjectMapper om = new ObjectMapper();
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        om.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
        jackson2JsonRedisSerializer.setObjectMapper(om);
        //配置序列化（解决乱码的问题），过期时间600秒
        RedisCacheConfiguration configuration = RedisCacheConfiguration.defaultCacheConfig().entryTtl(Duration.ofSeconds(600))
                .serializeKeysWith(RedisSerializationContext.SerializationPair.fromSerializer(redisSerializer))
                .serializeValuesWith(RedisSerializationContext.SerializationPair.fromSerializer(jackson2JsonRedisSerializer))
                .disableCachingNullValues();
        RedisCacheManager cacheManager = RedisCacheManager.builder(factory)
        .cacheDefaults(configuration)
        .build();
        return cacheManager;
    }
}
```

#### 5.4 测试

```
package com.atguigu.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("redisTest")
public class RedisTest {
    
    @Autowired
    private RedisTemplate redisTemplate;
    
    @GetMapping
    public String testRedis(){
        //设置值到redis
        redisTemplate.opsForValue().set("name","lucy");
        //从redis里获取值
        Object name = redisTemplate.opsForValue().get("name");
        return name.toString();
    }
    
}
```

### 6.redis事务

redis事务时一个单独的隔离操作：事务中的所有命令都会序列化、按顺序的执行。事务在执行的过程中，不会被其他客户端发送来的命令请求所打断。

redis事务的主要作用就是串联多个命令防止别的命令插队。

#### 6.1 事务命令Mulit、Exec、discard

从输入Mulit命令开始，输入的命令都会依次进入命令队列中（组队阶段）queued，但不会执行，直到注入Exec（执行阶段）后，redis会将之前命令队列中的命令依次执行。

组队的过程中可以通过discard来放弃组队。

![指令阶段](https://gitee.com/liujunrull/image-blob/raw/master/202206051622359.png)

#### 6.2 事务的错误处理

1. 组队中某个命令出现报告错误，执行时整个队列都会被取消。
2. 执行阶段某个命令出现错误，则只有报错的命令不会执行，其他的命令都会执行，不会回滚。

#### 6.3 事务冲突

##### 6.3.1 悲观锁

每次拿数据时认为别人会修改，所以每次拿数据时都会上锁，这样别人想拿数据就会block直到他拿到锁。传统的关系型数据库都用到了很多这这种锁机制，比如行锁、表锁、读锁、写锁等，都是在操作之前先上锁。

##### 6.3.2 乐观锁

每次拿数据时认为别人不会修改，所以每次拿数据时不会上锁，但是在更新时会判断一下此期间有没有别人去更新数据，可以使用版本号机制。乐观锁适用于多读的应用类型，这样可以提高吞吐量。redis就是利用这种check-and-set机制实现事务的。

在执行multi之前，先执行watch key1 [key2],可以监视一个或多个key，如果在事务执行之前这个或这些key被其他命令所改动，那么事务将被打断。 

unwatch 取消监视

#### 6.4 redis事务三特性

- 单独的隔离操作
  
    - 事务中的所有命令都会被序列化、按顺序执行。事务在执行过程中，不会被其他客户端发送来的请求打断

- 没有隔离级别的概念

    - 队列中的命令没有被提交之前都不会实际被执行，因为事务提交前人恶化指令都不会被实际执行。

- 不保证原子性

    - 事务中如果有一条命令执行失败，其后的命令仍然会被执行，没有回滚。

#### 6.5 秒杀案例

模拟并发工具-ab工具/Jmeter

```
服务器端访问本地
安装指令：yum install httpd-tools
测试指令：ab -n [请求次数] -c [并发数] [请求地址]
 如：ab -n 1000 -c 100 http://本机ip:8080/seckill
 -p: post请求
 -T: content-type
 如：ab -n 1000 -c 100 -p ~/postfile -T application/x-www-form-urlencoded http://本机ip:8080/seckill
```

- 连接超时问题-连接池
```
package com.atguigu.jedis;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

public class JedisPoolUtil {
    private static volatile JedisPool jedisPool = null;
    private JedisPoolUtil(){}

    public static JedisPool getJedisPoolInstance(){
        if(null == jedisPool){
         synchronized (JedisPoolUtil.class){
             if(null == jedisPool){
                 JedisPoolConfig poolConfig = new JedisPoolConfig();
                 poolConfig.setMaxTotal(200);
                 poolConfig.setMaxIdle(32);
                 poolConfig.setMaxWaitMillis(100*1000);
                 poolConfig.setBlockWhenExhausted(true);
                 poolConfig.setTestOnBorrow(true);

                 jedisPool = new JedisPool(poolConfig,"192.168.53.10",6379,60000);
             }
         }
        }
        return jedisPool;
}
    //释放资源
    public static void release(JedisPool jedisPool, Jedis jedis){
        if(null != jedis){
            jedisPool.isClosed();
        }
    }
}
```

- 超卖问题-乐观锁

```
public static boolean doSecKill(String uid,String prodid) throws IOException {
        //1.uid和prodid非空判断，都为空秒杀未开始
        if(uid == null || prodid == null){
            return false;
        }
        //2.连接redis
        //Jedis jedis = new Jedis("192.168.53.10",6379);
        //通过连接池得到jedis对象
        JedisPool jedisPool = JedisPoolUtil.getJedisPoolInstance();
        Jedis jedis = jedisPool.getResource();
        //3.拼接key
        //3.1 库存key
        String kcKey = "sk:" + prodid + ":qt";
        //3.2 秒杀成功用户key
        String userKey = "sk:" + prodid + ":user";
        //监视库存
        jedis.watch(kcKey);
        //4.获取库存，如果库存null，秒杀还没有开始
        if(jedis.get(kcKey) == null){
            System.out.println("秒杀还未开始，请等待！");
            jedis.close();
            return false;
        }
        //5.判断用户是否重复秒杀,使用set避免重复
        if(jedis.sismember(userKey,uid)){
            System.out.println("已经秒杀成功，不能重复秒杀！");
            jedis.close();
            return false;
        }
        //6.判断商品数量，如果小于1秒杀结束
        if(Integer.parseInt(jedis.get(kcKey)) <= 0){
            System.out.println("秒杀已经结束");
            jedis.close();
            return false;
        }
        //7.秒杀过程,加入事务
        Transaction multi = jedis.multi();
        //组队操作，库存-1。把秒杀成功用户加入到清单里+1
        multi.decr(kcKey);
        multi.sadd(userKey,uid);
        //执行
        List<Object> list = multi.exec();
        if(list == null || list.size() == 0){
            System.out.println("秒杀失败");
            jedis.close();
            return false;
        }
        //7.1 库存-1
        //jedis.decr(kcKey);
            //7.2 把秒杀成功用户加入到清单里
        //jedis.sadd(userKey,uid);
        System.out.println("秒杀成功...");
        jedis.close();
        return true;
    }
```

- 库存遗留问题-LUA脚本

    将复杂的或者多步的redis操作，写为一个脚本，一次提交给redis执行，减少吩咐连接redis的次数，提升性能

    LUA脚本是类似redis事务，有一定的原子性，不会被其他命令插队，可以完成一些redis事务性的操作

    但是注意redis的lua脚本功能，只有在redis 2.6以上的版本猜哦可以使用

    利用Lua脚本淘汰用户，解决超卖问题。

    redis 2.6以后，通过lua脚本解决争抢问题，实际上是redis利用单线程的特性，用任务队列的方式解决多任务并发问题。

 ```
        static String secKillScript = "local userid=KEYS[1];" +
            "local prodid=KEYS[2];" +
            "local qykey='SecKill:'..prodid..\":kc\";" +
            "local userskey='Seckill:'..prodid..\":user\";" +
            "local userExists=redis.call(\"sismeber\",usersKey,userid);" +
            "if tonumer(userExists)==1 then" +
            "   return 2;" +
            "end" +
            "local num=redis.call(\"get\",qtkey);" +
            "if tonumber(num)<=0 then" +
            "   return 0;" +
            "else" +
            "   redis.call(\"desc\",qukey);" +
            "   redis.call(\"sadd\",usersKey,userid);" +
            "end" +
            "return 1";
   
    public static boolean doSecKill(String uid,String prodid) throws IOException {
        JedisPool jedisPool = JedisPoolUtil.getJedisPoolInstance();
        Jedis jedis = jedisPool.getResource();

        String sha1 = jedis.scriptLoad(secKillScript);
        Object result = jedis.evalsha(sha1,2,uid,prodid);

        String reString = String.valueOf(result);
        if("0".equals(reString)){
            System.out.println("已抢空！");
        }else if("1".equals(reString)){
            System.out.println("抢购成功！");
        }else if("2".equals(reString)){
            System.out.println("该用户已抢过！");
        }else{
            System.err.println("抢购异常！");
        }

        jedis.close();
        return true;
    }
 ```

### 7. redis的持久化

分为RDB和AOF

#### 7.1 RDB

在指定的**时间间隔**内将内存中的**数据集快照**写入磁盘

##### 7.1.1 备份如何执行的

redis会单独创建（fork）一个子进程来进行持久化，会先将数据写入到一个**临时文件**中，待持久化过程都结束了，再用这个临时文件替换上次持久化好的文件。整个过程中，主进程是不进行任何IO操作的，这就确保了极高的性能，如果需要规模的数据的恢复，且对于数据恢复的完整性不是很敏感，那RDB方式要比AOF方式更加高效。RDB的缺点是**最后一个持久化后的数据可能丢失**。

##### 7.1.2 Fork

- Fork的作用是复制一个于当前进行一样的进程。新进程的所有数据（变量、环境变量、程序计数器等）数值都和原进程一致，但是是一个全新的进程，并作为原进程的子进程
- 在linux程序中，fork()会产生一个和父进程完全相同的子进程，但子进程在此后多次exec系统调用，出于效率考虑，linux引入了“写时复制技术”
- 一般情况父进程与子进程会共用一段物理内存，只有进程空间的各段的内容要发生变化时，才会将父进程的内容复制一份给子进程。

##### 7.1.3 config文件

stop-writes-on-bgsave-error:当redis无法写入磁盘的话，直接关掉redis的写操作。推荐yes

rdbcompression：对于存储到磁盘中的快照，可以设置是否进行压缩存储。如果是的话，redis会使用**LZF算法**进行压缩。如果不想消耗CPU进行压缩，可以设为No。推荐yes

rdbchecksum:检查完整性。在存储快照后，可以让redis使用CRC64算法来进行数据校验。但是这样做会增加大约10%的性能消耗。推荐yes

save 秒钟：写操作次数。RDB时整个内存的压缩过的snapshot，RDB的数据结构，可以配置复合的快照触发条件。手动保存

bgsave：redis会在后台异步进行快照操作，快照同时还可以响应客户端请求。

可以通过lastsave命令来获取最后一次成功执行快照的时间

##### 7.1.4 优势

- 适合大规模的数据恢复
- 对数据完整性和一致性要求不高更适合使用
- 节省磁盘空间
- 恢复速度快

##### 7.1.5 劣势

- fork的时候，内存中的数据被克隆了一份，大致2倍的膨胀性服用考虑
- 虽然redis在fork时使用了写时拷贝技术，但是如果数据庞大时还是比较消耗性能
- 在备份周期在一定时间间隔做一次备份，所有如果redis意外down掉的话，就是丢失最后一次快照后所有的修改


#### 7.2 AOF(Append Only File)

以日志的形式来记录每个写操作（增量保存），将redis执行过的所有写指令记录下来（读操作不记录），只许追加文件但不可以改写文件，redis启动之初会读取该文件重新构建数据，换言之，redis重启的话就根据日志文件的内容将写指令从前到后执行一次以完成数据的恢复工作。

##### 7.2.1 AOF持久化流程

1) 客户端的请求命令会被append追加到AOF缓冲区内
2) AOF缓冲区根据AOF持久化策略[always,everysec,no]将操作sync同步到磁盘的AOF文件中
3) AOF文件大小超过重写策略或手动重写时，会对AOF文件rewrite重写，压缩AOF文件容量
4) redis服务重启时，会重新load加载AOF文件中的写操作达到数据恢复的目的

#### 7.2.2 AOF默认不开启,优先读取AOF数据

可以在redis.conf打开，appendonly yes,配置文件名称，默认为appendonly.aof

AOF文件的保存路径同RDB文件路径一致。

``` 
查看redis进程：ps -ef | grep redis
杀掉redis进程：kill -9 端口号
```
##### 7.2.3 AOF的异常恢复

如果AOF文件损坏，通过/usr/local/bin/redis-check-aof--fix appendonly.aof进行恢复

备份被写坏的AOF文件

重启redis，然后重新加载

##### 7.2.4 AOF同步频率设置

appendfsync always:始终同步，每次redis的写入都会立刻计入日志；性能较差但数据完整性较好

appendfsync everysec:每秒同步，每秒计入日志一次，如果宕机，本秒的数据可能丢失。

appendfsync no：不主动进行同步，把同步交给操作系统

##### 7.2.5 Rewrite压缩

AOF采用文件追加方式，文件会越来越大。为了避免出现此种情况，新增加了**重写**机制。当AOF文件的大小超过所设定的阈值时，redis会启动AOF文件的内容压缩，只保留可以恢复数据的最小指令集，可以使用命令bgrewriteaof

**重写原理**

AOF文件持续增长而过大时，会fork一条新进程来将文件重新（也是先写临时文件最后在rename），redis4.0之后的重写，是指就是把rdb的快照，以二进制的形式附在新的AOF头部，作为已有的历史数据，替换到原来的流水账操作。

no-appendfsync-on-rewrite

触发条件：默认AOF文件大小是上次rewrite后大小的一倍且文件大于64M

auto-aof-rewrite-percentage：设置重写的基准值。文件达到100%时开始重写（大小是原来重写后文件的2倍）

auto-aof-rewrite-min-size:设置重写的基准值。最小64M

系统载入时或许上次重写完毕时，redis会记录此时AOF大小，设为**base_size**

##### 7.2.6 优势

- 备份机制更稳健。丢失数据概率耕地
- 可读的日志文本，通过操作AOF文件，可以处理误操作。

##### 7.2.7 劣势

- 比起RDB占用更多的磁盘空间
- 恢复备份速度要慢
- 每次读写都操作的话有一定的性能压力
- 存在个别bug，造成不能恢复

####  7.3 使用哪个

官方推荐两个都启动。如果对数据不敏感，可单独使用RDB。不建议单独使用AOF，因为可能胡i出现bug。如果只是做纯内存缓存，可以都不用

### 8. 主从复制

主机数据更新后根据配置和策略，自动同步到备机的master/slaver机制，Master以写为主，Slave以读为主。

读写分离，容灾快速恢复（一主多从，当主机挂断，考虑集群）

#### 8.1 简单一主两从搭建

1) 创建/myredis文件夹
2) 复制redis.conf配置文件到文件夹中
3) 配置一主两从，创建三个配置文件

    - redis6379.conf
    - redis6380.conf
    - redis6381.conf

4) 在三个配置文件写入内容
   include/myredis/redis.conmf
   prifile/var/run/redis_6379.pid
   port6379
   dbfilename dump6379.rbd

```
info replication：查看主从信息
```

在从机上执行slaveof 主机Ip 端口号 设为从机

#### 8.2 复制原理

1) slave启动成功连接到master后发送一个sync命令
2) master接到命令启动后台的存盘进程，同时收集所有接收到的用于修改数据集命令，在后台进程执行完毕之后，master将传送这个数据文件到slave，以完成一次完全同步
3) 全量复制：slave服务在接受到数据库文件数据后，将其存盘并加载到内存中。

    增量复制：master继续将新的所有收集到的修改命令依次传给slave，完成同步
4) 只要重新连接master，一次完全同步（全量复制）将被自动执行。

#### 8.3 三个特点

1) 一主两仆

    主服务器挂掉之后再重新启动还是主服务器，从服务器不会晋升为主服务器

2) 薪火相传

    从服务器下还可以再设置从服务器

3) 反客为主

    slaveof no no：设置从服务器为主服务器

#### 8.4 哨兵模式

反客为主的自动版，能够后台监控主机是否故障，如果故障了根据投票数自动将从库转换为主库。

#### 8.4.1 创建哨兵模式

    - 在自定义的/myredis目录下新建sentinel.conf文件，名字不能错
    - 配置哨兵，填写内容
      - sentinel monitor mystater 127.0.0.1 6379 1
      - 其中mymaster为监控对象起的服务器名称，1为至少有多少个哨兵同意迁移
    - 启动哨兵
      - redis -sentinel /myredis/sentinel.conf

#### 8.4.2 复制延迟

由于所有的写操作都是先在master上操作，然后更新到slave上，所以从master同步到slave机器有一定的延迟，当系统很繁忙的时候，延迟问题更加严重，slave机器数量的增加也会使这个问题更严重

#### 8.4.3 选举规则

- 选择优先级靠前的，优先级在redis.conf默认：replica-priority 100,值越小优先级越高,0的从机永不会被选为主机
- 选择偏移量大的，偏移量指获得原主机数据最全的
- 选择runid最小的从服务，每个redis实例启动后都会随机生成一个40位的runid

### 9. 集群

容量不够，redis如何进行扩容？

并发写操作，redis如何分摊？

另外，主从模式，薪火相传模式，主机宕机，导致ip地址发生变化，应用程序中配置需要修改对应的主机地址、端口等信息

之前通过**代理主机**来解决，但是redis3.0提供了解决方案，就是无中心化集群配置。

#### 9.1 什么是集群

redis集群实现了对redis的水平扩容，即启动N个redis节点，将整个数据库分布存储在这N个节点中，每个节点存储总数据的1/N。

redis集群通过分区来提供一定程度的可用性：即使集群中有一部分节点生效或者无法进行通讯，集群也可以继续处理命令请求。

#### 9.2 创建集群

    - 制作6个实例，端口号6379-6391
    - redis cluster配置修改
      - include/home/bigdata/redis.conf
      - port 6379
      - pidfile "var/run/redis_6379.pid"
      - dbfilename "dump6379.rdb"
      - dir "/home/bigdarta/redis_cluster"
      - logfile "/home/bigdata/redis_cluster/redis_err_6379.log"
    - 将6个集群合成一个集群
      - 组合之前确保所有redis实例启动，nodes-xxxx.conf文件都生成正常
      - cd /opt/redis-6.2.1/src
      - redis -cli --cluster created --cluster-replicas 192.168.11.101:6379到6391
      - 此处不要用127.0.0.1，要用真实ip地址
      - replicas 1 采用最简单的方式配置集群，一台主机，一台从机
    - redis-cli -c -p 6379 采用集群策略连接，设置数据会自动切换到相应的写主机
    - 通过cluster nodes 查看集群信息

#### 9.3 什么是slot

一个redis集群包含16384个插槽，数据库中每个键都属于这16384个插槽中的一个

集群使用公式CRC16(key)%16384来计算键key属于哪个槽，其中CRC6(key)语句用于计算键key的CRC16的校验和

集群中每一个节点负责处理一部分插槽

不在一个slot下的键值，不能使用mget、mset大声多键操作。

可以通过{}来定义组的概念，从而使key中{}内相同内容的键值对放到一个slot中去。

mset k1{cust} v1 k2{cust} v2

mget k1{cust}  k2{cust} 

#### 9.4 查询集群中的值

cluster getkeysinslot slot count:返回count个slot槽中的键

#### 9.5 故障恢复

如果主节点下线，在15秒重启不影响主关系，超过15秒还未恢复，从节点自动升为主节点。主节点恢复后，会变成从机

cluster-require-full-coverage为默认是yes，如果一段插槽的主从都挂掉，整个集群都挂掉，no，不影响整个集群。

集成jedis

```java
public class JedisClusterTest{
    public static void ain(String[] args){
        Set<HostAndPort> set = new HashSet<HostAndPort>;
        set.add(new HostAndPort("192.168.31.211",6379);
        JedisCluster  jesidCluster = new JedisCluster(set);
        jedisCluster.set("k1","v1");
        sout(jedisCluster.get("k1"));
    }
}
```

#### 9.6 优点

实现扩容

分摊压力

无中心配置相对简单

#### 9.7 缺点

多键操作不支持

多键的redis事务不支持。lua脚本不支持

由于集群方案出现较晚，很多公司已经采用了其他的集群方案，而代理或者客户端分片的方案想要迁移到redis cluster，需要整体迁移而不是逐步过渡，复杂度较高。

### 10. redis应用问题

#### 10.1 缓存穿透

redis缓存命中率低，导致大量数据请求数据库，造成数据库压力过大而崩溃。

**解决方案**

- 对空值缓存：如果一个查询返回数据为空，将空值进行缓存，设置空值的国企时间很短，最长不超过5分钟。
- 设置可访问的白名单：使用bitmaps类型定义一个可以访问的名单，名单id作为bitmaps的偏移量，每次访问和Bitmap中的id进行比较，如果id查不到，进行拦截，不允许访问。
- 布隆过滤器：实际上是一个很长的二进制向量（位图）和一系列随机映射函数（哈希函数）。本质是bitmaps，可以用于检测一个元素是否在一个集合中。优点是空间效率和查询实际都远远超过一般的算法，缺点是有一定的误识别率和删除困难。
- 进行实时监控：当发现redis命中率急剧降低，需要排查访问对象和访问的数据，和运维人员配合，设置黑名单限制服务。

#### 10.2 缓存击穿

redis中某个热门key过期吗，此时有大量访问使用这个key

**解决方案**

- 预先设置热门数据：在redis访问高峰之前，把一些人们数据提前存入到redis里面，加大这些热门数据key的时长。
- 实时调整 ：现场监控热门数据，实时调整key的过期时长。
- 使用锁：在缓存失效的时候（判断拿出来的值为空），不是立即去load db。先使用缓存工具的某些带成功操作返回值的操作（比如redis的sentnx）

#### 10.3 缓存雪崩

极少时间内，查询大量key集中过期的情况

**解决方案**

- 构建多级缓存架构：nginx缓存+redis缓存+其他缓存（ehcache等）
- 使用锁或队列：使用加u送或者队列的方式保证不会有大量的线程对数据库进行一次性读写，从而避免失效时大量的并发请求落到底层存储系统上。不适用于高并发情况。
- 设置过期标志更新缓存：记录缓存数据是否过期（设置提前量），如果过期会触发通知另外的线程在后台去更新实际key的缓存
- 将缓存失效实际分散开：比如可以在原有的失效实际基础上增加一个随机值，比如1-5分钟随机，这样每个缓存的过期实际重复率会降低。

### 11. 分布式锁

跨JVM的互斥机制来控制共享资源的访问

分布式锁的主流实现方案：

    - 基于数据库实现分布式锁
    - 基于缓存（redis)，性能最好
    - 基于Zookeeper，可靠性最高

#### 11.1 redis实现分布式锁

上锁并设置过期时间

set sku:1:info "ok" NX PX 10000

EX second:设置键的过期实际为second秒

set key value EX second效果等同于 set EX key second value

#### spring boot整合

```java
import org.springframework.web.bind.annotation.GetMapping;

@GetMapping("testLock")
public void testLock(){
    String uuid = UUID.random().toString()
    //1.获取锁，setne
        Boolean lock = redisTemplate.opsForValue().setIfAbsent("lock",uuid,3,TimeUnit.SE CONDS);
        //2.获取锁成功查询num的值
        if(lock){
           Object value = redisTemplate.opsForValue().get("num");
           //2.1 判断num 为空返回
        if(StringUtils.isEmpty(value)){
            return;
        }
        //2.2 如果有值转为int
        int num = Integer.parseInt(value + "");
        //2.3 把redis的num+1
        redisTemplate.opsForValue().set("num",++num);
        //释放锁
        String lockUuid = (String) redisTemplate.opsForValue().get("lock");
        if(uuid.equals(lockUuid){
        redisTemplate.delete("lock");
        }
     
        }else{
            //3.获取锁失败，每隔0.1秒后再 获取
            try{
               Thread.sleep(100);
               testLock();
        }catch(InterruptedException e){
                e.printStackTrace();
        }
    }
}
```

LUA脚本保证释放锁的原子性





















