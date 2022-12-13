package main
import(
	"fmt"
) 
func main(){
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "aaa"
	a["no2"] = "bbb"
	a["no1"] = "ccc"
	fmt.Println(a)//map[no1:ccc no2:bbb]
}