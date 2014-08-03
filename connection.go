package redis

func (c *Client) Auth(password string) error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Echo(message string) (string, error) {
	// TODO(ttacon): do it
	return "", nil
}

func (c *Client) Ping() (string, error) {
	// TODO(ttacon): do it
	return "", nil
}

func (c *Client) Quit() error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Select(database int) error {
	// TODO(ttacon): do it
	return nil
}
