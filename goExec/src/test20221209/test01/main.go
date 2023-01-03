package main
import(
	"fmt"
	"encoding/json"
) 
type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}
func main(){
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	//将monster变量序列化为json字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
	//jsonStr {"name":"牛魔王","age":500,"skill":"芭蕉扇"}
}