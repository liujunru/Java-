package main
import(
	"fmt"
) 
type Goods struct{
	Name string
	Price float64
}

type Brand struct{
	Name string 
	Address string 
}

type TV struct{
	Goods
	Brand
}

type TV2 struct{
	*Goods
	*Brand
}

func main(){
	tv := TV{Goods{"电视机01",5000.9}, Brand{"海尔","山东"}}
	
	tv2 := TV{
		Goods{
			Name : "电视机02",
			Price : 3000.01,
		},
		Brand{
			Name : "夏普",
			Address : "北京",
		},
	}

fmt.Println("TV", tv)//TV {{电视机01 5000.9} {海尔 山东}}
fmt.Println("TV2", tv2)//TV2 {{电视机02 3000.01} {夏普 北京}}

tv3 := TV2{&Goods{"电视机03", 600.19}, &Brand{"创维","河南"}}

tv4 := TV2{
	&Goods{
		Name : "电视机04",
		Price : 8000.34,
	},
	&Brand{
		Name : "长虹",
		Address : "成都",
	},
}

fmt.Println("TV3", *tv3.Goods,*tv3.Brand)//TV3 {电视机03 600.19} {创维 河南}
fmt.Println("TV4", *tv4.Goods,*tv4.Brand)//TV4 {电视机04 8000.34} {长虹 成都}
}