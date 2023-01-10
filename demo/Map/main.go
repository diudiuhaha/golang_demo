package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	//创建一个map
	var m = make(map[string]string)
	fmt.Println(m) // 空值为：map[]
	//赋值方式1
	m["a"] = "b"
	fmt.Println(m)
	//赋值方式2
	m1 := map[string]string{"one": "1", "two": "2"}
	fmt.Println(m1) //map[one:1 two:2]

	//循环遍历
	for a, b := range m1 {
		fmt.Println(a, b)
	}

	//判断键值对是否存在
	_, statue := m1["three"] //返回值和状态
	if statue {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

	//删除键值对delete(map,key)
	m1["end"] = "E"
	delete(m1, "one")
	fmt.Println(m1) //map[end:E two:2]

	//元素类型为map的切片
	//s := make([]string, 5, 5) // string类型替换为map类型
	s := make([]map[string]int, 5, 5)
	log.Println(s) // 切片空值为：[map[] map[] map[]]
	//赋值
	key := [5]string{"a", "b", "c", "d", "e"}

	for i := 0; i < 5; i++ {

		v := map[string]int{
			key[i]:    i + 1,
			"default": 0,
		}
		s[i] = v
	}
	log.Println(s)

	//再循环遍历去读这个切片
	for a, b := range s {
		log.Println(a, "-", b)
	}

	//值为切片的map
	var stumap = make(map[string][]string)
	//赋值
	stumap["tom"] = []string{"a", "b", "c"}
	log.Println(stumap) //map[tom:[a b c]]

	//安全map
	//只需要声明，不需要make
	var m4 sync.Map //不需要加，无法赋初值
	log.Println(m4) //{{0 0} {<nil>} map[] 0}

	//赋值
	m4.Store("a", 1)
	m4.Store("b", "b")
	m4.Store("c", true)

	//输出为：{{map[] true}} map[a:0xc0000a4000 b:0xc0000a4008 c:0xc0000a4010] 0}
	//读取值
	tmp, bool1 := m4.Load("a")   //返回值和布尔值
	log.Println(tmp, "-", bool1) //1 - true

	_, bool1 = m4.Load("d") //这里是= ，后续研究
	if bool1 {
		log.Println("d存在")
	} else {
		log.Println("d不存在")
	}

}
