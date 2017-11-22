/**
 * Created: 2017/11/8
 * @author: Jonn
 */

package main

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ", os.Args[0], "ip:port")
		os.Exit(1)
	}
	addr := os.Args[1]
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dialhttp: ", err)
	}

	var reply struct{
		C int
	}

	var divi struct{
		Quo int
		Rem int
	}


	args := &Args{2, 2}
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		log.Fatal("call hellorpc: ", err)
	}
	fmt.Println(reply.C)

	diviValue := &Args{78, 2}
	err = client.Call("Arith.Divide", diviValue, &divi)
	if err != nil {
		log.Fatal("call hellorpc: ", err)
	}
	fmt.Println(divi)


}
