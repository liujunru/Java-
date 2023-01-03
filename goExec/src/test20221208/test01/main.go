package main
import(
	"fmt"
) 
func main(){
	//第一种方式
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "aaa"
	a["no2"] = "bbb"
	a["no3"] = "ccc"
	fmt.Println(a)//map[no1:aaa no2:bbb no3:ccc]

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "南京"
	fmt.Println(cities)//map[no1:北京 no2:上海 no3:南京]

	//第三种方式:如果闭合大括号跟最后一个元素在同一行不需要逗号
	//不在同一行需要加逗号
	heroes := map[string]string{
		"hero1" : "宋江",
		"hero2" : "吴用",
		"hero3" : "卢俊义",
	}
	fmt.Println(heroes)//map[hero1:宋江 hero2:吴用 hero3:卢俊义]
}