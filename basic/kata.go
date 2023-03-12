package main

import "fmt"

func main() {
	// キャスト
	// var f float64 = 10
	// var n int = int(f)
	// println(n)

	// // 組み込み
	// var sum int
	// sum = 5 + 6 + 1
	// ave := sum / 3
	// if ave > 4 {
	// 	println("good")
	// }

	// // 配列操作
	// ns := [...]int{10, 20, 30, 22}
	// println(ns[3])

	// スライスの初期化
	var ns1 []int
	ns1 = make([]int, 3, 10)
	println(ns1[2])

	// appendの挙動
	a := []int{10, 20}
	// [10 20] 2
	fmt.Println(a, cap(a))

	b := append(a, 30)
	a[0] = 100
	fmt.Println(b, cap(b))

	// // ポインタ
	// n, m := 10, 20
	// swap2(&n, &m)
	// println(n, m)

	for i := 1; i <= 100; i++ {
		print(i)
		if i%2 == 0 {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}

	// I/O
	scanner := bufio

}
