// hello.go
package main

import "fmt"

func main() {
	// var x int = 1
	// xx := float64(x)
	// fmt.Printf("%T %v %f\n", xx, xx, xx)

	// var y float64 = 1.2
	// yy := int(y)
	// fmt.Printf("%T %v %d\n", yy, yy, yy)

	// var s string = "14"
	// i, _ := strconv.Atoi(s)
	// fmt.Printf("%T %v", i, i)

	// var a [2]int
	// a[0] = 100
	// a[1] = 200
	// fmt.Println(a)

	// var b = []int{100, 200}
	// b = append(b, 300)
	// fmt.Println(b)

	// n := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(n)
	// fmt.Println(n[2])
	// fmt.Println(n[2:4])
	// fmt.Println(n[:2])
	// fmt.Println(n[2:])
	// fmt.Println(n[:])

	// n[2] = 100
	// fmt.Println(n)

	// var board = [][]int{
	// 	[]int{0, 1, 2},
	// 	[]int{3, 4, 5},
	// 	[]int{6, 7, 8},
	// }

	// fmt.Println(board)

	// n = append(n, 100, 200, 300, 400)
	// fmt.Println(n)

	// n := make([]int, 3, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 0, 0)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 1)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// a := make([]int, 3)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	// b := make([]int, 0)
	// var c []int
	// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	// c = make([]int, 5)
	// // c = make([]int, 0, 5)
	// for i := 0; i < 5; i++ {
	// 	c = append(c, i)
	// 	fmt.Println(c)
	// }
	// fmt.Println(c)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["peach"] = 500
	fmt.Println(m)

	fmt.Println(m["noting"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["orange"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	var s []int
	if s == nil {
		fmt.Println("nil")
	}
}
