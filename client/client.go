package main

import (
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	// 导入grpc包
	"google.golang.org/grpc"
	// 导入刚才我们生成的代码所在的proto包。
	pb "zhaor.com/mapper-grpc/pkg/apis/dmi/v1"
)

func UnixConnect(context.Context, string) (net.Conn, error) {
	unixAddress, err := net.ResolveUnixAddr("unix", "/tmp/a.sock")
	conn, err := net.DialUnix("unix", nil, unixAddress)
	return conn, err
}

func main() {
	// 连接grpc服务器
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	conn, err := grpc.Dial("/tmp/a.sock", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithContextDialer(UnixConnect))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := pb.NewMapperClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	_, err = c.MapperRegister(ctx, &pb.MapperRegisterRequest{
		Mapper: &pb.MapperInfo{
			Name:       "test",
			Version:    "1.0",
			ApiVersion: "1.1",
			Protocol:   "ble",
			Address:    []byte("127.0.0.1"),
			State:      "ok",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greetings: %s", "test success")

	// 调用SayHello接口，发送一条消息
	
	_, err = c.ReportDeviceStatus(ctx, &pb.ReportDeviceStatusRequest{
		DeviceName:     "device-1",
		ReportedDevice: &pb.DeviceStatus{
			Twins: []*pb.Twin{&pb.Twin{
				PropertyName: "temperature",
				Desired:      &pb.TwinProperty{
					Value:    "27",
					Metadata: nil,
				},
				Reported:     &pb.TwinProperty{
					Value:    "30",
					Metadata: nil,
				},
			}},
			State: "online",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greetings: %s", "test success")
}