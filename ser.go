/**
 * Created: 2017/11/16
 * @author: Jonn
 */

package main

import (
	"myGoStudy/server"
	"net/rpc"
	"fmt"
	"net"
	"os"
)

func main() {
	arith := new(server.Arith)
	if err := rpc.Register(arith); err != nil {
		fmt.Print(err)
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:1235")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		fmt.Println("hello")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
