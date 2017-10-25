package main

import (
	"fmt"
	"net"
	"os"
	"syscall"

	"github.com/Sirupsen/logrus"
)

const (
	optUDPENCAP         = 100
	optUDPENCAPESPINUDP = 2
)

func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":4500")
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}
	defer ServerConn.Close()

	//udpconn, ok := ServerConn.(*net.UDPConn)
	//if !ok {
	//	fmt.Println("error in casting *net.Conn to *net.UDPConn!")
	//	os.Exit(1)
	//}

	//file, err := udpconn.File()
	file, err := ServerConn.File()
	if err != nil {
		fmt.Println("error in getting file for the connection!")
		os.Exit(1)
	}
	err = syscall.SetsockoptInt(int(file.Fd()), syscall.IPPROTO_UDP, optUDPENCAP, optUDPENCAPESPINUDP)
	file.Close()
	if err != nil {
		fmt.Println("error in setting priority option on socket:", err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
