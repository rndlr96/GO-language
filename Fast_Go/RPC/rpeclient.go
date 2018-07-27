package main

import (
    "fmt"
    "net/rpc"
)

type Args struct {
    Width, Height, Factor int
}

type Reply struct {
    C, D int
}

func main() {
    client, err := rpc.Dial("tcp", "127.0.0.1:6000")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer client.Close()

    args := &Args{1, 2, 3}
    reply := new(Reply)

    err = client.Call("Rectangle.RectangleScaleA", args, reply)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("reply.A : ", reply.C)
    fmt.Println("reply.B : ", reply.D)


    args.Width = 4
    args.Height = 9
    args.Factor = 2

    sumCall := client.Go("Rectangle.RectangleScaleA", args, reply, nil)
    <-sumCall.Done
    fmt.Println("reply.A : ", reply.C)
    fmt.Println("reply.B : ", reply.D)
}
