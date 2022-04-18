package socket

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// Create Listener in random ports
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	// starting go goroutine, access to client
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			// the method of accept clears blocking and return err
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			//call the method of close
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					// return io.EOF that close opposite connection
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Log("recived: %q", buf[:n])
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
	<-done
	//
	listener.Close()
	<-done
}
