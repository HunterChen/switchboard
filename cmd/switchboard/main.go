package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	. "github.com/pivotal-cf-experimental/switchboard"
)

var (
	pidfile = flag.String("pidfile", "", "The location for the pidfile")
	port    = flag.Uint("port", 3306, "Port to listen on")

	backendIp          = flag.String("backendIp", "", "IP address of backend")
	backendPort        = flag.Uint("backendPort", 3306, "Port of backend")
	healthcheckPort    = flag.Uint("healthcheckPort", 9200, "Port for healthcheck endpoints")
	healthcheckTimeout = flag.Duration("healthcheckTimeout", 5*time.Second, "Timeout for healthcheck")
)

func acceptClientConnection(l net.Listener) net.Conn {
	clientConn, err := l.Accept()
	if err != nil {
		log.Fatal("Error accepting client connection: %v", err)
	}
	return clientConn
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error listening on port %d: %v\n", *port, err.Error()))
	}
	defer l.Close()

	err = ioutil.WriteFile(*pidfile, []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot write pid to file: %s", *pidfile))
	}

	fmt.Printf("Proxy started on port %d\n", *port)
	fmt.Printf("Backend ipAddress: %s\n", *backendIp)
	fmt.Printf("Backend port: %d\n", *port)
	fmt.Printf("Healthcheck port: %d\n", *healthcheckPort)

	backend := NewBackend("backend1", *backendIp, *backendPort)

	healthcheck := NewHttpHealthCheck(*backendIp, *healthcheckPort, *healthcheckTimeout)
	healthcheck.Start(backend.RemoveAndCloseAllBridges)

	for {
		clientConn := acceptClientConnection(l)
		defer clientConn.Close()

		backendConn, err := backend.Dial()
		if err != nil {
			log.Fatal("Error connection to backend: %s", err.Error())
		}
		defer backendConn.Close()

		bridge := NewConnectionBridge(clientConn, backendConn)
		backend.AddBridge(bridge)

		go func() {
			bridge.Connect()
			backend.RemoveBridge(bridge)
		}()
	}
}