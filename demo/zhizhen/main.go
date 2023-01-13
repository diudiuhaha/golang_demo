package main

import "fmt"

func main() {
	a := 1
	fmt.Println("&取a的地址：", &a)
	//fmt.Println(*0xc00001c098)
	b := &a
	fmt.Printf("b赋值为a的地址，b的值为%#v  \n", b)
	//b赋值为a的地址，b的值为(*int)(0xc00001c098)

	//更改地址b的值
	*b = 2
	//改变内存中的值，会改变原来的值
	fmt.Printf("改变了内存值的b为：%v \n", b)
	//改变了内存值的b为：0xc00001c098
	fmt.Printf("b的值为：%#v,a的值为:%#v \n", b, a)
	//b的值为：(*int)(0xc00001c098),a的值为:2
	//所以通过更改内存中的值，会改变对地址对应的值，也就是指针对应的值
	*(&a) = 3
	fmt.Println(a) //a的值为3

}
