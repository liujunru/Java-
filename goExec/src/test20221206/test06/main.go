package main
import(
	"fmt"
	"time"
) 
func main(){
	//获取当前时间
	now := time.Now();
	fmt.Printf("now=%v now type=%T\n", now, now)
	//now=2022-12-06 19:21:08.279572 +0800 CST m=+0.001794201 now type=time.Time

	//通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	// 年=2022
	// 月=December
	// 月=12
	// 日=6
	// 时=19
	// 分=25
	// 秒=58

	//格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	//当前年月日 2022-12-6 19:28:54

	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	//2022-12-06 19:30:52
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	//2022-12-06
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	//19:30:52
}