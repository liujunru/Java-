### 1.Feature.OrderedField
fastjson在把字符串解析成Json对象时，可以通过指定Feature.OrderedField来使Json对象中的元素按字符串中的顺序排列。

````java
public static void main(String[] args) {
        String str = "{\"b\":321,\"a\":123}";
        JSONObject json = JSONObject.parseObject(str);
        System.out.println(json);//结果为：{"a":123,"b":321}
         
        JSONObject json2 = JSONObject.parseObject(str, Feature.OrderedField);
        System.out.println(json2);//结果为：{"b":321,"a":123}
         
        String arrStr = "[{\"b\":321,\"a\":123}]";
        JSONArray array = JSONArray.parseArray(arrStr);
        System.out.println(array);//结果为：[{"a":123,"b":321}]
         
        int temp = JSON.DEFAULT_PARSER_FEATURE;
        JSON.DEFAULT_PARSER_FEATURE = Feature.config(JSON.DEFAULT_PARSER_FEATURE, Feature.OrderedField, true);
        JSONArray array2 = JSONArray.parseArray(arrStr);
        JSON.DEFAULT_PARSER_FEATURE = temp;
        System.out.println(array2);//结果为：[{"b":321,"a":123}]
    }
````

如果没有设置Feature.OrderedField则json对象中的元素按**字典顺序**排列，

如果设置了Feature.OrderedField则按**字符串中的顺序**排列。

JSON.DEFAULT_PARSER_FEATURE是全局的，如果不希望影响其他功能，建议在完成字符串解析后再把JSON.DEFAULT_PARSER_FEATURE设置为默认值。

