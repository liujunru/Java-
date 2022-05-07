- 删除重复数据保留一条sql

```mysql
delete from testdelete where id not in (select a.id from(select min(id) as id from testdelete group by testname)as a);
```

