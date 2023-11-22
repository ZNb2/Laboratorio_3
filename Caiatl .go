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

func Borrar_Base(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.BorrarBase(context.Background(), &pb.Message{Body: mensaje})

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
func Actualizar_Valor(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.ActualizarValor(context.Background(), &pb.Message{Body: mensaje})

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

func Renombrar_Base(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.RenombrarBase(context.Background(), &pb.Message{Body: mensaje})

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

func Agregar_Base(Servidor string, mensaje string){
	
	conn, err := grpc.Dial(Servidor, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	for {
		_, err := c.AgregarBase(context.Background(), &pb.Message{Body: mensaje})

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
	server_name := "Informante Caiatl"
	fmt.Println("Starting " + server_name + " . . .\n")

	fmt.Print("\nOpciones disponibles:\n
		\t1 = Agregar Base\n
		\t2 = Renombrar Base\n
		\t3 = Actualizar Valor\n
		\t4 = Borrar Base\n"
	)

	for {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese una opciones: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	fmt.Print("Ingrese un comando: ")
	msj, _ := reader.ReadString('\n')
	msj = strings.TrimSuffix(input, "\n")


	if input == "1" {
		Agregar_Base(To_broker(server_broker, msj), msj)
	} else if == "2" {
		Renombrar_Base(To_broker(server_broker, msj), msj)
	} else if == "3" {
		Actualizar_Valor(To_broker(server_broker, msj), msj)
	} else if == "4" {
		Borrar_Base(To_broker(server_broker, msj), msj)
	} else{
		fmt.Println("\nComando no reconocido!\n")
	}
	}
}