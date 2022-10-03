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

