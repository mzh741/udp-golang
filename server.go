package main

import "fmt"
import "net"

func main() {
	udpaddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:11111");
	listener, _ := net.ListenUDP("udp", udpaddr)
	
	defer listener.Close() //release resource
	
	//recevier
	buffer := make([]byte, 2048)
	len, remoteAddr ,_ := listener.ReadFromUDP(buffer)
	fmt.Printf(string(buffer[:len]))
	fmt.Printf("\n")
	
	//sender
	_, _ = listener.WriteToUDP([]byte("DATA"), remoteAddr)
	
}
