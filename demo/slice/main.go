// 切片学习demo
package main

import (
	"fmt"
	"sort"
)

func main() {
	// var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// // 切片，左开右闭
	// fmt.Println(a[:])   // [0 1 2 3 4 5 6 7 8 9]
	// fmt.Println(a[2:6]) // [2 3 4 5]
	// // 长度和容量
	// fmt.Printf("a的长度为%v,容量为%v \n", len(a), cap(a))           // a的长度为10,容量为10
	// fmt.Printf("b的长度为%v,容量为%v \n", len(a[2:6]), cap(a[2:6])) // b的长度为4,容量为8

	// // 声明创建一个切片
	// a := make([]int, 4, 4)
	// fmt.Println(a) // [0 0 0 0]
	// // 扩容(在后面扩容)
	// a = append(a, 1, 2, 3, 4)
	// fmt.Println(a) // [0 0 0 0 1 2 3 4]
	// // 扩容+合并
	// // 合并
	// b := []int{5, 6, 7, 8}
	//
	// a = append(a, b...)
	// fmt.Println(a) // [0 0 0 0 1 2 3 4 5 6 7 8]

	// 需要复制的切片
	a := []int{0, 1, 2, 3, 4}
	// 创建一个空的切片b，长度和容量与a相同
	b := make([]int, len(a), cap(a))
	fmt.Println(b) // [0 0 0 0 0]
	copy(b, a)
	fmt.Println(b) // [0 1 2 3 4]
	// 修改切片
	b[0] = 10
	fmt.Println(b) // [10 1 2 3 4]
	// 删除切片的元素
	// 删除切片a下标为1和2的元素
	a = append(a[:1], a[3:]...)
	fmt.Println(a) // [0 3 4]

	// sort排序
	sort.Ints(b)
	fmt.Println(b) // [1 2 3 4 10]
	// 倒叙
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b) // [10 4 3 2 1]

}
