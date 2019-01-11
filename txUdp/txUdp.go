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
	fmt.Println("Sending to " + dst)

	// Connect to server
	conn, erConn := net.Dial("udp", dst)
	// In case of error
	if erConn != nil {
		log.Fatal("Error: net.Dial function.")
	}
	defer conn.Close()

	// Set message
	msg := []byte("UDP GO")

	// Send message
	_, erWrite := conn.Write(msg)
	if erWrite != nil{
		log.Fatal("Error: Can NOT send message.")
	}

}
