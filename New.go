package sock

import (
	"fmt"
	"net"
	"os"
)

type (
	// Map is alias for map[string]interface{}
	Map map[string]interface{}
)

// New makes new Channel with "net/http".Request
func New(addr string) *Channel {
	channel := newChannel()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go channel.handler(conn)
		}
	}()
	return channel
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
