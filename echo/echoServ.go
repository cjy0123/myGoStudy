/**
 * Created: 2017/11/20
 * @author: Jonn
 */

package echo

import (
	"net/rpc"
	"net"
)

type EchoServer bool

func (s *EchoServer) Echo(req *string, res *string) error {
	res = req
	return nil
}

func main() {
	echo := new(EchoServer)
	if err:=rpc.Register(echo);err!=nil{
		print(err)
	}
	var listener *net.TCPListener
	if tcpAddr, err:=net.ListenTCP("tcp",addr)

}