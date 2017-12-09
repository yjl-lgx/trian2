package main

import (
	// "os"
	// "fmt"
    "io"
	"log"
    "net"
    "time" 
    _ "myapp/routers"
	"github.com/astaxie/beego"
    "google.golang.org/grpc"
	pb "myapp/echoapp/Myecho"
    "google.golang.org/grpc/reflection"
)

const (
    port = ":60001"
    ppt = ":60002"
)

type Server struct{}

type MyTime struct{}

type Time2 struct{
    te string
}

//回显函数检查流中是否有数据，若有则发送回客服端
func (this *Server) GetInfo(stream pb.Data_GetInfoServer)(error){
    for{
        in, err := stream.Recv()
        if err == io.EOF {
			return nil
		}
        if err != nil{ 
           return err   
        }
        stream.Send(&pb.MyInfoResponse{Name:in.Name})     
    }
}
func Svers(){
    //开启服务
	lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    //新建服务
    pb.RegisterDataServer(s, &Server{})	
    //注册到关联
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
    }
}

func (this *MyTime) TimePush(stream pb.TimeData_TimePushServer)(error){
    ticker := time.NewTicker(20 * time.Second)
        for {
            time := <-ticker.C
            timrout := "01__02-2006 3.04.05 PM"
            var now = time.Format(timrout)
            var testtime pb.TimeResponse
            testtime.Ti = now 
            //stream.Send(&pb.TimeResponse{Ti:now})
            stream.Send(&testtime)
          }
          return nil
 }

func Mytime(){
    //开启服务
	lis, err := net.Listen("tcp", ppt)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    t := grpc.NewServer()
    //新建服务
    pb.RegisterTimeDataServer(t, &MyTime{})	
    //注册到关联
    reflection.Register(t)
    if err := t.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    go Svers()     
    go Mytime()
	beego.Run()
}









