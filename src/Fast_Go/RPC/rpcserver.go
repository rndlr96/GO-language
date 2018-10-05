package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Rectangle int

type Args struct {
	Width, Height, Factor int // A : width, B : height
}

type Reply struct {
	C, D int // C : reply width, D : reply height
}

func (r *Rectangle) RectangleScaleA(args Args, reply *Reply) error {
	reply.C = args.Width * args.Factor
	reply.D = args.Height * args.Factor
	return nil
}

func main() {
	rpc.Register(new(Rectangle))
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()

		go rpc.ServeConn(conn)
	}
}
