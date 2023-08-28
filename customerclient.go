package main

import (
	"context"
	"fmt"
	pb "grpc/customer"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed1 %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomerserviceClient(conn)
	customer := &pb.CustomerDetails{
		Name:    "sini",
		Balance: 0,
		BankID:  0,
	} 
	// response,err:=client1.SayHello(context.Background(),&hw.HelloRequest{Name:name,Age:int32(age)})
	//   addres,err:=client.FetchDetails(context.Background(),pb.CustomerDetails)
	// 	if err!=nil{        log.Fatal("failed2 %v",err)
	// 	}
	// 	 fmt.Printf("Response of add :%s\n",addres.Id)
	taskres, err := client.AddAccount(context.Background(),customer)
	if err != nil {
		log.Fatal("failed2 %v", err)
	}
	fmt.Printf("Response of get :",taskres)
	// for _, task := range taskres.Tasks {
	// 	fmt.Printf("ID:%s,Title:%s,Completed:%v\n", customer.AccID, cus, task.Completed)
	// }
}
