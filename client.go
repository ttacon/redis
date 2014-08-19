package redis

import (
	"bufio"
	"net"
)

type Client struct {
	conn   net.Conn
	reader bufio.Reader
}

func NewClient(address string) (*Client, error) {
	// TODO(ttacon): support unix sockets
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	// TODO(ttacon): make size configurable
	return &Client{
		conn:   conn,
		reader: bufio.NewReaderSize(conn, 4096),
	}
}
