package system_grpc

import (
	"google.golang.org/grpc"
	"log"
)

func GetConn(serverUri string) (*grpc.ClientConn) {
	//arr := strings.Split(serverUri, "/")
	//service := arr[2]
	//myFun := arr[3]
	//firstCode := string([]rune(myFun)[:1])
	//myFun = strings.ToUpper(firstCode) + string([]rune(myFun)[1:])
	//if regStruct[service] != nil {
	//	m := regStruct[service]
	//	myFun := reflect.ValueOf(m).MethodByName(myFun)
	//}
	//fmt.Printf("service:::::%s", service)
	////impl.Factory =   service+="Impl"
	//return impl.Factory
	//arr := strings.Split(serverUri, "/")
	//service := arr[2]
	//myFun := arr[3]
	//firstCode := string([]rune(myFun)[:1])
	//myFun = strings.ToUpper(firstCode) + string([]rune(myFun)[1:])
	var address string
	if serverUri == "/mine/pool" {
		address = "localhost:50051"
	}
	//address = "localhost:50051"
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
	//switch service {
	//case "mine/pool":
	//	return pb.NewPoolClient(conn)
	//case "mine/miner":
	//	return pb.NewMinerClient(conn)
	//}
	//
	//
	//
	//c := pb.NewGreeterClient(conn)
	//
	//// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.Message)
}
