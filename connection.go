package redis

import "fmt"

func (c *Client) Auth(password string) error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Echo(message string) (string, error) {
	resp, err := c.exec("ECHO", message)
	if err != nil {
		return "", err
	}
	return c.bulkString(resp)
}

func (c *Client) Ping() (string, error) {
	resp, err := c.exec("PING")
	if err != nil {
		return "", err
	}
	return c.stringResp(resp)
}

func (c *Client) Quit() (bool, error) {
	resp, err := c.exec("QUIT")
	if err != nil {
		return false, err
	}
	return c.boolResp(resp)
}

func (c *Client) Select(database int) (bool, error) {
	// TODO(ttacon): do it
	resp, err := c.exec("SELECT", fmt.Sprintf("%d", database))
	if err != nil {
		return false, err
	}

	return c.boolResp(resp)
}
