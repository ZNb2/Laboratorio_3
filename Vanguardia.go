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


func To_broker(Servidor string, mensaje string) (response string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		response, err := c.ReceiveFromInformant(context.Background(), &pb.Message{Body: mensaje})
		if err != nil {
			log.Println(Servidor, "not responding")
			log.Println("Trying again in 10 seconds . . .")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	return
}


func main() {

	server_broker := "dist105.inf.santiago.usm.cl:50051"
	server_name := "Vanguardia"
	fmt.Println("Starting " + server_name + " . . .\n")

	fmt.Print("\nOpciones disponibles:\n
		\tGetSoldados\n"
	)

	for {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese un comando: ")
	msj, _ := reader.ReadString('\n')
	msj = strings.TrimSuffix(input, "\n")
	response := To_broker(server_broker, msj)
	log.Print("Soldados: " + response)
	}
}