package main

import (
	"fmt"
	"net"
)

func main() {
	localudpaddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:11112");
	remoteudpaddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:11111");
	dial, _ := net.DialUDP("udp", localudpaddr, remoteudpaddr)
	
	defer dial.Close() //release resource
	
	//sender
	Data := []byte("DATA")
	_, _ = dial.Write(Data)
	
	//recevier
	buffer := make([]byte, 2048)
	len, _, _ := dial.ReadFromUDP(buffer)
	fmt.Printf(string(buffer[:len]))
	fmt.Printf("\n")
}
