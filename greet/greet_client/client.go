package main

import (
	"context"
	"fmt"

	// "io"
	"log"
	"myproject/goGrpc_course/greet/greetpb"

	// "time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm a client")
	//创建客户端拨号连接 //WithInsecure: 不使用ssl
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect: ", err)
	}
	defer cc.Close()

	//创建 client API 实例 (interface{})
	cli := greetpb.NewGreetServiceClient(cc)

	doUnary(cli)
	// doServerStream(cli)
	// doClientStream(cli)
	// doBiDiStream(cli)
}

func doUnary(cli greetpb.GreetServiceClient) {
	fmt.Println("starting do Unary RPC...")
	// request 内容(填写结构体)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Squall",
			LastName:  "Dong",
		},
	}
	//向server 的 Greet 发送seq, 并返回resp(result)
	resp, err := cli.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("Greet is failed", err)
	}
	log.Printf("Response form Greet:", resp.Result)

}

// func doServerStream(cli greetpb.GreetServiceClient) {
// 	fmt.Println("starting doServerStream RPC...")
// 	rst := &greetpb.GreetManytimesRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "Squall",
// 			LastName:  "dong",
// 		},
// 	}
// 	// 向服务器端的 GreetManytimes 方法传递数据
// 	respStream, err := cli.GreetManytimes(context.Background(), rst)
// 	if err != nil {
// 		log.Fatalln("Failed with calling GreetManytimes RPC: ", err)
// 	}
// 	for {
// 		msg, err := respStream.Recv()
// 		if err == io.EOF {
// 			fmt.Println("reached the end of the stream !")
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalln("failed while recv message: ", err)
// 		}
// 		fmt.Println("Response from GreetManytimes: ", msg.GetResult())
// 	}
// }

// func doClientStream(cli greetpb.GreetServiceClient) {
// 	fmt.Println("starting doClientStream RPC...")
// 	// 向服务器端的 LongGreet 方法传递数据
// 	stream, err := cli.LongGreet(context.Background())
// 	if err != nil {
// 		log.Fatalln("send client stream failed : ", err)
// 	}
// 	// 创建需要发送的 stream 列表
// 	requests := []*greetpb.LongGreetRequest{
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Squall",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Evin",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Emily",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Kimbarly",
// 			},
// 		},
// 		&greetpb.LongGreetRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Fanny",
// 			},
// 		},
// 	}

// 	// 循环发送
// 	for _, req := range requests {
// 		fmt.Printf("sending request: %v\n", req)
// 		stream.Send(req)
// 		time.Sleep(1000 * time.Millisecond)
// 	}
// 	// 发送完毕，关闭steam并接收相应
// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalln("error while receiving response form LongGreet: ", err)
// 	}
// 	// 查看服务器的相应
// 	fmt.Println("LogGreeting response:", resp.GetResult())
// 	// resp.GetResult()

// }

// func doBiDiStream(cli greetpb.GreetServiceClient) {
// 	fmt.Println("starting doBiDiStream RPC")
// 	// create a stream by invoking the client
// 	stream, err := cli.GreetEvery(context.Background())
// 	if err != nil {
// 		log.Fatalln("Error while send client stream ", err)
// 	}
// 	// request := []*greetpb.GreetEveryRequest{
// 	// 	&greetpb.GreetEveryRequest{
// 	// 		Greeting: greetpb.Greeting{
// 	// 			FirstName: "Squall",
// 	// 		}
// 	// 	}
// 	// }
// 	// stream.Send(&greetpb.GreetEveryRequest)

// 	requests := []*greetpb.GreetEveryRequest{
// 		&greetpb.GreetEveryRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Squall",
// 			},
// 		},
// 		&greetpb.GreetEveryRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Evin",
// 			},
// 		},
// 		&greetpb.GreetEveryRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Emily",
// 			},
// 		},
// 		&greetpb.GreetEveryRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Kimbarly",
// 			},
// 		},
// 		&greetpb.GreetEveryRequest{
// 			Greeting: &greetpb.Greeting{
// 				FirstName: "Fanny",
// 			},
// 		},
// 	}

// 	waitc := make(chan struct{})

// 	// send a bunch of message to the client(go routine)
// 	go func() {
// 		for _, req := range requests {
// 			fmt.Printf("sending message: %v\n", req)
// 			stream.Send(req)
// 			time.Sleep(1000 * time.Millisecond)
// 		}
// 		stream.CloseSend()
// 	}()

// 	// receive a bunch of message from the client (go routine)
// 	go func() {
// 		for {
// 			res, err := stream.Recv()
// 			if err == io.EOF {
// 				break
// 			}
// 			if err != nil {
// 				log.Fatalln("Error while recving: ", err)
// 				break
// 			}
// 			fmt.Printf("Received: %v\n", res.GetResult())
// 		}
// 		close(waitc)
// 	}()

// 	// block until everything is done
// 	<-waitc
// }
