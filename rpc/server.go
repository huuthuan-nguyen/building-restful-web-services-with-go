package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {

}

type TimeServer int64
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// crate new RPC server
	timeserver := new(TimeServer)
	// register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// Listen for requests on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}