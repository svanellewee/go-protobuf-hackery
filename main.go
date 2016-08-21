package main

import (
	"fmt"
	// "github.com/golang/protobuf/proto"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func serverCode(responder *zmq.Socket) {
	msg, _ := responder.Recv(0)
	fmt.Println("Server RX", msg)
	result := fmt.Sprintf("CHANGED %s", msg)
	responder.Send(result, 0)
}

func runServer(doSomething func(*zmq.Socket)) {
	responder, _ := zmq.NewSocket(zmq.REP)
	responder.Bind("tcp://*:5555")
	defer responder.Close()

	for {
		doSomething(responder)
	}
}

func clientCode(requester *zmq.Socket) {
	for i := 0; i< 10; i++ {
		msg := fmt.Sprintf(">>%d>>", i)
		requester.Send(msg, 0)
		retval, _ := requester.Recv(0)
		fmt.Println("Getting ", retval)
	}
}

func runClient(doSomething func(*zmq.Socket)){
	requester, _ := zmq.NewSocket(zmq.REQ)
	requester.Connect("tcp://localhost:5555")
	defer requester.Close()
	doSomething(requester)
}
/*
HOw about we create a new protobuf, and send it via REQ REP?
*/
func main() {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		runServer(serverCode)
		return
	}
	runClient(clientCode)
	
}
