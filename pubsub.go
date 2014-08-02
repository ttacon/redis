package redis

// TODO(ttacon): is there a return type for this command? none is documented
func (c *Client) PSubscribe(pattern string, patterns ...string) error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) Publish(channel, message string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

// TODO(ttacon): double check return type
func (c *Client) PUnsubscribe(patterns ...string) ([]string, error) {
	// TODO(ttacon): do it
	return "", nil
}

// TODO(ttacon): double check return type, same issue as PSubscribe
func (c *Client) Subscribe(channel string, channels ...string) error {
	// TODO(ttacon): nil
	return nil
}

// TODO(ttacon): again, double check return type
func (c *Client) Unsubscribe(channels ...string) error {
	// TODO(ttacon): do it
	return nil
}
