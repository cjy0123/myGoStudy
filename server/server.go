/**
 * Created: 2017/11/8
 * @author: Jonn
 */

package server

import (
	"errors"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
