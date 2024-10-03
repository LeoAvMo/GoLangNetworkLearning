package tcpconnection

import (
	"net"
	"testing"
)

// Server function
func Listener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0") // Host 127.0.0.1 accepts a tcp connection at port 0"
	// listener, err := net.Listen("tcp4", "127.0.0.1:0") // Host 127.0.0.1 accepts a tcp only for ipv4 addresses connection at port 0"
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("connection from %q", conn.RemoteAddr())
		go func(c net.Conn) {
			defer c.Close()
			// handle connection
		}(conn)
	}
}
