package main

func main() {
	// キャスト
	var f float64 = 10
	var n int = int(f)
	println(n)

	// 組み込み
	var sum int
	sum = 5 + 6 + 1
	ave := sum / 3
	if ave > 4 {
		println("good")
	}

	// 配列操作
	ns := [...]int{10, 20, 30, 22}
	println(ns[3])

	// スライスの初期化
	var ns1 []int
	ns1 = make([]int, 3, 10)
	println(ns1[2])

}
