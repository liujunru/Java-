package main
import(
	"fmt"
) 
func main(){
	str := "hello,worldv地方的色!"
	for index, val := range str{
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
	//index=0,val=h
	// index=1,val=e
	// index=2,val=l
	// index=3,val=l
	// index=4,val=o
	// index=5,val=,
	// index=6,val=w
	// index=7,val=o
	// index=8,val=r
	// index=9,val=l
	// index=10,val=d
	// index=11,val=!

	str2 := "hekjndfr参数"
str3 := []rune(str2)
	for i:= 0; i < len(str3); i++{
		fmt.Printf("%c \n",str3[i])
	}
}