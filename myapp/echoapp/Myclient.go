package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "myapp/echoapp/Myecho"
)

const (
    address    = "192.168.34.13:60001"
    address2   = "192.168.34.13:60002"
)

func Mecho(){
    conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewDataClient(conn)
    stream, err := c.GetInfo(context.Background())
    if err != nil {
        fmt.Println(c,err)
    }
    waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(in.Name)
		}
    }()
    for {
        Reader := bufio.NewReader(os.Stdin)
        ss, _:= Reader.ReadString('\n')
        stream.Send(&pb.MyInfoResponse{Name:ss})
    }
    stream.CloseSend()
	<-waitc
}
func Mtime(){
    conn2, err2 := grpc.Dial(address2, grpc.WithInsecure())
    if err2 != nil {
		log.Fatalf("did not connect2: %v", err2)
	}
    defer conn2.Close()
    c2 := pb.NewTimeDataClient(conn2)
    stream2, err2 := c2.TimePush(context.Background())
    if err2 != nil {
        fmt.Println(c2,err2)
        return
    }
    for {
        in, _:= stream2.Recv()
        fmt.Println(in.Ti)
    }
       // stream2.CloseRecv()
}


func main() {
    for {
    go Mecho()
       Mtime() 
    }
}

