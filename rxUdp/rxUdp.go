package main 

import(
	"fmt"
	"log"
	"os"
	"net"
)

func main (){

	// Check Arguments
	if len(os.Args) != 3{
		argErMsg := string("Usage: " + os.Args[0] + "[IP] [PORT]")
		log.Fatal(argErMsg)
	}

	// Get parameters from arguments
	ip := os.Args[1]
	port := os.Args[2]
	dst := string(ip + ":" + port)
	fmt.Println("Start Listening via " + dst)

	// Listen
	conn, erConn := net.ListenPacket("udp", dst)
	// In case of error
	if erConn != nil {
		log.Fatal("Error: Can NOT listen.")
	}
	defer conn.Close()

	// Create buffer
	msgBuf := make([]byte, 1000)

	// Receive message
	for{
		_, srcAddr, erRead :=conn.ReadFrom(msgBuf)
		// In case of error
		if erRead != nil{
			log.Fatal("Error: Can NOT receive.")
		}
		fmt.Println("Receive message from : " + srcAddr.String())
		fmt.Println(" \"" + string(msgBuf[:]) + "\"")
	}

}
