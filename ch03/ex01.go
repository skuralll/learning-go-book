package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== 3.1 配列 =====")
	// 配列宣言
	var x1 = [12]int{1, 5: 4, 6, 10: 100, 15}
	for _, v := range x1 {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("")
	var x2 = [...]int{1, 2, 3}
	for _, v := range x2 {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("")
	// インデックス
	fmt.Println(x1[0])
	// 長さ
	fmt.Println(len(x2))
	fmt.Println("===== 3.2 スライス =====")
	var x3 = []int{10, 20, 5: 30} // [n]や[...]のように書かなければスライスになる
	echoSlice(x3)
	fmt.Println(len(x3))
	// append
	var x4 = []int{0}
	x4 = append(x4, 1)
	echoSlice(x4)
	// キャパシティ
	var x5 = []int{99: 0}
	fmt.Println(cap(x5))
	x5 = append(x5, 0)
	fmt.Println(cap(x5)) // スライスのサイズを増やすと多めに容量を確保する
	// make
	x6 := make([]int, 0, 10)
	x6 = append(x6, 1, 2, 3, 4)
	fmt.Printf("%d:%d\n", len(x6), cap(x6)) // 予めキャパシティを設定しているため増えない
	// スライスのスライス
	var x7 = []int{0, 1, 2, 3, 4, 5}
	y7 := x7[2:4]
	fmt.Println(y7)
	// スライスは記憶領域を共有する
	var x8 = []int{0, 1, 2, 3, 4, 5}
	y8 := x8[0:3]
	y8[0] = 100
	fmt.Println(x8) // xの中身まで変わってしまう
	// メモリを共有しないスライスのコピー
	var x9 = []int{0, 1, 2, 3, 4, 5}
	y9 := make([]int, 3)
	copy(y9, x9[0:3])
	fmt.Println(y9)
	fmt.Println("===== 3.3 文字列、rune、バイト =====")
	var s1 string = "Hello there"
	var b1 byte = s1[6]
	fmt.Println(b1)
	// スライス式を同じように使える
	fmt.Println(s1[0:3])
	// マルチバイト文字
	fmt.Println("len(☀) = ", len("☀"))
	// rune, byte はstringに変換できる
	var r1 rune = 'x'
	var b2 byte = 'y'
	fmt.Println(string(r1), string(b2))
}

func echoSlice(s []int) {
	for _, v := range s {
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
}
