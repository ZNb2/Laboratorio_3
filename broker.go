package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo27-Laboratorio-3/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var fulcrumServers = []string{
		"dist105.inf.santiago.usm.cl:50053",
		"dist106.inf.santiago.usm.cl:50053",
		"dist107.inf.santiago.usm.cl:50053",
}


type brokerLunaServer struct{}


func (s *Server)ReceiveFromInformant(ctx context.Context, in *pb.Message)(*pb.Message, error){
	
	if strings.Contains(in.Body, "inconsistencia"){
		//Enviar mensaje a Fulcrum y resolver inconsistencia
		return &pb.Message{Body: "Reparada"}, nil
	} else {
		return &pb.Message{Body: Fulcrum[rand.Intn(3)]}, nil
	}
}


func Get_Soldados(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.GetSoldados(context.Background(), &pb.Message{Body: mensaje})

		// Revisar y registrar consistencia

		if err != nil {
			log.Println(Servidor, "not responding")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
}

func (s *Server)ReceiveFromVanguardia(ctx context.Context, in *pb.Message)(*pb.Message, error){
	
	if strings.Contains(in.Body, "inconsistencia"){
		//Enviar mensaje a Fulcrum y resolver inconsistencia
		return &pb.Message{Body: "Reparada"}, nil
	} else {
		//Enviar mensaje de la vanguardia a fulcrum
		response := Get_Soldados(Fulcrum[rand.Intn(3)], in.Body)
		return &pb.Message{Body: response}, nil
	}
}


func main() {

	rand.Seed(time.Now().UnixNano())
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	brokerLuna := &brokerLunaServer{}
	fulcrum := &fulcrumServer{}

	grpcServer := grpc.NewServer()
	pb.RegisterBrokerLunaServer(grpcServer, brokerLuna)
	pb.RegisterFulcrumServerServer(grpcServer, fulcrum)

	log.Println("gRPC server running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
