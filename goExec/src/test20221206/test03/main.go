package main
import(
	"fmt"
) 
var (
	Fun1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)
func main(){
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10,20)

	fmt.Println("res1=", res1)//res1= 30

	a := func (n1 int, n2 int) int {
		return n1 + n2
	}

	res2 := a(10, 30)

	fmt.Println("res2=", res2)//res2= 40

	res4 := Fun1(10, 30)
	fmt.Println("res4=", res4)//res4= 300

}