/**
 * Created: 2017/11/8
 * @author: Jonn
 */
package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
func main() {
	c = make(chan int)
	go ready("Tea", 3)
	go ready("Coffee", 6)
	fmt.Println("I'm waiting, but not too long")
	<-c
	<-c
}
