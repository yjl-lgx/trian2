package main

import (
	//"github.com/beego/bee/cmd/commands/server"
    //"log"
    //"net"
   // "io"
    //"golang.org/x/net/context"
    //"google.golang.org/grpc"
	pb "myapp/echoapp/Myecho"
    //"google.golang.org/grpc/reflection"
   // "fmt"
)

const (
    PORT = ":50001"
)

type Server struct{}

func (this *Server) GetInfo(stream pb.Data_GetInfoServer)(error){
  /*  //in, err := stream.Recv()
  //  if err != nil {
       // err := stream.Send(in)
    }else{
        return nil 
    }
    //return err*/
}
/*func main() {
    lis, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterDataServer(s, &Server{})
    s.Serve(lis)
    
}*/