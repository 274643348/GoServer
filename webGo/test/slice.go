package main

import "fmt"

func main() {
	ar := []byte{'a','b','c','d','e'};
	//切片是引用类型（map也是）
	arr := ar
	fmt.Printf("ar 地址：%p\n", &ar[0])
	fmt.Printf("arr 地址：%p\n", &arr[0])

	arrr1 := [10]byte{'a','b','c','d','e'};
	//基本类型是赋值
	arrr2 := arrr1
	fmt.Printf("arrr1 地址：%p\n", &arrr1[0])
	fmt.Printf("arrr2 地址：%p\n", &arrr2[0])
	fmt.Println("ar len：", len(ar))
	fmt.Println("ar cap：", cap(ar))

	ar2 := []byte{'a','b','c','d','e'};

	fmt.Printf("ar2 地址：%p\n", &ar2[0])
	fmt.Println("ar2 len：", len(ar2))
	fmt.Println("ar2 cap：", cap(ar2))

	//ar3与ar共享一片空间
	ar3 := ar[0:3]
	fmt.Printf("ar3 地址：%p\n", &ar3[0])
	fmt.Println("ar3 len：", len(ar3))
	fmt.Println("ar3 cap：", cap(ar3))

	//通过copy复制，不用一片空间
	ar4 := make([]byte, len(ar))
	copy( ar4,ar[0:3])
	fmt.Printf("ar4 地址：%p\n", &ar4[0])
	fmt.Println("ar4 len：", len(ar4))
	fmt.Println("ar4 cap：", cap(ar4))


	//ar3动态扩展后与ar不用同一片空间
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	fmt.Printf("ar 地址：%p\n", &ar[0])
	fmt.Printf("ar3 地址：%p\n", &ar3[0])
	fmt.Println("ar3 len：", len(ar3))
	fmt.Println("ar3 cap：", cap(ar3))


	var array [20]int
	slice :=array[2:4]

	fmt.Println(len(slice));
	fmt.Println(cap(slice));

	slice2 :=array[2:4:7]
	fmt.Println(len(slice2));
	fmt.Println(cap(slice2));
}
