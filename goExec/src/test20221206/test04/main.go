package main
import(
	"fmt"
) 

func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}
func main(){
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	// str= helloa
	// 11
	// str= helloaa
	// 13
	// str= helloaaa
	// 16

}