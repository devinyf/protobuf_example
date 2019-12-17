package main

import (
	"context"
	"fmt"

	// "io"
	"log"
	"myproject/goGrpc_course/greet/greetpb"
	"net"

	// "strconv"
	// "time"

	"google.golang.org/grpc"
)

type server struct{}

// 接收客户端 cli.Greet 方法的 request
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with: ", req)
	firstName := req.GetGreeting().GetFirstName()

	// 返回给客户端的应答
	result := "Hello " + firstName
	rst := &greetpb.GreetResponse{
		Result: result,
	}

	return rst, nil
}

// // 接收客户端 cli.GreetManytimes 方法的 request
// func (*server) GreetManytimes(req *greetpb.GreetManytimesRequest, stream greetpb.GreetService_GreetManytimesServer) error {
// 	fmt.Println("GreetManytimes function was invoked with:", req)
// 	firstName := req.GetGreeting().GetFirstName()
// 	for i := 0; i < 10; i++ {
// 		result := "Hello" + firstName + "number" + strconv.Itoa(i)
// 		resp := &greetpb.GreetManytimesResponse{
// 			Result: result,
// 		}
// 		stream.Send(resp)
// 		time.Sleep(1000 * time.Millisecond)
// 	}
// 	return nil
// }

// // 接收客户端 cli.LongGreet 方法的 request
// func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
// 	fmt.Println("LongGreet function was invoked with:", stream)
// 	result := ""
// 	// 循环接收若干 stream.send() 的请求
// 	for {
// 		req, err := stream.Recv()
// 		// 全部接收完毕：一次回应全部 result
// 		if err == io.EOF {
// 			fmt.Println("finished reading client stream")
// 			// SendAndClose 函数 return 了 error
// 			return stream.SendAndClose(&greetpb.LongGreetResponse{
// 				Result: result,
// 			})
// 		}
// 		if err != nil {
// 			log.Fatalln("reading client stream error: ", err)
// 		}
// 		firstName := req.GetGreeting().GetFirstName()
// 		result += "Hello " + firstName + "! "
// 	}
// }

// func (*server) GreetEvery(stream greetpb.GreetService_GreetEveryServer) error {
// 	fmt.Println("GreetEvery function was invoked with streaming request:")
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			log.Fatalln("Error while reading client stream")
// 			return err
// 		}
// 		firstName := req.GetGreeting().GetFirstName()
// 		result := "Hello " + firstName + "! "
// 		// 发送应答
// 		if err := stream.Send(&greetpb.GreetEveryResponse{
// 			Result: result,
// 		}); err != nil {
// 			log.Fatalf("Error while sending datea to client: %v", err)
// 			return err
// 		}
// 	}
// }

func main() {
	fmt.Println("Hello I'm servers!")
	// 50051 是 gRPC 的默认端口
	lst, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//创建 server 并注册
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// 启动server
	if err := s.Serve(lst); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}
