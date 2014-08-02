package redis

func (c *Client) Discard() error {
	// TODO(ttacon): do it
	return nil
}

// TODO(ttacon): does this ([][]string) cover all cases?
func (c *Client) Exec() ([][]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) Multi() error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Unwatch() error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Watch(key string, keys ...string) error {
	// TODO(ttacon): do it
	return nil
}
