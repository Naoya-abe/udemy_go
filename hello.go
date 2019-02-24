// hello.go
package main

// func add(x, y int) (int, int) {
// 	// fmt.Println(x + y)
// 	return x - y, x + y

// }

// func cal(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func convert(price int) float64 {
// 	return float64(price)
// }

// func incrementGenerator() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func circleArea(pi float64) func(radius float64) float64 {
// 	return func(radius float64) float64 {
// 		return pi * radius * radius
// 	}
// }

// func foo(params ...int) {
// 	fmt.Println(len(params), params)
// 	for _, param := range params {
// 		fmt.Println(param)
// 	}
// }

/*
func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}
*/

/*
func getOsName() string {
	return "mac"
}
*/

/*
func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}
*/

/*
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multilogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multilogFile)
}
*/

/*
func thirdPartyConnectDB() {
	panic("Unable to connect database")
}
func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
*/
/*
func one(x *int) {
	*x = 1
}
*/

type Vertex struct {
	X, Y int
	S    string
}

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

	// m := map[string]int{"apple": 100, "banana": 200}
	// fmt.Println(m)
	// fmt.Println(m["apple"])
	// m["banana"] = 300
	// fmt.Println(m)
	// m["peach"] = 500
	// fmt.Println(m)

	// fmt.Println(m["noting"])

	// v, ok := m["apple"]
	// fmt.Println(v, ok)

	// v2, ok2 := m["orange"]
	// fmt.Println(v2, ok2)

	// m2 := make(map[string]int)
	// m2["pc"] = 5000
	// fmt.Println(m2)

	// var s []int
	// if s == nil {
	// 	fmt.Println("nil")
	// }

	// b := []byte{72, 73}
	// fmt.Println(b)
	// fmt.Println(string(b))

	// c := []byte("HI")
	// fmt.Println(c)
	// fmt.Println(string(c))

	// r1, r2 := add(10, 20)
	// fmt.Println(r1, r2)

	// r3 := cal(100, 2)
	// fmt.Println(r3)

	// f := func(x int) {
	// 	fmt.Println("inner func", x)
	// }
	// f(1)

	// func(x int) {
	// 	fmt.Println("inner func", x)
	// }(1)

	// f1 := convert(20)
	// fmt.Printf("type=%T, value=%v", f1, f1)

	// x := 0
	// increment := func() int {
	// 	x++
	// 	return x
	// }
	// fmt.Println(increment())
	// fmt.Println(increment())
	// fmt.Println(increment())

	// counter := incrementGenerator()
	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())
	// fmt.Println(counter())

	// c1 := circleArea(3.14)
	// fmt.Println(c1(2))

	// c2 := circleArea(3)
	// fmt.Println(c2(2))

	// foo()
	// foo(10, 20)
	// foo(10, 20, 30)

	// s := []int{1, 2, 3}
	// foo(s...)

	// // Q1
	// f := 1.11
	// fmt.Println(int(f))

	// // Q2
	// // 5
	// // 6

	// // Q3
	// m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	// fmt.Printf("%T %v\n", m, m)
	// fmt.Println(m["Mike"])

	// num := 6
	// if num%2 == 0 {
	// 	fmt.Println("by 2")
	// } else if num%3 == 0 {
	// 	fmt.Println("by 3")
	// } else {
	// 	fmt.Println("else")
	// }

	/*
		x, y := 10, 11
		if x == 10 && y == 10 {
			fmt.Println("&&")
		} else if x == 10 || y == 10 {
			fmt.Println("||")
		}
	*/

	/*
		result := by2(10)
		if result == "ok" {
			fmt.Println("great")
		}

		if result2 := by2(10); result2 == "ok" {
			fmt.Println("great2")
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				fmt.Println("continue")
				continue
			}

			if i > 5 {
				fmt.Println("break")
				break
			}
			fmt.Println(i)
		}
		sum := 1
		for sum < 10 {
			sum += sum
			fmt.Println(sum)
		}
		fmt.Println(sum)
	*/

	/*
		l := []string{"python", "go", "java"}

		for i := 0; i < len(l); i++ {
			fmt.Println(i, l[i])
		}

		for i, v := range l {
			fmt.Println(i, v)
		}

		for _, v := range l {
			fmt.Println(v)
		}

		m := map[string]int{"apple": 100, "banana": 200}

		for k, v := range m {
			fmt.Println(k, v)
		}

		for k := range m {
			fmt.Println(k)
		}

		for _, v := range m {
			fmt.Println(v)
		}
	*/

	/*
		// os := getOsName()
		switch os := getOsName(); os {
		case "mac":
			fmt.Println("Mac!!")
		case "windows":
			fmt.Println("Windows!!")
		default:
			fmt.Println("default!")
		}

		t := time.Now()
		fmt.Println(t.Hour())
		switch {
		case t.Hour() < 12:
			fmt.Println("morning")
		case t.Hour() > 12:
			fmt.Println("afternoon")
		}
	*/
	/*
		defer fmt.Println("world")

		foo()

		fmt.Println("hello")
	*/
	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/
	/*
		file, _ := os.Open("./hello.go")
		defer file.Close()
		data := make([]byte, 100)
		file.Read(data)
		fmt.Println(data)
	*/

	/*
		LoggingSettings("test.log")
		_, err := os.Open("fafavae")
		if err != nil {
			log.Fatalln("Exit", err)
		}

		log.Println("logging!")
		log.Printf("%T %v", "test", "test")
		log.Fatalf("%T %v", "test", "test")
		log.Fatalln("error!")

		fmt.Println("ok!")
	*/

	/*
		file, err := os.Open("./hello.go")
		if err != nil {
			log.Fatalln("Error!")
		}
		defer file.Close()
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatalln("Error")
		}
		fmt.Println(count, string(data))
	*/

	/*
		save()
		fmt.Println("OK?")
	*/

	/*
		l := []int{100, 300, 23, 11, 23, 2, 3, 4, 3}
		var min int
		fmt.Println(min)
		for i, num := range l {
			if i == 0 {
				min = num
				continue
			}
			if min >= num {
				min = num
			}
		}
		fmt.Println(min)

		m := map[string]int{
			"apple":   200,
			"banana":  300,
			"grapes":  150,
			"orange":  80,
			"papaiya": 500,
			"kiwi":    90,
		}

		sum := 0

		for _, v := range m {
			sum += v
		}

		fmt.Println(sum)
	*/

	/*
		l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
		var min int
		for i, num := range l {
			if i == 0 {
				min = num
			}
			if min >= num {
				min = num
			}
		}
		fmt.Println(min)

		m := map[string]int{
			"apple":   200,
			"banana":  100,
			"peach":   300,
			"papaiya": 500,
			"orange":  80,
			"kiwi":    90,
		}
		sum := 0
		for _, v := range m {
			sum += v
		}
		fmt.Println(sum)
	*/

	/*
		var n int = 100
		fmt.Println(n)

		var p *int = &n
		fmt.Println(p)
		fmt.Println(*p)
	*/
	/*
		var n int = 100
		one(&n)
		fmt.Println(n)
	*/

	/*
		var p *int = new(int)
		fmt.Println(*p)
		*p++
		fmt.Println(*p)

		var p2 *int
		fmt.Println(p2)
	*/

	/*
		v := Vertex{X: 1, Y: 2}

		fmt.Println(v)
		fmt.Println(v.X, v.Y)
		v.X = 100
		fmt.Println(v.X, v.Y)

		v2 := Vertex{X: 1}
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3)

		v4 := Vertex{}
		fmt.Println(v4)

		var v5 Vertex
		fmt.Println(v5)

		v6 := new(Vertex)
		fmt.Printf("%T %v\n", v6, v6)

		v7 := &Vertex{}
		fmt.Printf("%T %v\n", v7, v7)
	*/
}
