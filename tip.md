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

