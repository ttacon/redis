package redis

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"strconv"
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

// currently this is synchronous
func (c *Client) exec(command string, args ...string) ([]byte, error) {
	data := bytesForCommand(command, args)

	_, err := c.conn.Write(data)
	if err != nil {
		return nil, err
	}

	buf, err := c.reader.ReadBytes(crlfBytes)
	if err != nil {
		return nil, err
	}

	fmt.Printf("resp: %q\n", string(buf))
	return buf, nil
}

// ----------------------------------------------------------------------------
// Wire
// ----------------------------------------------------------------------------

// protocol's special bytes
const (
	crByte    byte = byte('\r')
	lfByte         = byte('\n')
	spaceByte      = byte(' ')
	errByte        = byte('-')
	okByte         = byte('+')
	countByte      = byte('*')
	sizeByte       = byte('$')
	numByte        = byte(':')
	trueByte       = byte('1')

	crlfBytes = []byte{crByte, lfByte}
)

func bytesForCommand(command string, args []string) []byte {
	cmdBytes := []byte(command)

	buffer := &bytes.Buffer{}

	buffer.WriteByte(countByte)
	buffer.Write([]byte(strconv.Itoa(len(args) + 1)))
	buffer.Write(crlfBytes)
	buffer.WriteByte(sizeByte)
	buffer.Write([]byte(strconv.Itoa(len(cmdBytes))))
	buffer.Write(crlfBytes)
	buffer.Write(cmdBytes)
	buffer.Write(crlfBytes)

	for _, s := range args {
		buffer.WriteByte(sizeByte)
		buffer.Write([]byte(strconv.Itoa(len(s))))
		buffer.Write(crlfBytes)
		buffer.Write(s)
		buffer.Write(crlfBytes)
	}

	return buffer.Bytes()
}
