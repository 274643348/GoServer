package main

import "fmt"

func main() {
	ar := []byte{'a','b','c','d','e'};

	fmt.Printf("ar 地址：%p\n", &ar[0])
	fmt.Println("ar len：", len(ar))
	fmt.Println("ar cap：", cap(ar))

	ar2 := []byte{'a','b','c','d','e'};

	fmt.Printf("ar2 地址：%p\n", &ar2[0])
	fmt.Println("ar2 len：", len(ar2))
	fmt.Println("ar2 cap：", cap(ar2))

	//ar3与ar共享一片控件
	ar3 := ar[0:3]
	fmt.Printf("ar3 地址：%p\n", &ar3[0])
	fmt.Println("ar3 len：", len(ar3))
	fmt.Println("ar3 cap：", cap(ar3))

	//通过copy复制
	ar4 := make([]byte, len(ar))
	copy( ar4,ar[0:3])
	fmt.Printf("ar4 地址：%p\n", &ar4[0])
	fmt.Println("ar4 len：", len(ar4))
	fmt.Println("ar4 cap：", cap(ar4))


	//ar3与ar共享一片控件
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	ar3= append(ar3, 1)
	fmt.Printf("ar3 地址：%p\n", &ar3[0])
	fmt.Println("ar3 len：", len(ar3))
	fmt.Println("ar3 cap：", cap(ar3))


	fmt.Printf("ar 地址：%p\n", &ar[0])
	fmt.Printf("ar3 地址：%p\n", &ar3[0])
}
