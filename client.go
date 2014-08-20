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
	reader *bufio.Reader
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
	}, nil
}

// currently this is synchronous
func (c *Client) exec(command string, args ...string) ([]byte, error) {
	data := bytesForCommand(command, args)

	_, err := c.conn.Write(data)
	if err != nil {
		return nil, err
	}

	return c.readTillCRLF()
}

// TODO(ttacon): change this to use bufio.Scanner/SplitFunc
func (c *Client) readTillCRLF() ([]byte, error) {
	buf, err := c.reader.ReadBytes(crByte)
	if err != nil {
		return nil, err
	}
	lf, err := c.reader.ReadByte()
	if err != nil {
		return nil, err
	}
	if lf != lfByte {
		return nil, fmt.Errorf("expected end of wire format delimeter, got %v", lf)
	}

	buf = buf[0 : len(buf)-1]
	return buf, nil
}

func (c *Client) stringResp(data []byte) (string, error) {
	if len(data) == 0 || data[0] != okByte {
		return "", fmt.Errorf(
			"invalid response data, cannot read string, data was: %v", data)
	}

	return string(data[1:]), nil
}

func (c *Client) boolResp(data []byte) (bool, error) {
	if len(data) == 0 {
		return false, fmt.Errorf(
			"invalid response data, cannot read OK/bool, data was: %v", data)
	}

	if data[0] == errByte {
		return false, fmt.Errorf(string(data[1:]))
	}

	if data[0] != okByte {
		return false, fmt.Errorf(
			"invalid response data, cannot read OK/bool, data was: %v", data)
	}

	return string(data[1:]) == "OK", nil
}

func (c *Client) bulkString(data []byte) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf(
			"invalid response data, cannot read OK/bool, data was: %v", data)
	}

	if data[0] == errByte {
		return "", fmt.Errorf(string(data[1:]))
	}

	if data[0] != sizeByte {
		return "", fmt.Errorf(
			"invalid response data, cannot read OK/bool, data was: %v", data)
	}

	ln, err := strconv.ParseInt(string(data[1:]), 10, 64)
	if err != nil {
		return "", err
	}

	if ln < 0 {
		return "", nil
	}

	respBytes, err := c.readTillCRLF()
	return string(respBytes), err
}

func (c *Client) stringSlice(data []byte) ([]string, error) {
	if len(data) == 0 || data[0] != countByte {
		return nil, fmt.Errorf(
			"invalid response data, cannot read string slice, data was: %v", data)
	}

	n, err := strconv.ParseInt(string(data[1:]), 10, 64)
	if err != nil {
		return nil, err
	}

	var (
		res []string
		i   int64 = 0
	)
	for ; i < n; i++ {
		next, err := c.readTillCRLF()
		if err != nil {
			// TODO(ttacon): do we need to clean the wire/conn up?
			return nil, err
		}

		// sanity check
		if len(next) < 2 {
			return nil, fmt.Errorf("invalid amount of data to read: %v", string(next))
		}
		// TODO(ttacon): use encoding/binary?
		ln, err := strconv.ParseInt(string(next[1:]), 10, 64)
		if err != nil {
			return nil, err
		}

		val, err := c.readTillCRLF()
		if err != nil {
			return nil, err
		}
		if int64(len(val)) != ln {
			return nil, fmt.Errorf(
				"read invalid amount of data off wire, expected %d bytes got %v",
				ln,
				val)
		}

		res = append(res, string(val))

	}
	return res, nil
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
)

var (
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
		buffer.Write([]byte(s))
		buffer.Write(crlfBytes)
	}

	return buffer.Bytes()
}
