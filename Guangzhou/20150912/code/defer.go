package main

import (
	"fmt"
)

func c() (i int) {
	i = 0
	defer func() {
		i++ // 闭包引用 i 指针。
	}()
	return
}

func c1() int {
	i := 0
	defer func() {
		i++
	}()
	return i // 返回值存储空间由调用栈准备，也就是说返回值相当于另外一个 "外部变量"。
	// return 先将右值保存到返回值变量，然后才 call defer ...
}

func c2() int {
	i := 0
	defer func(i int) {
		i++
	}(i)
	return i // 同上，写返回值在 call defer 之前。
}

func c3() (i int) {
	i = 0
	defer func(i int) {
		i++ // 和 c3 返回值不是同一个变量。
	}(i) // i 值拷贝
	return
}

func main() {
	fmt.Println(c())  // 1
	fmt.Println(c1()) // 0  为什么c和c1会出现不同结果
	fmt.Println(c2()) // 0
	fmt.Println(c3()) // 0
}
