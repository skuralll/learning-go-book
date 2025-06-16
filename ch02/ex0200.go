package main

import (
	"fmt"
	"math"
)

// x := 1 // error:「:=」は関数内でしか使えない

func main() {
	fmt.Println("===== 2.1 基本型 =====")
	fmt.Println("===== 2.1.1 ゼロ値 =====")
	{
		var x int
		var y float32
		var z string

		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		fmt.Printf("z: 「%s」\n", z)
	}
	fmt.Println("===== 2.1.2 リテラル =====")
	fmt.Println(1_234)
	fmt.Println(1_234.1_234)
	{
		fmt.Printf("%c\n", 'a')          // 1文字Unicode
		fmt.Printf("%c\n", '\141')       // 8bit8進数
		fmt.Printf("%c\n", '\x61')       // 8bit16進数
		fmt.Printf("%c\n", '\u0061')     // 16bit16進数
		fmt.Printf("%c\n", '\U00000061') // 32bitUnicode
	}
	fmt.Println("string literal")
	fmt.Println("===== 2.1.3 論理型bool =====")
	var flag bool // ゼロ値はfalse
	fmt.Println(flag)
	flag = true
	fmt.Println(flag)
	fmt.Println("===== 2.1.4 数値型 =====")
	fmt.Printf("int8 min:%d max:%d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("uint8 min:%d max:%d\n", 0, math.MaxUint8)
	fmt.Println("===== 2.1.5 文字列型とrune型 =====")
	fmt.Println("string1" + " " + "string2")

	fmt.Println("===== 2.2 変数宣言 =====")
	// var x int = 10 // 型指定
	// var x = 10 // 型推論
	// var x int // ゼロ値を代入
	// var x, y int = 10, 20 // 複数宣言, 多重代入
	// var x, y int // ゼロ値が代入される
	// 宣言リスト
	// var (
	// 	x    int
	// 	y        = 20
	// 	z    int = 30
	// 	d, e     = 40, "hello"
	// 	f, g string
	// )
	// x := 10 // varの省略と型推論
	// x := 10
	// x, y := 30, "hello" // 既存の変数にも値を代入できる
}
