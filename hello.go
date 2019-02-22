// hello.go
package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
func init() {
	fmt.Println("init")
}
*/

func bazz() {
	fmt.Println("bazz")
}

func showTime() {
	fmt.Println(time.Now())
}

func main() {
	bazz()
	showTime()
	fmt.Println("Hello world!!")
	fmt.Println(user.Current())
}
